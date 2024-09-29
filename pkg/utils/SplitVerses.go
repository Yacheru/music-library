package utils

import "strings"

func SplitVerses(s string) string {
	return strings.ReplaceAll(s, "\n\n", " ")
}
