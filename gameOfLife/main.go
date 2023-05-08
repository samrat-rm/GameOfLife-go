package main

import "fmt"

func main() {
	// print the row and column of the address
	addr := NewAddress(1, 2)
	// print the row and column of the address
	fmt.Printf("Row: %d, Col: %d\n", addr.Row, addr.Col)
}
