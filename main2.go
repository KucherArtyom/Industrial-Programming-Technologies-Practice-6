package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var num1, num2 int

	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)

	result := gcd(num1, num2)

	fmt.Printf("Наибольший общий делитель чисел %d и %d: %d\n", num1, num2, result)
}
