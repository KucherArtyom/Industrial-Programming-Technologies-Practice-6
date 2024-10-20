package main

import "fmt"

func reverseNumber(n int) int {
	reversed := 0

	for n != 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}

	return reversed
}

func main() {
	var num int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)
	reversed := reverseNumber(num)
	fmt.Printf("Число в обратном порядке: %d\n", reversed)
}
