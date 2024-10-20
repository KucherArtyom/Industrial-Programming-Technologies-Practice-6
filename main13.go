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
