package iteration

import "strings"

const defaultCount = 5

func Repeat(character string, count int) string {
	var repeated strings.Builder

	if count == 0 {
		count = defaultCount
	}

	for i := 0; i < count; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
