package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return 0, errors.New("string are not the same")
	}

	for index := range a {
		if b[index] != a[index] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
