package main

import (
	"fmt"
	"math"
)

func isArmstrong(n int) bool {
	temp := n
	numDigits := 0
	for temp != 0 {
		numDigits++
		temp /= 10
	}

	sum := 0
	temp = n
	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	return sum == n
}

func main() {
	var num int
	fmt.Print("Введите число для проверки на число Армстронга: ")
	fmt.Scan(&num)

	if isArmstrong(num) {
		fmt.Printf("Число %d является числом Армстронга.\n", num)
	} else {
		fmt.Printf("Число %d не является числом Армстронга.\n", num)
	}
}
