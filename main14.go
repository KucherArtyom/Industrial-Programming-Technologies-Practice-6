package main

import (
	"fmt"
)

func digitalRoot(n int) int {
	if n < 10 {
		return n
	}

	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return digitalRoot(sum)
}

func main() {
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	root := digitalRoot(num)

	fmt.Printf("Цифровой корень числа: %d\n", root)
}
