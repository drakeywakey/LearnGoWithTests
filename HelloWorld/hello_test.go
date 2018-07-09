package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s', but wanted '%s'", got, want)
		}
	}

	t.Run("saying hello to myself", func(t *testing.T) {
		got := Hello("Drake", "")
		want := "Sup, Drake?"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Sup, World?"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Drake", "Spanish")
		want := "Cenar, Drake?"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Drake", "French")
		want := "Ca va, Drake?"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in German", func(t *testing.T) {
		got := Hello("Drake", "German")
		want := "Guten tag, Drake?"
		assertCorrectMessage(t, got, want)
	})
}
