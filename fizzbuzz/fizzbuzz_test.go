package fizzbuzz

import "testing"

func TestFizzBuzzModuloV1(t *testing.T) {
	fizzBuzzModuloV1()
}

func BenchmarkFizzBuzzModuloV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzzModuloV1()
	}
}

func BenchmarkFizzBuzzModuloV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzzModuloV2()
	}

}

func BenchmarkFizzBuzzModuloV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzzModuloV3()
	}
}

func BenchmarkFizzBuzzNoModulo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fizzBuzzNoModulo()
	}
}
