package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("lookup word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, err := dict.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("word not in dictionary", func(t *testing.T) {
		dict := Dictionary{}
		_, err := dict.Search("test")

		assertError(t, err, ErrMissingWord)
	})

	t.Run("Add word", func(t *testing.T) {
		dictionary := Dictionary{}
		word, definition := "test", "this is just a test"

		err := dictionary.Add(word, definition)
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("word already in dictionary", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "testing again"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)
		assertNoError(t, err)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("cannot update word if not in dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "testing again"

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrMissingWord)
		_, err = dictionary.Search(word)
		assertError(t, err, ErrMissingWord)
	})

	t.Run("delete", func(t *testing.T) {
		word, definition := "test", "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertNoError(t, err)
		
		_, err = dictionary.Search(word)
		assertError(t, err, ErrMissingWord)
	})
}

func assertDefinition(t testing.TB, dict Dictionary, word, definition string) {
	got, err := dict.Search(word)
	assertNoError(t, err)
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got unexpected error: ", err)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()

	if err == nil {
		t.Fatal("wanted err, got none")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
