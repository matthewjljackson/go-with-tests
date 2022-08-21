package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Cristina")
	want := "Hello, Cristina"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
