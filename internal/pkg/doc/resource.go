package doc

import "sort"

// Output represents a Terraform output.
type Resource struct {
	Type        string
	Name        string
}

// HasDescription indicates if a Terraform output has a description.
func (o *Resource) HasDescription() bool {
	return false
}

type resourcesSortedByType []Resource

func (a resourcesSortedByType) Len() int {
	return len(a)
}

func (a resourcesSortedByType) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a resourcesSortedByType) Less(i, j int) bool {
	var result bool

	if a[i].Type != a[j].Type {
		result = a[i].Type < a[j].Type
	} else {
		result = a[i].Name < a[j].Name
	}
	return result
}

// SortResourcesByName sorts a list of outputs by name.
func SortResourcesByType(resources []Resource) {
	sort.Sort(resourcesSortedByType(resources))
}
