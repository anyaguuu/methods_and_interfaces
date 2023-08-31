package base

import (
	"sort"
	"strings"
)

type Base struct {
	baseMap map[string]int
}

func New(inputMap map[string]int) Base { // returns pointer to base
	return Base{baseMap: inputMap} // maps words to frequencies
}

// implements the interface!
// return up to the first 10 words alphabetically that have the given base
func (b Base) Complete(base string) []string {
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
