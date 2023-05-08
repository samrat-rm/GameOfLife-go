package main

import (
	"testing"
)

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
