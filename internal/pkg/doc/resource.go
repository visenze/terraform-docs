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
	return a[i].Type < a[j].Type
}

// SortResourcesByName sorts a list of outputs by name.
func SortResourcesByType(resources []Resource) {
	sort.Sort(resourcesSortedByType(resources))
}
