package citsaae

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	if strings.Join(word1, "") == strings.Join(word2, "") {
		return true
	}
	return false
}
