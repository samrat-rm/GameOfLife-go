package main

type GridOperations interface {
	UpdateGrid() [][]*Cell
	GetNeighbors(address *Address) []*Cell
}

func (g *BaseGrid) UpdateGrid() [][]*Cell {
	newGrid := make([][]*Cell, g.Rows)
	for row := 0; row < g.Rows; row++ {
		newRow := make([]*Cell, g.Cols)
		for col := 0; col < g.Cols; col++ {
			cell := g.GetCell(row, col)
			address := cell.Address
			neighbors := g.GetNeighbors(address)
			state := cell.UpdateState(neighbors)
			var updatedCell *Cell
			if state {
				updatedCell, _ = NewCell(address, true)
			} else {
				updatedCell, _ = NewCell(address, false)
			}
			newRow[col] = updatedCell
		}
		newGrid[row] = newRow
	}
	g.Grid = newGrid
	return newGrid
}

func (g *BaseGrid) GetNeighbors(address *Address) []*Cell {
	row := address.Row
	col := address.Col
	neighbors := make([]*Cell, 0)
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if r >= 0 && r < g.Rows && c >= 0 && c < g.Cols && !(r == row && c == col) {
				neighbors = append(neighbors, g.GetCell(r, c))
			}
		}
	}
	return neighbors
}
