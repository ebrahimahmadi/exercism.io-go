package acronym

import (
	"strings"
)

func Abbreviate(s string) (out string) {
	s = filterChr("_", s)
	s = filterChr("-", s)

	words := strings.Fields(s)

	for _, word := range words {
		out += word[:1]
	}

	return strings.ToUpper(out)
}

func filterChr(toFilter string, sourceString string) string {
	replaceWith := " "
	allOccurance := -1

	filtered := strings.Replace(sourceString, toFilter, replaceWith, allOccurance)
	return filtered
}
