package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Cristina", "")
		want := "Hello, Cristina"
	
		assertCorrentMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when arg is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrentMessage(t, got, want)
	})

	t.Run("If language is Spanish respond with 'Hola'", func(t *testing.T) {
		got := Hello("Barry", spanish)
		want := "Hola, Barry"

		assertCorrentMessage(t, got, want)
	})

	t.Run("If language is French respond with 'Bonjour'", func(t *testing.T) {
		got := Hello("Piere", french)
		want := "Bonjour, Piere"

		assertCorrentMessage(t, got, want)
	})

	t.Run("If language is German respond with 'Hallo'", func(t *testing.T) {
		got := Hello("Gary", german)
		want := "Hallo, Gary"

		assertCorrentMessage(t, got, want)
	})
}

func assertCorrentMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
