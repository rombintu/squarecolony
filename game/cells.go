package game

import "reflect"

// Set points of cells. Checking repeatability
func setPoints(cell *Cell, cells []*Cell, size [2]int) {
	point := [2]int{
		randint(1, size[0]),
		randint(1, size[1]),
	}
	for _, cell := range cells {
		if reflect.DeepEqual(cell.Points, point) {
			setPoints(cell, cells, size)
		}
	}

	cell.Points = point
}
