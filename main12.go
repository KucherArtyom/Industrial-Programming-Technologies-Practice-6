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
