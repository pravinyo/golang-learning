package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Pravin")
	want := "Hello, Pravin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
