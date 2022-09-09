package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new words", func(t *testing.T) {
		dictionary := Dictionary{}
		err1 := dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err2 := dictionary.Search("test")

		assertNoError(t, err2)
		assertNoError(t, err1)
		assertStrings(t, got, want)
	})

	t.Run("existing words", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just a test")

		assertError(t, err, ErrWordExists)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
