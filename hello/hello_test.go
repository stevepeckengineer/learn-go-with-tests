// Tests for hello world
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello with name", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertEqual(t, got, want)
	})

	t.Run("Saying 'hello, world!' if no name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertEqual(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, world!"
		assertEqual(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, world!"
		assertEqual(t, got, want)
	})
}

func assertEqual(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
