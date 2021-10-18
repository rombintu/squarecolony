package game

import "sort"

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

// Reurn all resource sorted of distance
func (b *Base) GetAllResourcesSortedOfDist(resources []Resource) []Resource {
	var sortedResources []Resource
	mainCell := b.Cell
	buff := make(map[int]Resource)
	keys := make([]int, len(resources))
	for i := 0; i < len(resources); i++ {
		p := DistBetweenCells(mainCell, resources[i].Cell)
		buff[p] = resources[i]
		keys[i] = p
	}

	sort.Ints(keys)

	for _, key := range keys {
		sortedResources = append(sortedResources, buff[key])
	}

	return sortedResources
}

// Return one resource the closest to base
func (b *Base) NearestResource(resources []Resource) Resource {
	return b.GetAllResourcesSortedOfDist(resources)[0]
}

// Return (count) resources the closest to base
func (b *Base) NearestResources(resources []Resource, count int) []Resource {
	arr := b.GetAllResourcesSortedOfDist(resources)

	if count >= len(arr) {
		count = len(arr)
	} else if count <= 0 {
		count = 0
	}

	return arr[:count]
}
