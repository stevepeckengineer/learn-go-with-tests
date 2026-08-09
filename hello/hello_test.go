// Tests for hello world
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello with name", func(t *testing.T){
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q \n want %q", got, want)
		}
	})
	t.Run("Saying hello, world if no name", func(t *testing.T){
		got := Hello("")
		want := "Hello, world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}