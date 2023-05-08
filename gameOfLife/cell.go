package main

import "fmt"

type Cell struct {
	Address *Address
	State   bool
}

// NewCell is a constructor function that creates a new Cell
func NewCell(addr *Address, state bool) (*Cell, error) {
	if addr == nil {
		return nil, fmt.Errorf("invalid parameters for cell class")
	}
	return &Cell{Address: addr, State: state}, nil
}

// UpdateState is a method that updates the state of a Cell based on its neighbors
func (c *Cell) UpdateState(neighbors []*Cell) error {
	// TODO: implement updateState method
	return fmt.Errorf("not implemented")
}
