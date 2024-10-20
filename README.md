# Industrial-Programming-Technologies-Practice-6 (ЭФМО-02-24 Кучер Артем Сергеевич)
## Задачи по языку программирования Go
### 1.Проверка на простоту
```
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
```
### 2.Наибольший общий делитель (НОД)
```
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
```
### 3.Сортировка пузырьком
```
package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				fmt.Printf("Шаг %d-%d: %v\n", i+1, j+1, arr)
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Введите количество элементов: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Println(i+1, "-й элемент массива:")
		fmt.Scan(&arr[i])
	}

	fmt.Println("Исходный массив:", arr)

	bubbleSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
```
### 4.Таблица умножения в формате матрицы
```
package main

import "fmt"

func main() {
	fmt.Println("Таблица умножения 10x10:")

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
```
### 5.Фибоначчи с мемоизацией
```
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
```
### 6.Обратные числа
```
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
```
### 7.Треугольник Паскаля
```
package main

import "fmt"

func pascalTriangle(levels int) {
	triangle := make([][]int, levels)

	for i := 0; i < levels; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
		fmt.Println(triangle[i])
	}
}

func main() {
	var levels int
	fmt.Print("Введите количество уровней треугольника Паскаля: ")
	fmt.Scan(&levels)

	pascalTriangle(levels)
}
```
### 8.Число палиндром
```
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
```
### 9.Нахождение максимума и минимума в массиве
```
package main

import (
	"fmt"
	"math"
)

func findMaxMin(arr []int) (int, int) {
	if len(arr) == 0 {
		fmt.Println("Массив пуст")
		return 0, 0
	}

	maxVal, minVal := math.MinInt64, math.MaxInt64

	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}

	return maxVal, minVal
}

func main() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&n)
	arr := make([]int, n)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d-й элемент массива: ", i+1)
		fmt.Scan(&arr[i])
	}
	maxVal, minVal := findMaxMin(arr)

	fmt.Printf("Максимум: %d, Минимум: %d\n", maxVal, minVal)
}
```
### 10.Игра "Угадай число"
```
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
```
### 11.Числа Армстронга
```
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
```
### 12.Подсчет слов в строке
```
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(strings.ToLower(input))

	words := strings.Fields(input)

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Printf("Количество уникальных слов: %d\n", len(wordCount))

	fmt.Println("Слова и их частоты:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
```
### 13.Игра "Жизнь" (Conway's Game of Life)
```
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const rows = 10
const cols = 10

func initializeGrid() [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
		for j := range grid[i] {
			grid[i][j] = rand.Intn(2)
		}
	}
	return grid
}

func countAliveNeighbors(grid [][]int, row, col int) int {
	aliveNeighbors := 0
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			aliveNeighbors += grid[newRow][newCol]
		}
	}

	return aliveNeighbors
}

func nextGeneration(grid [][]int) [][]int {
	newGrid := make([][]int, rows)
	for i := range grid {
		newGrid[i] = make([]int, cols)
		for j := range grid[i] {
			aliveNeighbors := countAliveNeighbors(grid, i, j)

			if grid[i][j] == 1 {
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					newGrid[i][j] = 1
				} else {
					newGrid[i][j] = 0
				}
			} else {
				if aliveNeighbors == 3 {
					newGrid[i][j] = 1
				} else {
					newGrid[i][j] = 0
				}
			}
		}
	}
	return newGrid
}

func printGrid(grid [][]int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	grid := initializeGrid()

	fmt.Println("Начальная генерация:")
	printGrid(grid)

	for generation := 1; generation <= 10; generation++ {
		fmt.Printf("Поколение %d:\n", generation)
		grid = nextGeneration(grid)
		printGrid(grid)
		time.Sleep(500 * time.Millisecond)
	}
}
```
### 14.Цифровой корень числа
```
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
```
### 15.Римские цифры
```
package main

import (
	"fmt"
)

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	return roman
}

func main() {
	var num int
	fmt.Print("Введите арабское число: ")
	fmt.Scan(&num)

	roman := intToRoman(num)

	fmt.Printf("Римское число: %s\n", roman)
}
```
