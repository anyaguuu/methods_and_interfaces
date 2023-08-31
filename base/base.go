package base

// provides completion based on search order

import (
	"sort"
	"strings"

	"github.com/anyaguuu/methods_and_interfaces/search"
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
func (b base) Complete(prefix string) []string {
	var words []string
	index, ok := search.BinarySearch(prefix, b.words) // finds the first word in words that starts with prefix
	if !ok {
		return words
	}

	// return up to 10 words w input prefix
	last := min(index+10, len(b.words)) //last index we want to go til
	for i := index; i < last; i++ {
		word := b.words[i]
		if strings.HasPrefix(word, prefix) {
			words = append(words, word)
		} else {
			break
		}
	}
	return words
}
