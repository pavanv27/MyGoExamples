package main

import "fmt"

// Fibonacci sequence :  each number is the sum of the two preceding numbers, starting from 0 and 1
// Series : 1,1,2,3,5,8,13,21...

func main() {

	for n := 0; n < 9; n++ {
		fmt.Println("Fibonacci series", n, "element is", fibonacciEasy(n))
	}

	for n := 0; n < 9; n++ {
		fmt.Println("Fibonacci series", n, "element is", fibonacciRecursionOpt(n, 1, 1))
	}

	for n := 0; n < 9; n++ {
		fmt.Println("Fibonacci series", n, "element is", fibonacciNoRecursion(n))
	}

}

// fibonacciEasy returns the nth element of the fibonacci series
func fibonacciEasy(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonacciEasy(n-1) + fibonacciEasy(n-2)
}

// fibonacciRecursionOpt shows the further optimized way to achieve it
func fibonacciRecursionOpt(n int, a int, b int) int {
	if n == 0 || n == 1 {
		return b //b holds the final added value from each recursion
	}
	// since we are calling recursive function twice we have to have 2 variables
	return fibonacciRecursionOpt(n-1, b, a+b)
}

// fibonacciNoRecursion shows the no recursion way to achieve it
func fibonacciNoRecursion(n int) int {
	a := 1 //1st
	b := 1 //2nd
	c := 1 //out
	for i := 1; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c //c holds the final added value after each iteration
}

//Series : 1,1,2,3,5,8,13,21...
