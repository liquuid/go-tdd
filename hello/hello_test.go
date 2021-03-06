package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

	t.Run("Saying hello!", func(t *testing.T) {
		got := hello("Cris", "")
		want := "Hello, Cris!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("spanish mode", func(t *testing.T) {
		got := hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("french mode", func(t *testing.T) {
		got := hello("Claudia", "French")
		want := "Bonjour, Claudia!"
		assertCorrectMessage(t, got, want)
	})
}
