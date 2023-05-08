package main

import (
	"testing"
)

func TestNewCellValid(t *testing.T) {
	addr := NewAddress(1, 2)
	cell, err := NewCell(addr, true)
	if err != nil {
		t.Errorf("NewCell returned error: %v", err)
	}
	if cell.Address != addr {
		t.Errorf("NewCell returned cell with wrong address: %v", cell.Address)
	}
	if cell.State != true {
		t.Errorf("NewCell returned cell with wrong state: %v", cell.State)
	}
}

func TestNewCellInvalid(t *testing.T) {
	_, err := NewCell(nil, true)
	if err == nil {
		t.Errorf("NewCell did not return error for nil address")
	}
}
