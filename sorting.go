package strcmp

import "sort"

// Assert that Naturally implements sort.Interface.
var _ sort.Interface = Naturally{}

// Naturally is a sort.Interface implementation
// for sorting a string slice in natural order.
type Naturally []string

func (s Naturally) Less(i, j int) bool {
	return Natural(s[i], s[j]) < 0
}

func (s Naturally) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Naturally) Len() int {
	return len(s)
}
