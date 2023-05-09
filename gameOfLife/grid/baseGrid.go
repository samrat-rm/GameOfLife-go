package grid

import (
	"math/rand"
	"time"

	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/address"
	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/cell"
)

type Grid interface {
	GetCell(row, col int) *cell.Cell
	CreateCell(addr *address.Address) *cell.Cell
	createGrid() [][]*cell.Cell
}

type BaseGrid struct {
	Rows int
	Cols int
	Grid [][]*cell.Cell
}

// NewBaseGrid is a constructor function that creates a new BaseGrid
func NewBaseGrid(rows, cols int) *BaseGrid {
	grid := make([][]*cell.Cell, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]*cell.Cell, cols)
		for col := 0; col < cols; col++ {
			addr := address.NewAddress(row, col)
			grid[row][col], _ = cell.NewCell(addr, false)
		}
	}
	return &BaseGrid{Rows: rows, Cols: cols, Grid: grid}
}

// GetCell is a method that returns the cell at the given row and column
func (g *BaseGrid) GetCell(row, col int) *cell.Cell {
	if row >= 0 && row < g.Rows && col >= 0 && col < g.Cols {
		return g.Grid[row][col]
	} else {
		return nil
	}
}

// CreateCell is a method that creates a new cell at the given address
func (g *BaseGrid) CreateCell(addr *address.Address) *cell.Cell {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	// Generate a random float64 between 0 and 1
	randomFloat := rng.Float64()
	// fmt.Println(randomFloat)
	state := randomFloat < 0.5 // Randomly assign true or false with 50% probability
	cell, _ := cell.NewCell(addr, state)
	return cell
}

// createGrid is a method that creates a new grid of cells with random states
func (g *BaseGrid) CreateGrid() [][]*cell.Cell {

	grid := make([][]*cell.Cell, g.Rows)
	for row := 0; row < g.Rows; row++ {
		grid[row] = make([]*cell.Cell, g.Cols)
		for col := 0; col < g.Cols; col++ {
			addr := address.NewAddress(row, col)
			grid[row][col] = g.CreateCell(addr)
		}
	}
	g.Grid = grid
	return grid
}
