package array

import "testing"

func assertCorrectNumbers(t *testing.T, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected array numbers sum, expected %d, while got %d", expected, got)
	}

}

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	assertCorrectNumbers(t, got, want)
}
