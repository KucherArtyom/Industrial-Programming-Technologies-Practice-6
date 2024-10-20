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
