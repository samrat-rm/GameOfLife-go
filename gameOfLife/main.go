package main

import "fmt"

func main() {
	grid := NewBaseGrid(3, 3)
	fmt.Println(grid.GetCell(1, 1).State)
}
