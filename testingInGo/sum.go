//Package sum provides us functions which give sum of integers , square of sumof integers and cube of sum of integers
package sum

import "fmt"

//Sum gives the sum of all the integers of the slice fed to the function
func Sum(in []int) int {
	res := 0
	for _, r := range in {
		res += r
	}
	return res
}

//Square gives the square of the sum of all integers of the slice fed to the function
func Square(in []int) int {
	res := 0
	for _, r := range in {
		res += r
	}
	//fmt.Println("res=", res*res)
	return res * res
}

func Cube(in []int) int {
	res := 0
	for _, r := range in {
		//fmt.Println("r=",r)
		res += r
	}
	fmt.Println("res=", res*res*res)
	return res * res * res
}
