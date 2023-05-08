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
