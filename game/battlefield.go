package game

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Metadata struct {
	CellsCount int
	Cells      []*Cell
}

type Battlefield struct {
	Players   []*Player
	Resources []*Resource
	Metadata  Metadata
	SizeField [2]int
}

func (bf *Battlefield) AddPlayer(name string) {
	id := uuid.NewV4().String()
	player := newPlayer(id, name)
	bf.Players = append(bf.Players, player)

	player.Base.Cell.SetNewPoint(bf.SizeField, bf.Metadata.Cells)

	bf.Metadata.Cells = append(bf.Metadata.Cells, &player.Base.Cell)
	bf.Metadata.CellsCount++

}

// func (bf *Battlefield) InitPlayers(playersInfoList [][]string) {

// 	for i, info := range playersInfoList {
// 		player := newPlayer(i+1, info)
// 		bf.Players = append(bf.Players, player)
// 		bf.Metadata.Cells = append(bf.Metadata.Cells, player.Base.Cell)
// 		bf.Metadata.CellsCount++
// 	}

// }

func (bf *Battlefield) InitResources(CountResorces int) {
	resourcesInfoList := RandomSelectFromArr(
		resourceTypeListNames,
		CountResorces,
	)

	for i, info := range resourcesInfoList {
		resource := newResource(i+1, info)
		bf.Resources = append(bf.Resources, resource)
		bf.Metadata.Cells = append(bf.Metadata.Cells, &resource.Cell)
		bf.Metadata.CellsCount++
	}
}

// Generate new point of cells
func (bf *Battlefield) RefreshPoints(SizeField [2]int) {
	var cells []*Cell

	for i := 0; i < len(bf.Players); i++ {
		cells = append(cells, &bf.Players[i].Base.Cell)
	}

	for i := 0; i < len(bf.Resources); i++ {
		cells = append(cells, &bf.Resources[i].Cell)
	}

	for i, cell := range cells {
		cell.SetID(i + 1)
		cell.SetNewPoint(SizeField, cells)
	}
}

func ReturnCell(cell Cell) string {
	switch {
	case cell.IsBase:
		return cell.toString
	case cell.IsResrc:
		return cell.toString
	default:
		return "[ ]"
	}
}

// Show main battlefield in console
func (bf *Battlefield) ShowBF(size [2]int) {
	fmt.Print("[ .]")
	for h := 0; h < size[1]; h++ {
		fmt.Printf("[%2.d]", h+1)

	}
	fmt.Print("\n")
	for x := 0; x < size[0]; x++ {
		if x == 0 {
			continue
		}

		fmt.Printf("[%2.d]", x)
		for y := 0; y < size[1]; y++ {
			var pointsA [2]int = [2]int{x, y}
			continFlag := false
			for _, c := range bf.Metadata.Cells {
				pointsB := c.points
				if pointsB == pointsA {
					if c.IsBase {
						fmt.Print("[🐜]")
					} else if c.IsResrc {
						fmt.Print("[🍫]")
					}
					continFlag = true
					continue
				}

			}
			if continFlag {
				continue
			}
			fmt.Print("[  ]")
		}
		fmt.Print("\n")
	}
}

func (bf *Battlefield) GetInfo() {
	// log.
}

func (bf *Battlefield) GetPlayers() []*Player {
	return bf.Players
}

func (bf *Battlefield) GetBases() []Base {
	var arr []Base
	for _, player := range bf.Players {
		arr = append(arr, player.Base)
	}
	return arr
}

func (bf *Battlefield) GetResources() []Resource {
	var arr []Resource
	for _, resource := range bf.Resources {
		arr = append(arr, *resource)
	}
	return arr
}
