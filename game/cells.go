package game

import "reflect"

// Cell of Battlefield
type Cell struct {
	ID       int
	IsBase   bool
	IsResrc  bool
	toString string
	points   [2]int
}

func (c *Cell) SetID(ID int) {
	c.ID = ID
}

// Set points of cells. Checking repeatability
func (c *Cell) SetNewPoint(size [2]int, cells []*Cell) {
	point := [2]int{
		Randint(1, size[0]),
		Randint(1, size[1]),
	}
	for _, cell := range cells {
		if reflect.DeepEqual(cell.points, point) {
			c.SetNewPoint(size, cells)
		}
	}

	c.points = point
}
