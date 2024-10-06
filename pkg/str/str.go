package str

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.English)

// Converts snake_case to Title Case
func SnakeToTitle(snakeStr string) string {
	// Split the string by underscores
	words := strings.Split(snakeStr, "_")
	for i, word := range words {
		// Convert first letter to uppercase, the rest to lowercase
		words[i] = titleCaser.String(word)
	}

	return strings.Join(words, " ")
}

func ToTitle(s string) string {
	return titleCaser.String(s)
}
