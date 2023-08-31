package frequency

import "strings"

type Frequency struct {
	freqMap map[string]int
}

func New(inputMap map[string]int) Frequency {
	return Frequency{freqMap: inputMap}
}

// implement interface
// return up to the first ten words in decreasing
// frequency order that have the given base

func (f Frequency) Complete(base string) []string {
	newMap := make(map[string]int, len(f.freqMap))
	for word, freq := range f.freqMap {
		if strings.Index(word, base) == 0 { // check if has base first
			newMap[word] = freq
		}
	}

	// now sort new map in decreasing order of frequency
}
