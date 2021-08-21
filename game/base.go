package game

// CONFIG
var capacityBaseDefault int = 10000

type Base struct {
	Id          int
	OwnerID     int
	CapacityMax int
	Capacity    int
	Name        string
	Cell        Cell
}

// Create new Base{...}
func newBase(id int, name string) (Base, *Cell) {
	cell := Cell{
		isBase:   true,
		isResrc:  false,
		toString: "B",
		Points:   [2]int{},
	}

	return Base{
		id, id, capacityBaseDefault, 1000, name, cell,
	}, &cell
}
