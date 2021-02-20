package main

import "testing"

func BenchmarkFibonacciEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciEasy(30)
	}
}

func BenchmarkFibonacciRecursionOpt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursionOpt(30, 1, 1)
	}
}

func BenchmarkFibonacciNoRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciNoRecursion(30)
	}
}
