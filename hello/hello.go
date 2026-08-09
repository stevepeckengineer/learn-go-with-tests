// Hello world, can accept a name to personalize, and a language for the greeting
// Default language is English if language unavailable
package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	frenchPrefix  = "Bonjour, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world!"
	}

	return greetingPrefix(lang) + name
}

// set greeting prefix by supported language or default to english
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
