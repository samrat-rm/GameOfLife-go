// address.go
package main

// Address is a struct that holds a row and column
type Address struct {
	Row int
	Col int
}

// NewAddress is a constructor function that creates a new Address
func NewAddress(row, col int) *Address {
	return &Address{Row: row, Col: col}
}
