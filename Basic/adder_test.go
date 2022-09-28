package main

import "testing"

func getAssertCorrectInteger(t *testing.T, sum, want int) {
	t.Helper()
	if sum != want {
		t.Errorf("Sum expected: %d, got %d", want, sum)
	}

}

func TestAdder(t *testing.T) {
	sum := Add(2, 4)
	expect := 4

	getAssertCorrectInteger(t, sum, expect)

}
