package frequency

import (
	"sort"
	"strings"

	"github.com/anyaguuu/methods_and_interfaces/search"
)

// must store words and their frequencies
type freq struct {
	words       []string
	frequencies map[string]int
}

func New(inputMap map[string]int) freq {
	// extract the keys from map
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	sort.Strings(keys) // sort keys alphabetically
	return freq{words: keys, frequencies: inputMap}
}

// implement interface
// return up to the first ten words in decreasing
// frequency order that have the given base

func (f freq) Complete(prefix string) []string {
	var words []string

	// find idx of first word w given prefix
	index, ok := search.BinarySearch(prefix, f.words)

	if !ok {
		return words
	}

	// find all words with prefix
	for index < len(f.words) {
		word := f.words[index]
		if strings.HasPrefix(word, prefix) {
			words = append(words, word)
			index++
		} else {
			break // bc sorted
		}
	}

	// sort matching words based on frequency vals in descending order
	sort.Slice(words, func(i, j int) bool {
		return f.frequencies[words[i]] > f.frequencies[words[j]]
	})

	if len(words) >= 10 {
		words = words[:10]
	}

	return words
}
