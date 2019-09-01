package group

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	assertCorrectSum := func(t *testing.T, want int, got int, numbers []int) {
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		got := Sum(numbers)
		want := 15
		assertCorrectSum(t, want, got, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkFSums := func(got []int, want []int, t *testing.T) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkFSums(got, want, t)
	})

	t.Run("sum empty array", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkFSums(got, want, t)
	})
}

func ExampleSum() {
	got := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(got)
	// Output: 15
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(got)
	// Output: [3 9]
}

func ExampleSumAllTails() {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	fmt.Println(got)
	// Output: [2 9]
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
	}
}
