package array

import (
	"reflect"
	"testing"
)

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

func TestSumWithSlice(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 1}
		got := Sum(numbers)
		want := 11

		assertCorrectNumbers(t, got, want)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumWithSlice(numbers)
		want := 6

		assertCorrectNumbers(t, got, want)
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 5})
	want := []int{3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v while want %v", got, want)
	}

}

func TestSumAppendDynamic(t *testing.T) {
	got := SumAppendDynamicParams([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v while want %v", got, want)
	}

}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v while want %v", got, want)
	}

}
