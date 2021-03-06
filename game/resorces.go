package game

var ordinaryResources []string = []string{"Leaves", "Berry", "Corpses"}
var raresResources []string = []string{"Honey", "Leftovers"}
var resourceTypeListNames [][]string = [][]string{ordinaryResources, raresResources}

// CONFIG VARS
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

func newResource(id int, name string) *Resource {
	resourceSize := Randint(1, 4)
	var capacity int
	switch resourceSize {
	case 1:
		capacity = Randint(capacityDefault, capacityDefault*2)
	case 2:
		capacity = Randint(capacityDefault*3, capacityDefault*4)
	case 3:
		capacity = Randint(capacityDefault*5, capacityDefault*6)
	}

	cell := Cell{
		IsBase:   false,
		IsResrc:  true,
		toString: "Rr",
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

	return &res
}

func (r *Resource) GetCell() Cell {
	return r.Cell
}

// Change resource status {bool}
func (r *Resource) ChangeStatus() {
	if r.IsMining {
		r.IsMining = false
	} else {
		r.IsMining = true
	}
}

// Get owner id
func (r *Resource) GetOwnerId() int {
	return r.OwnerID
}

// Set owner id
func (r *Resource) SetOwnerId(newID int) {
	r.OwnerID = newID
}
