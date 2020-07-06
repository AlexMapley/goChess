package main

// CreateBoard will return an empty 8 by 8 board
func CreateBoard(xMax int, yMax int) map[Tile]rune {
	grid := map[Tile]rune{}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			tile := Tile{
				X: x,
				Y: y,
			}
			grid[tile] = '.'
		}
	}
	return grid
}