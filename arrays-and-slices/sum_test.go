package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("A collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectHelper(t, got, want, numbers)
	})
}

func assertCorrectHelper(t *testing.T, got int, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d when given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	reflectHelper(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("Make sums of slices with length > 1", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		reflectHelper(t, got, want)
	})

	t.Run("return 0 for empty strings", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		reflectHelper(t, got, want)
	})
}

func reflectHelper(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
