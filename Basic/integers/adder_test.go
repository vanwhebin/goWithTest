package integers

import (
	"fmt"
	"testing"
)

func getAssertCorrectInteger(t *testing.T, sum, want int) {
	t.Helper()
	if sum != want {
		t.Errorf("Sum expected: %d, got %d", want, sum)
	}

}

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expect := 4

	getAssertCorrectInteger(t, sum, expect)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
