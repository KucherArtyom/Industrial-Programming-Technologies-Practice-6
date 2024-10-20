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
