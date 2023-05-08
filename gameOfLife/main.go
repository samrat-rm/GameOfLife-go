package main

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

func main() {
	grid := NewBaseGrid(10, 10)
	grid.createGrid()
	PrintCurrentGrid(grid)
	grid.UpdateGrid()
	PrintCurrentGrid(grid)
}
