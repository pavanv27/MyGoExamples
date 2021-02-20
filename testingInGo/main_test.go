package sum

import (
	"fmt"
	"testing"
)

type testpair struct {
	values []int
	exsum  int
}

var tests = []testpair{
	{[]int{1, 2, 3, 4, 5}, 15},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 45},
	{[]int{-1, -2, -3, -4, -5}, -15},
	{[]int{}, 0},
}

func TestSum(t *testing.T) {
	for _, tp := range tests {
		got := Sum(tp.values)
		if got != tp.exsum {
			t.Errorf("Sum(%v) = %d; want %d", tp.values, got, tp.exsum)
		}
	}
}

func ExampleSum() {
	res := Sum(tests[0].values)
	fmt.Println(res)
	// Output:
	// 15
}

func ExampleSquare() {
	res := Square(tests[0].values)
	fmt.Println(res)
	// Output:
	// 225
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(tests[0].values)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(tests[0].values)
	}
}
