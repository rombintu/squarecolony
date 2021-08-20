package game

var ordinaryResources []string = []string{"Leaves", "Berry"}
var raresResources []string = []string{"Honey", "Corpse", "Leftovers"}
var resourceTypeList [][]string = [][]string{ordinaryResources, raresResources}

var capacityDefault int = 40

type resourceType struct {
	Name     string
	Size     int // 1 - small, 3 - big
	Capacity int
}

type Resource struct {
	Type    resourceType
	ID      int
	OwnerID int
	// Points   [2]int
	Cell     Cell
	IsMining bool
}

func newResource(id int, name string, sizeField [2]int) (Resource, Cell) {
	var points [2]int = [2]int{
		randint(sizeField[0], sizeField[1]),
		randint(sizeField[0], sizeField[1]),
	}
	resourceSize := randint(1, 4)
	var capacity int
	switch resourceSize {
	case 1:
		capacity = randint(capacityDefault, capacityDefault*2)
	case 2:
		capacity = randint(capacityDefault*3, capacityDefault*4)
	case 3:
		capacity = randint(capacityDefault*5, capacityDefault*6)
	}

	cell := Cell{
		isBase:  false,
		isResrc: true,
		Points:  points,
	}

	return Resource{
		resourceType{
			Name:     name,
			Size:     resourceSize,
			Capacity: capacity,
		},
		id,
		0,
		cell,
		false,
	}, cell
}
