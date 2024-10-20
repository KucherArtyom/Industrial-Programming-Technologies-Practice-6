package main

import "fmt"

func Palindrome(n int) bool {
	if n < 0 {
		return false
	}
	original := n
	reversed := 0
	for n != 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	return original == reversed
}

func main() {
	var num int
	fmt.Print("Введите целое число: ")
	fmt.Scan(&num)
	if Palindrome(num) {
		fmt.Printf("Число %d является палиндромом.\n", num)
	} else {
		fmt.Printf("Число %d не является палиндромом.\n", num)
	}
}
