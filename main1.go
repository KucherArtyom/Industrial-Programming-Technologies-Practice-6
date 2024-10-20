package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	fmt.Print("Введите число, которое хотите проверить на простоту: ")
	fmt.Scan(&num)
	if isPrime(num) {
		fmt.Printf("%d является простым числом.\n", num)
	} else {
		fmt.Printf("Первый делитель числа %d: %d\n", num, firstDivisor(num))
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	maxDivisor := int(math.Sqrt(float64(n)))
	for i := 3; i <= maxDivisor; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func firstDivisor(n int) int {
	if n <= 1 {
		return 1
	}
	maxDivisor := int(math.Sqrt(float64(n)))
	for i := 2; i <= maxDivisor; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}
