package doc

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/segmentio/terraform-docs/internal/pkg/fs"
	hclparser "github.com/segmentio/terraform-docs/internal/pkg/hcl"
)

// Doc represents a Terraform module.
type Doc struct {
	Comment string
	Inputs  []Input
	Outputs []Output
	Resources []Resource
	Modules []Module
}

// HasComment indicates if the document has a comment.
func (d *Doc) HasComment() bool {
	return len(d.Comment) > 0
}

// HasInputs indicates if the document has inputs.
func (d *Doc) HasInputs() bool {
	return len(d.Inputs) > 0
}

// HasOutputs indicates if the document has outputs.
func (d *Doc) HasOutputs() bool {
	return len(d.Outputs) > 0
}

// HasResources indicates if the document has resources.
func (d *Doc) HasResources() bool {
	return len(d.Resources) > 0
}

// HasModules indicates if the document has modules.
func (d *Doc) HasModules() bool {
	return len(d.Modules) > 0
}

// Value represents a Terraform value.
type Value struct {
	Type  string
	Value interface{}
}

// IsAggregateType indicates if a Terraform value is an aggregate type.
func (v Value) IsAggregateType() bool {
	return v.Type == "list" || v.Type == "map"
}

// CreateFromPaths creates a new document from a list of file or directory paths.
func CreateFromPaths(paths []string) (*Doc, error) {
	names := make([]string, 0)

	for _, path := range paths {
		if fs.DirectoryExists(path) {
			matches, err := filepath.Glob(fmt.Sprintf("%s/*.tf", path))
			if err != nil {
				log.Fatal(err)
			}

			names = append(names, matches...)
		} else if fs.FileExists(path) {
			names = append(names, path)
		}
	}

	files := make(map[string]*ast.File)

	for _, name := range names {
		bytes, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		ast, err := hcl.ParseBytes(bytes)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		files[name] = ast
	}

	return Create(files), nil
}

// Create creates a new document from a map of filenames to *ast.Files.
func Create(files map[string]*ast.File) *Doc {
	doc := new(Doc)

	for name, file := range files {
		objects := file.Node.(*ast.ObjectList)

		doc.Inputs = append(doc.Inputs, getInputs(objects)...)
		doc.Outputs = append(doc.Outputs, getOutputs(objects)...)
		doc.Resources = append(doc.Resources, getResources(objects)...)
		doc.Modules = append(doc.Modules, getModules(objects)...)

		filename := filepath.Base(name)
		comments := file.Comments
		if filename == "main.tf" && len(comments) > 0 {
			doc.Comment = header(comments[0])
		}
	}

	return doc
}

// getInputs returns a list of inputs from an ast.ObjectList.
func getInputs(list *ast.ObjectList) []Input {
	var result []Input

	for _, item := range list.Items {
		if isItemOfKindVariable(item) {
			result = append(result, Input{
				Name:        getItemName(item),
				Description: getItemDescription(item),
				Default:     getItemDefault(item),
				Type:        getItemType(item),
			})
		}
	}

	return result
}

// getOutputs returns a list of outputs from an ast.ObjectList.
func getOutputs(list *ast.ObjectList) []Output {
	var result []Output

	for _, item := range list.Items {
		if isItemOfKindOutput(item) {
			result = append(result, Output{
				Name:        getItemName(item),
				Description: getItemDescription(item),
			})
		}
	}

	return result
}

// getResources returns a list of resources from an ast.ObjectList.
func getResources(list *ast.ObjectList) []Resource {
	var result []Resource

	for _, item := range list.Items {
		if isItemOfKindResource(item) {
			result = append(result, Resource{
				Name: getItemResourceName(item),
				Type: getItemResourceType(item),
			})
		}
	}

	return result
}

// getModules returns a list of resources from an ast.ObjectList.
func getModules(list *ast.ObjectList) []Module {
	var result []Module

	for _, item := range list.Items {
		if isItemOfKindModule(item) {
			result = append(result, Module{
				Name:    getItemModuleName(item),
				Type:    getItemModuleType(item),
				Version: getItemModuleVersion(item),
			})
		}
	}

	return result
}



func getItemByKey(items []*ast.ObjectItem, key string) *Value {
	for _, item := range items {
		if isItemOfKind(item, key) {
			result := new(Value)

			if literal, ok := item.Val.(*ast.LiteralType); ok {
				result.Type = "string"

				if value, ok := literal.Token.Value().(string); ok {
					result.Value = value
				} else {
					result.Value = literal.Token.Text
				}

				return result
			}

			if _, ok := item.Val.(*ast.ObjectType); ok {
				result.Type = "map"

				data, err := hclparser.ParseAstNode(&item.Val, result.Type)
				if err == nil {
					result.Value = data
				}

				return result
			}

			if _, ok := item.Val.(*ast.ListType); ok {
				result.Type = "list"

				data, err := hclparser.ParseAstNode(&item.Val, result.Type)
				if err == nil {
					result.Value = data
				}

				return result
			}

			return nil
		}
	}

	return nil
}

func getItemDefault(item *ast.ObjectItem) *Value {
	items := item.Val.(*ast.ObjectType).List.Items
	return getItemByKey(items, "default")
}

func getItemDescription(item *ast.ObjectItem) string {
	var result string

	items := item.Val.(*ast.ObjectType).List.Items

	var description = getItemByKey(items, "description")
	if description != nil {
		result = description.Value.(string)
	}

	if result == "" {
		if item.LeadComment != nil {
			result = getItemDescriptionFromComment(item.LeadComment.List)
		}
	}

	return result
}

func getItemDescriptionFromComment(comments []*ast.Comment) string {
	var result string

	for _, comment := range comments {
		line := strings.TrimSpace(comment.Text)
		line = strings.TrimPrefix(line, "#")
		line = strings.TrimPrefix(line, "//")
		result += strings.TrimSpace(line)
	}

	return result
}

func getItemName(item *ast.ObjectItem) string {
	name, err := strconv.Unquote(item.Keys[1].Token.Text)
	if err != nil {
		name = item.Keys[1].Token.Text
	}

	return name
}

func getItemResourceName(item *ast.ObjectItem) string {
	val, err := strconv.Unquote(item.Keys[2].Token.Text)
	if err != nil {
		val = item.Keys[2].Token.Text
	}

	return val
}

func getItemResourceType(item *ast.ObjectItem) string {
	val, err := strconv.Unquote(item.Keys[1].Token.Text)
	if err != nil {
		val = item.Keys[1].Token.Text
	}

	return val
}

func getItemModuleName(item *ast.ObjectItem) string {
	val, err := strconv.Unquote(item.Keys[1].Token.Text)
	if err != nil {
		val = item.Keys[1].Token.Text
	}

	return val
}

func getItemModuleType(item *ast.ObjectItem) string {
	items := item.Val.(*ast.ObjectType).List.Items
	val := getItemByKey(items, "source").Value.(string)

	strToSearch := "/modules"

	index := strings.Index(val, strToSearch)
	if index  >= 0 {
		val = val[index + len(strToSearch) + 1:]
	}

	return val
}

func getItemModuleVersion(item *ast.ObjectItem) string {
	var version string

	items := item.Val.(*ast.ObjectType).List.Items
	val := getItemByKey(items, "version")

	if val != nil {
		version = val.Value.(string)
	}

	return version
}

func getItemType(item *ast.ObjectItem) string {
	var result string

	items := item.Val.(*ast.ObjectType).List.Items

	_type := getItemByKey(items, "type")
	value := getItemByKey(items, "default")
	if _type == nil || _type.Value == "" {
		if value == nil {
			result = "string"
		} else {
			result = value.Type
		}
	} else {
		result = _type.Value.(string)
	}

	return result
}

func isItemOfKind(item *ast.ObjectItem, kind string) bool {
	if len(item.Keys) > 0 {
		return item.Keys[0].Token.Text == kind
	}

	return false
}

func isItemOfKindOutput(item *ast.ObjectItem) bool {
	return isItemOfKind(item, "output")
}

func isItemOfKindVariable(item *ast.ObjectItem) bool {
	return isItemOfKind(item, "variable")
}

func isItemOfKindResource(item *ast.ObjectItem) bool {
	return isItemOfKind(item, "resource")
}

func isItemOfKindModule(item *ast.ObjectItem) bool {
	return isItemOfKind(item, "module")
}

// Header returns the header comment from the list
// or an empty comment. The head comment must start
// at line 1 and start with `/**`.
func header(c *ast.CommentGroup) (comment string) {
	if len(c.List) == 0 {
		return comment
	}

	if c.Pos().Line != 1 {
		return comment
	}

	cm := strings.TrimSpace(c.List[0].Text)

	if strings.HasPrefix(cm, "/**") {
		lines := strings.Split(cm, "\n")

		if len(lines) < 2 {
			return comment
		}

		lines = lines[1 : len(lines)-1]
		for _, l := range lines {
			l = strings.TrimSpace(l)
			switch {
			case strings.TrimPrefix(l, "* ") != l:
				l = strings.TrimPrefix(l, "* ")
			default:
				l = strings.TrimPrefix(l, "*")
			}
			comment += l + "\n"
		}
	}

	return comment
}
