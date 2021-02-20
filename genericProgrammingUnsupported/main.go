package main

import (
	"fmt"
)

func main() {

	// string slice
	s := []string{"this", "is", "a", "test", "input"}
	fmt.Println(s)
	ReverseStringSlice(s)
	fmt.Println(s)

	// int slice
	i := []int{1, 2, 3, 4, 5}
	fmt.Println(i)
	ReverseIntSlice(i)
	fmt.Println(i)
}

// ReverseIntSlice reverses only slice of int
func ReverseIntSlice(s []int) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// ReverseStringSlice reverses only slice of string
func ReverseStringSlice(s []string) {
	first := 0
	last := len(s) - 1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

// No Generic function popssible. Below deosn't work !!

// func ReverseSliceGeneric(s interface{}) {
// 	first := 0
// 	last := 0
// 	switch x := s.(type) {
// 	case int:
// 	case string:
// 		last = len(x) - 1
// 		for first < last {
// 			s[first], s[last] = s[last], s[first]
// 			first++
// 			last--
// 		}
// 	}
// }
