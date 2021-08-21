package game

var ordinaryResources []string = []string{"Leaves", "Berry"}
var raresResources []string = []string{"Honey", "Corpse", "Leftovers"}
var resourceTypeListNames [][]string = [][]string{ordinaryResources, raresResources}

var capacityDefault int = 40

type resourceType struct {
	Name        string
	Size        int // 1 - small, 3 - big
	CapacityMax int
	Capacity    int
}

type Resource struct {
	Type     resourceType
	ID       int
	OwnerID  int
	Cell     Cell
	IsMining bool
}

func newResource(id int, name string) Resource {
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
		IsBase:   false,
		IsResrc:  true,
		toString: "R",
		points:   [2]int{},
	}

	res := Resource{
		Type: resourceType{
			Name:        name,
			Size:        resourceSize,
			CapacityMax: capacity,
			Capacity:    capacity,
		},
		ID:       id,
		OwnerID:  0,
		Cell:     cell,
		IsMining: false,
	}

	return res
}
