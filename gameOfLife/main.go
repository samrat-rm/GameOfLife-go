package main

import (
	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/grid"
)

func main() {
	grid := grid.NewGrid(3, 3)
	grid.CreateGrid()
	grid.UpdateGrid()

}

// func PrintGridStates(grid [][]*Cell) {
// 	rows := len(grid)
// 	if rows == 0 {
// 		fmt.Println("The grid is empty.")
// 		return
// 	}
// 	columns := len(grid[0])

// 	fmt.Println("\n ")

// 	for row := 0; row < rows; row++ {
// 		for col := 0; col < columns; col++ {
// 			cellState := grid[row][col].State
// 			fmt.Printf("%t ", cellState)
// 		}
// 		fmt.Println()
// 	}

// }
