package service

import (
	"bufio"
	"encoding/json"
	"strings"
	"unicode"
)

// WordCount stores the count of unique words and
// a map of {word: count}.
type WordCount struct {
	Count   int            `json:"count,omitempty"`
	WordMap map[string]int `json:"wordmap,omitempty"`
}

// CountWordsInText cleans the string of punctuation,
// makes a map of the words and their counts, then
// returns the WordCount object as json.
func CountWordsInText(s string) ([]byte, error) {
	s = strings.ToLower(s)
	strSlice := cleanWords(s)
	m := make(map[string]int)
	for _, word := range strSlice {
		m[word]++
	}

	w := WordCount{
		Count:   len(m),
		WordMap: m,
	}
	wJSON, err := json.Marshal(w)
	if err != nil {
		return []byte{}, err
	}
	return wJSON, nil
}

// cleanSpace separate the words from the spaces.
func cleanWords(s string) []string {
	var strSlice []string
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		stripped := strings.TrimFunc(scanner.Text(), stripNonWord)
		if stripped != "" {
			strSlice = append(strSlice, stripped)
		}
	}
	return strSlice
}

// stripNonWord will remove leading and trailing unicode
// points from the word if they are not letters.
func stripNonWord(r rune) bool {
	return !unicode.IsLetter(r)
}
