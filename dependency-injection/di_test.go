package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("simple buffer test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
