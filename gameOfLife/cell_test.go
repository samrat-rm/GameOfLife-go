package main

import (
	"testing"
)

func TestNewCell(t *testing.T) {
	addr := NewAddress(1, 1)
	cell, err := NewCell(addr, true)
	if err != nil {
		t.Errorf("Error creating new cell: %v", err)
	}

	if cell.Address != addr {
		t.Errorf("Expected address %v, got %v", addr, cell.Address)
	}

	if cell.State != true {
		t.Errorf("Expected state %v, got %v", true, cell.State)
	}
}

func TestUpdateState(t *testing.T) {
	addr1 := NewAddress(1, 1)
	addr2 := NewAddress(1, 2)
	addr3 := NewAddress(1, 3)

	c1, _ := NewCell(addr1, true)
	c2, _ := NewCell(addr2, true)
	c3, _ := NewCell(addr3, true)

	// Test for live cell with less than 2 live neighbors
	neighbors := []*Cell{c2}
	newState := c1.UpdateState(neighbors)
	if newState != false {
		t.Errorf("Expected false, got %v", newState)
	}

	// Test for live cell with 2 live neighbors
	neighbors = []*Cell{c2, c3}
	newState = c1.UpdateState(neighbors)
	if newState != true {
		t.Errorf("Expected true, got %v", newState)
	}

	// Test for live cell with 3 live neighbors
	neighbors = []*Cell{c1, c2, c3}
	newState = c2.UpdateState(neighbors)
	if newState != true {
		t.Errorf("Expected true, got %v", newState)
	}

	// Test for live cell with more than 3 live neighbors
	neighbors = []*Cell{c1, c2, c3, c1, c1}
	newState = c2.UpdateState(neighbors)
	if newState != false {
		t.Errorf("Expected false, got %v", newState)
	}

	// Test for dead cell with 3 live neighbors
	c4, _ := NewCell(NewAddress(2, 2), false)
	neighbors = []*Cell{c1, c2, c3}
	newState = c4.UpdateState(neighbors)
	if newState != true {
		t.Errorf("Expected true, got %v", newState)
	}

	// Test for dead cell with less than 3 live neighbors
	neighbors = []*Cell{c1}
	newState = c4.UpdateState(neighbors)
	if newState != false {
		t.Errorf("Expected false, got %v", newState)
	}

	// Test for dead cell with more than 3 live neighbors
	neighbors = []*Cell{c1, c2, c3, c1, c1}
	newState = c4.UpdateState(neighbors)
	if newState != false {
		t.Errorf("Expected false, got %v", newState)
	}
}
