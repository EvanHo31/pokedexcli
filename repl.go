package main

import (
	"strings"
)

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	splitted := strings.Split(trimmed, " ")
	cleaned := []string{}
	for _, word := range splitted {
		clean_word := strings.TrimSpace(word)
		if len(clean_word) == 0 {
			continue
		}
		clean_word = strings.ToLower(clean_word)
		cleaned = append(cleaned, clean_word)
	}
	return cleaned
}
