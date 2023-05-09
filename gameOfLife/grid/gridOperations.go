package grid

import (
	"fmt"
	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/address"
	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/cell"
)

type GridOperations interface {
	BaseGridInterface
	UpdateGrid() [][]*cell.Cell
	GetNeighbors(address *address.Address) []*cell.Cell
}

func (g *Grid) UpdateGrid() [][]*cell.Cell {
	newGrid := make([][]*cell.Cell, g.Rows)
	for row := 0; row < g.Rows; row++ {
		newRow := make([]*cell.Cell, g.Cols)
		for col := 0; col < g.Cols; col++ {
			currentCell := g.GetCell(row, col)
			cellAddress := currentCell.Address
			neighbors := g.GetNeighbors(cellAddress)
			state := currentCell.UpdateState(neighbors)
			var updatedCell *cell.Cell
			if state {
				updatedCell, _ = cell.NewCell(cellAddress, true)
			} else {
				updatedCell, _ = cell.NewCell(cellAddress, false)
			}
			fmt.Println(state, cellAddress.Row, cellAddress.Col)
			newRow[col] = updatedCell
		}
		newGrid[row] = newRow
	}
	PrintCurrentGrid(newGrid)
	g.Grid = newGrid
	return newGrid
}

func (g *Grid) GetNeighbors(address *address.Address) []*cell.Cell {
	row := address.Row
	col := address.Col
	neighbors := make([]*cell.Cell, 0)
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < g.Rows && c >= 0 && c < g.Cols && !(r == row && c == col) {
				neighbors = append(neighbors, g.GetCell(r, c))
			}
		}
	}
	return neighbors
}
