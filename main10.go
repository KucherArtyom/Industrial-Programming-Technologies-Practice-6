package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	const maxAttempts = 10

	fmt.Println("Я загадал число от 1 до 100. Попробуйте угадать его!")
	fmt.Printf("У вас есть %d попыток.\n", maxAttempts)

	var guess int

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		fmt.Printf("Попытка %d: Введите ваше число: ", attempts)
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Загаданное число больше.")
		} else if guess > target {
			fmt.Println("Загаданное число меньше.")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d с %d попытки.\n", target, attempts)
			return
		}
	}

	fmt.Printf("Вы не смогли угадать число %d. Попытки закончились.\n", target)
}
