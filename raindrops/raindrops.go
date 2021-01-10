package raindrops

import (
	"strconv"
	"strings"
)

var Mapping = []struct {
	factor int
	sound  string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {
	var sounds strings.Builder

	for _, mapping := range Mapping {
		hasFactor := number%mapping.factor == 0

		if hasFactor {
			sounds.WriteString(mapping.sound)
		}
	}

	if sounds.Len() == 0 {
		sounds.WriteString(strconv.Itoa(number))
	}

	return sounds.String()
}
