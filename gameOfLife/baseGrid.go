package main

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

// CreateCell is a method that creates a new cell at the given address
// func (g *BaseGrid) CreateCell(addr *Address) *Cell {
// 	panic("Not implemented")
// }

// // createGrid is a method that creates a new grid of cells
// func (g *BaseGrid) createGrid() [][]*Cell {
// 	grid := make([][]*Cell, g.Rows)
// 	for row := 0; row < g.Rows; row++ {
// 		grid[row] = make([]*Cell, g.Cols)
// 		for col := 0; col < g.Cols; col++ {
// 			addr := NewAddress(row, col)
// 			grid[row][col] = g.CreateCell(addr)
// 		}
// 	}
// 	return grid
// }
