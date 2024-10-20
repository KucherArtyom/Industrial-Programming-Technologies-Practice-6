package main

import "fmt"

var memo = map[int]int{}

func fibonacci(n int) int {
	if val, found := memo[n]; found {
		return val
	}

	if n <= 1 {
		return n
	}

	memo[n] = fibonacci(n-1) + fibonacci(n-2)
	return memo[n]
}

func main() {
	var n int
	fmt.Print("Введите число для вычисления Фибоначчи: ")
	fmt.Scan(&n)

	result := fibonacci(n)
	fmt.Printf("Число Фибоначчи для n = %d: %d\n", n, result)
}
