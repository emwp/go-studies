package lists

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of five numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make the sums of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9}, []int{1, 3, 5, 9})
		want := []int{2, 9, 17}
		checkSums(t, got, want)
	})

	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 10, 20})
		want := []int{0, 30}
		checkSums(t, got, want)
	})
}
