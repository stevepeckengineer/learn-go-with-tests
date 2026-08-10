package dictionary

import (
	"errors"
)

var ErrMissingWord = errors.New("Word not found!")
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, isFound := d[word]
	
	if !isFound {
		return "", ErrMissingWord
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}