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
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")
		
		assertNoError(t, err)
		assertStrings(t, got, want)
	})
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
		t.Errorf("wanted no error, got %q", err)
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
