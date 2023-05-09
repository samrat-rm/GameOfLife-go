package grid

import (
	"fmt"
	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/cell"
)

type PrintGridOperations interface {
	GridOperations
	PrintCurrentGrid()
}

func PrintCurrentGrid(g [][]*cell.Cell) {
	fmt.Println()

	rows := len(g)
	columns := len(g[0])
	var output string
	if rows > 5 {
		fmt.Println("  > ")
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			cell := g[row][col]
			if cell.State {
				output += " O "
			} else {
				output += " - "
			}
		}
		output += "\n \n"
	}

	fmt.Print(output)
}
