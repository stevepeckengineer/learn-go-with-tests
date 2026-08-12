package dictionary


const (
	ErrMissingWord = DictionaryErr("Word not found!")
	ErrWordExists  = DictionaryErr("Word already exists!")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, isFound := d[word]

	if !isFound {
		return "", ErrMissingWord
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrMissingWord:
		d[word] = definition
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}

}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
		return nil
	default:
		return err
	}

}
