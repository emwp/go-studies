package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Everton", "")
		want := "Hello, Everton"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to \"Hello, World\" when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Everton", "Spanish")
		want := "Hola, Everton"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Everton", "French")
		want := "Bonjour, Everton"
		assertCorrectMessage(t, got, want)
	})

}
