package print

import (
	"fmt"

	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/grid"
)

func PrintCurrentGrid(g *grid.BaseGrid) {
	fmt.Println()

	rows := g.Rows
	columns := g.Cols
	var output string
	if rows > 5 {
		fmt.Println("  > ")
	}
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			cell := g.GetCell(row, col)
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
