package grid

import (
	"github.com/samrat-rm/GameOfLife-go.git/gameOfLife/address"
	"testing"
)

func TestUpdateGrid(t *testing.T) {
	// Test 1: Any live cell with fewer than two live neighbors dies, as if caused by underpopulation.
	// Create a grid with a single live cell
	g1 := NewGrid(3, 3)
	cellAddr := address.NewAddress(1, 1)
	g1.CreateCell(cellAddr)
	g1.UpdateGrid()
	// Check that the cell dies due to underpopulation
	cell := g1.GetCell(1, 1)
	if cell.State {
		t.Errorf("Expected cell to die due to underpopulation, but it is still alive")
	}

	// Test 2: Any live cell with two or three live neighbors lives on to the next generation.
	// Create a grid with a live cell and two live neighbors
	g2 := NewGrid(3, 3)
	cellAddr = address.NewAddress(1, 1)
	g2.CreateCell(cellAddr)
	g2.CreateCell(address.NewAddress(0, 0))
	g2.GetCell(0, 0).State = true
	g2.CreateCell(address.NewAddress(0, 1))
	g2.GetCell(0, 1).State = true
	g2.CreateCell(address.NewAddress(0, 2))
	g2.GetCell(0, 2).State = true
	g2.UpdateGrid()
	// Check that the cell survives
	cell = g2.GetCell(1, 1)
	if !cell.State {
		t.Errorf("Expected cell to survive, but it is dead")
	}

	// Test 3: Any live cell with more than three live neighbors dies, as if by overpopulation.
	// Create a grid with a live cell and five live neighbors
	g3 := NewGrid(3, 3)
	cellAddr = address.NewAddress(1, 1)
	g3.CreateCell(cellAddr)
	g3.CreateCell(address.NewAddress(0, 0))
	g3.CreateCell(address.NewAddress(0, 1))
	g3.CreateCell(address.NewAddress(0, 2))
	g3.CreateCell(address.NewAddress(1, 0))
	g3.CreateCell(address.NewAddress(1, 2))
	g3.UpdateGrid()
	// Check that the cell dies due to overpopulation
	cell = g3.GetCell(1, 1)
	if cell.State {
		t.Errorf("Expected cell to die due to overpopulation, but it is still alive")
	}

	// Test 4: Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
	// Create a grid with a dead cell and three live neighbors
	g4 := NewGrid(3, 3)
	cellAddr = address.NewAddress(1, 1)
	g4.CreateCell(address.NewAddress(0, 0))
	g4.GetCell(0, 0).State = true
	g4.CreateCell(address.NewAddress(0, 1))
	g4.GetCell(0, 1).State = true
	g4.CreateCell(address.NewAddress(0, 2))
	g4.GetCell(0, 2).State = true
	g4.UpdateGrid()
	// Check that the cell is born due to reproduction
	cell = g4.GetCell(1, 1)
	if !cell.State {
		t.Errorf("Expected cell to be born due to reproduction, but it is still dead")
	}
}

// func TestUpdateGrid(t *testing.T) {
// 	baseGrid := NewGrid(5, 5)
// 	baseGrid.CreateGrid()
// 	oldGrid := baseGrid.Grid

// 	baseGrid.UpdateGrid()
// 	newGrid := baseGrid.Grid

// 	if oldGrid == nil {
// 		t.Error("Old grid should not be nil")
// 	}
// 	if newGrid == nil {
// 		t.Error("New grid should not be nil")
// 	}
// 	if len(oldGrid) != len(newGrid) {
// 		t.Error("Old grid and new grid should have the same number of rows")
// 	}
// 	if len(oldGrid[0]) != len(newGrid[0]) {
// 		t.Error("Old grid and new grid should have the same number of columns")
// 	}
// }
