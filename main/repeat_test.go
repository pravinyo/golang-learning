package main

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Should repeat a by 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("should repeat b by 10 times", func(t *testing.T) {
		repeated := Repeat("b", 10)
		expected := "bbbbbbbbbb"
		assertCorrectMessage(t, repeated, expected)
	})
}

// go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
