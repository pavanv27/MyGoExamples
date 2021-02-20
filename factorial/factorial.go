package main

import (
	"fmt"
)

/*
Factorial :
A factorial is a function that multiplies a number by every number below it.
For example 5!= 5*4*3*2*1=120.
*/

func main() {
	fmt.Println("Let's Factorial !..")

	n := 5

	fmt.Println("factorialEasy", n, " is", factorialEasy(n))

	fmt.Println("factorialRecursionOpt", n, " is", factorialRecursionOpt(n, 1))
	//function overloading not supported in go

	fmt.Println("factorialNoRecursion of ", n, " is", factorialNoRecursion(n))

}

// factorialEasy shows the easy way of coding to get the factorial
func factorialEasy(n int) int {
	if n == 0 || n == 1 { // 0!=1 and 1!=1
		return 1
	}
	return n * factorialEasy(n-1)
}

// factorialRecursionOpt shows the optimized way of doing it
func factorialRecursionOpt(n int, a int) int {
	if n == 0 || n == 1 {
		return a
	}
	return factorialRecursionOpt(n-1, n*a) //a holds the total "multiplied" value after each recursion
}

// factorialNoRecursion shows the way to code to have best performance
func factorialNoRecursion(n int) int {
	a := 1
	for n != 0 {
		a = n * a // this holds the total "multiplied" value after each iteration
		n = n - 1
	}
	return a
}
