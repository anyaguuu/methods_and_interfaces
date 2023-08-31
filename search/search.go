package search

import "strings"

// BinarySearch finds the first word in words that starts with prefix.  words
// must be sorted.  Returns the index of the first word in words that starts
// with prefix. The second return value is true if such a word was found
// (meaning the index is valid) or false if no such word exists in words
// (meaning the index is invalid).
func BinarySearch(prefix string, words []string) (int, bool) {
	if len(words) == 0 {
		return -1, false
	}

	index := 0
	low := 0
	high := len(words) - 1
	for low <= high {
		mid := (low + high) / 2
		word := words[mid]
		if strings.HasPrefix(word, prefix) || word > prefix {
			// Search lower half
			high = mid - 1
			index = mid
		} else {
			// Search upper half
			low = mid + 1
		}
	}

	if !strings.HasPrefix(words[index], prefix) {
		return -1, false
	}

	return index, true
}
