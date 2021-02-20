package main

import "testing"

func BenchmarkFactorialEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialEasy(30)
	}
}

func BenchmarkFactorialTailCallOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialRecursionOpt(30, 1)
	}
}

func BenchmarkFactorialNoRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialNoRecursion(30)
	}
}
