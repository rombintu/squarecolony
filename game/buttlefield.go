package game

import (
	"fmt"
	"strings"
)

// Main object
type Buttlefield struct {
	PlayerList []Player
	// BaseList   []*Base
	Resources  []Resource
	CellsCount int
	Size       [2]int
}

// Create new ButtleField of game
func NewButtlefield(
	countResorces int,
	playersInfoList [][]string,
	sizeField [2]int,
) Buttlefield {
	resourcesInfoList := RandomSelectFromArr(
		resourceTypeListNames,
		countResorces,
	)

	var playerList []Player
	// var baseList []*Base
	var resources []Resource
	var cells int

	for i, info := range playersInfoList {
		player := newPlayer(i+1, info)
		playerList = append(playerList, player)
		// baseList = append(baseList, &player.Base)
		cells++
	}

	for i, info := range resourcesInfoList {
		resource := newResource(i+1, info)
		resources = append(resources, resource)
		cells++
	}

	return Buttlefield{
		PlayerList: playerList,
		// BaseList:   baseList,
		Resources:  resources,
		CellsCount: cells,
		Size:       sizeField,
	}
}

// Show main buttlefield in console
func (bf *Buttlefield) ShowButtlefield(points [][2]int) {

	print("0  | ")

	for y0 := 1; y0 < bf.Size[1]+1; y0++ {
		if y0 < 10 {
			fmt.Printf(" %d ", y0)
		} else {
			fmt.Printf("%d ", y0)
		}
	}

	print("\n")
	fmt.Print(strings.Repeat("-+-", bf.Size[1]+2))
	print("\n")
	for x := 1; x < bf.Size[0]+1; x++ {
		if x < 10 {
			fmt.Printf("%d  | ", x)
		} else {
			fmt.Printf("%d | ", x)
		}
		for y := 1; y < bf.Size[1]+1; y++ {
			// TODO
			// for _, p := range points {
			// 	if reflect.DeepEqual([2]int{x, y}, p) {
			// 		fmt.Print("*")
			// 	}
			// }
		}
	}
}

func (bf *Buttlefield) GetInfo() {
	// log.
}

func (bf *Buttlefield) GetPlayerList() []Player {
	return bf.PlayerList
}

func (bf *Buttlefield) GetBaseList() []Base {
	var arr []Base
	for _, player := range bf.PlayerList {
		arr = append(arr, player.Base)
	}
	return arr
}
