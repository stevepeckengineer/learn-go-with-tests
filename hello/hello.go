// Hello world, can accept a name to personalize, and a language for the greeting
// Default language is English if language unavailable
package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french = "French"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world!"
	}

	prefix := englishPrefix

	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}