package main

import (
	"strings"
)

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	trimmedText := strings.Fields(loweredText)

	return trimmedText
}
