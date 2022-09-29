package iteration

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t *testing.T, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected iteration messages, expected %q, while got %q", expected, got)
	}

}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	assertCorrectMessage(t, repeated, expected)
}

func BeanchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func ExampleRepeat() {
	repeat := Repeat("Yay!", 100)
	fmt.Println(repeat)
	// "Yay! * 100"
}
