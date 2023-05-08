package main

import (
	"math/rand"
	"time"
)

type BaseGrid struct {
	Rows int
	Cols int
	Grid [][]*Cell
}

// NewBaseGrid is a constructor function that creates a new BaseGrid
func NewBaseGrid(rows, cols int) *BaseGrid {
	grid := make([][]*Cell, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]*Cell, cols)
		for col := 0; col < cols; col++ {
			addr := NewAddress(row, col)
			grid[row][col], _ = NewCell(addr, false)
		}
	}
	return &BaseGrid{Rows: rows, Cols: cols, Grid: grid}
}

// GetCell is a method that returns the cell at the given row and column
func (g *BaseGrid) GetCell(row, col int) *Cell {
	if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
		return g.Grid[row][col]
	} else {
		return nil
	}
}

func (g *BaseGrid) UpdateGrid() [][]*Cell {
	newGrid := make([][]*Cell, g.Rows)
	for row := 0; row < g.Rows; row++ {
		newRow := make([]*Cell, g.Cols)
		for col := 0; col < g.Cols; col++ {
			cell := g.GetCell(row, col)
			address := cell.Address
			neighbors := g.GetNeighbors(address)
			state := cell.UpdateState(neighbors)
			var updatedCell *Cell
			if state {
				updatedCell, _ = NewCell(address, true)
			} else {
				updatedCell, _ = NewCell(address, false)
			}
			newRow[col] = updatedCell
		}
		newGrid[row] = newRow
	}
	g.Grid = newGrid
	return newGrid
}

func (g *BaseGrid) GetNeighbors(address *Address) []*Cell {
	row := address.Row
	col := address.Col
	neighbors := make([]*Cell, 0)
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < g.Rows && c >= 0 && c < g.Cols && !(r == row && c == col) {
				neighbors = append(neighbors, g.GetCell(r, c))
			}
		}
	}
	return neighbors
}

// CreateCell is a method that creates a new cell at the given address
func (g *BaseGrid) CreateCell(addr *Address) *Cell {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	// Generate a random float64 between 0 and 1
	randomFloat := rng.Float64()
	state := randomFloat < 0.5 // Randomly assign true or false with 50% probability
	cell, _ := NewCell(addr, state)
	return cell
}

// createGrid is a method that creates a new grid of cells with random states
func (g *BaseGrid) createGrid() [][]*Cell {

	grid := make([][]*Cell, g.Rows)
	for row := 0; row < g.Rows; row++ {
		grid[row] = make([]*Cell, g.Cols)
		for col := 0; col < g.Cols; col++ {
			addr := NewAddress(row, col)
			grid[row][col] = g.CreateCell(addr)
		}
	}
	g.Grid = grid
	return grid
}
