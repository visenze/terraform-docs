
package doc

import "sort"

// Output represents a Terraform output.
type Module struct {
	Type        string
	Name        string
	Version     string
}

// HasDescription indicates if a Terraform output has a description.
func (o *Module) HasDescription() bool {
	return false
}

type modulesSortedByType []Module

func (a modulesSortedByType) Len() int {
	return len(a)
}

func (a modulesSortedByType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a modulesSortedByType) Less(i, j int) bool {
	var result bool

	if a[i].Type != a[j].Type {
		result = a[i].Type < a[j].Type
	} else {
		result = a[i].Name < a[j].Name
	}
	return result
}

// SortModuleByType sorts a list of modules by type.
func SortModuleByType(modules []Module) {
	sort.Sort(modulesSortedByType(modules))
}
