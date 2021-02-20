package fizzbuzz

import "fmt"

func fizzBuzzModuloV1() {
	for i := 1; i < 101; i++ {
		s := fmt.Sprintf("%d", i)
		if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else if i%15 == 0 {
			s = "FizzBuzz"
		}
		_ = s
		//fmt.Println(s)
	}
}

func fizzBuzzModuloV2() {
	for i := 1; i < 101; i++ {
		s := fmt.Sprintf("%d", i)
		if i%3 == 0 {
			s = "Fizz"
			if i%5 == 0 {
				s += "Buzz"
			}
		} else if i%5 == 0 {
			s = "Buzz"
		}
		//fmt.Println(s)
	}
}

func fizzBuzzModuloV3() {
	for i := 1; i < 101; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if s == "" {
			s = fmt.Sprintf("%d", i)
		}
		//fmt.Println(s)
	}
}

func fizzBuzzNoModulo() {
	c3 := 0
	c5 := 0
	for i := 1; i < 101; i++ {
		s := ""
		c3++
		c5++
		if c3 == 3 {
			s += "Fizz"
			c3 = 0
		}
		if c5 == 5 {
			s += "Buzz"
			c5 = 0
		}
		if s == "" {
			s = fmt.Sprintf("%d", i)
		}
		//fmt.Println(s)
	}
}
