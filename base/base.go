package base

// provides completion based on search order

import (
	"sort"
	"strings"
)

type base struct {
	words []string
}

func New(input map[string]int) base { // returns pointer to base
	// we want to get the words and sort them
	keys := make([]string, 0, len(input))

	for word := range input {
		keys = append(keys, word)
	}

	sort.Strings(keys)

	return base{words: keys}
}

// implements the interface!
// return up to the first 10 words alphabetically that have the given base
func (b base) Complete(base string) []string {
	ret := make([]string, 0, 10) // so empty slice of strings, cap = 10
	for word := range b.baseMap {
		if strings.Index(word, base) == 0 {
			ret = append(ret, word)
		}
		if len(ret) == 10 { // return when reached 10
			sort.Strings(ret)
			return ret
		}
	}
	// now sort alphabetically
	sort.Strings(ret)
	return ret
}
