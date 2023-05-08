package main

import (
	"testing"
)

func TestNewBaseGrid(t *testing.T) {
	// Create a new BaseGrid with 2 rows and 3 columns
	grid := NewBaseGrid(2, 3)

	// Test the Rows and Cols fields
	if grid.Rows != 2 {
		t.Errorf("Expected Rows to be 2, but got %d", grid.Rows)
	}
	if grid.Cols != 3 {
		t.Errorf("Expected Cols to be 3, but got %d", grid.Cols)
	}
}

func TestGetCell(t *testing.T) {
	// Create a new BaseGrid with 2 rows and 3 columns
	grid := NewBaseGrid(2, 3)

	// Test getting a valid cell
	cell := grid.GetCell(0, 0)
	if cell == nil {
		t.Error("Expected cell to not be nil")
	}
	if cell.Address.Row != 0 {
		t.Errorf("Expected cell row to be 0, but got %d", cell.Address.Row)
	}
	if cell.Address.Col != 0 {
		t.Errorf("Expected cell col to be 0, but got %d", cell.Address.Col)
	}

	// Test getting a cell out of bounds
	cell = grid.GetCell(-1, 0)
	if cell != nil {
		t.Error("Expected cell to be nil")
	}
	cell = grid.GetCell(0, 3)
	if cell != nil {
		t.Error("Expected cell to be nil")
	}
}
func TestCreateGrid(t *testing.T) {
	baseGrid := NewBaseGrid(5, 5)
	baseGrid.createGrid()

	for row := 0; row < baseGrid.Rows; row++ {
		for col := 0; col < baseGrid.Cols; col++ {
			if baseGrid.Grid[row][col] == nil {
				t.Errorf("Cell at row %d and col %d should not be nil", row, col)
			}
		}
	}
}
func TestUpdateGrid(t *testing.T) {
	baseGrid := NewBaseGrid(5, 5)
	baseGrid.createGrid()
	oldGrid := baseGrid.Grid

	baseGrid.UpdateGrid()
	newGrid := baseGrid.Grid

	if oldGrid == nil {
		t.Error("Old grid should not be nil")
	}
	if newGrid == nil {
		t.Error("New grid should not be nil")
	}
	if len(oldGrid) != len(newGrid) {
		t.Error("Old grid and new grid should have the same number of rows")
	}
	if len(oldGrid[0]) != len(newGrid[0]) {
		t.Error("Old grid and new grid should have the same number of columns")
	}
}
