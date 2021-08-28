package game

// ========= INTERNAL CONFIG ========= //

var capacityBaseDefault int = 10000

// ========= INIT ========= //

// Players base
type Base struct {
	ID          int
	OwnerID     int
	CapacityMax int
	Capacity    int
	Name        string
	Cell        Cell
}

// Create new Base{...}
func newBase(id int, name string) Base {
	cell := Cell{
		IsBase:   true,
		IsResrc:  false,
		toString: "B",
		points:   [2]int{},
	}

	base := Base{
		id,
		id,
		capacityBaseDefault,
		capacityBaseDefault / 100 * 10,
		name,
		cell,
	}
	return base
}

func (b *Base) GetCell() Cell {
	return b.Cell
}
