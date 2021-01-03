package twofer

import (
	"fmt"
)

func ShareWith(name string) string {
	if name == ""{
		defaultName := "you"
		name = defaultName
	}

	message := fmt.Sprintf("One for %v, one for me.", name)
	return message
}
