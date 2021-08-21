package game

import (
	"fmt"
	"log"
	"strings"
)

type Cell struct {
	ID       int
	isBase   bool
	isResrc  bool
	toString string
	Points   [2]int
}

type Buttlefield struct {
	PlayerList []Player
	BaseList   []*Base
	Resources  []Resource
	Cells      []*Cell
	Size       [2]int
}

// Create new ButtleField of game
func NewButtlefield(
	numResorce int,
	playerNameList [][]string,
	sizeField [2]int,
) Buttlefield {
	resourcesList := RandomSelectFromArr(resourceTypeList, numResorce)

	var playerList []Player
	var baseList []*Base
	var resources []Resource
	var cells []*Cell

	for i, name := range playerNameList {
		player, cell := newPlayer(i+1, name[0]+"/"+name[1])
		baseList = append(baseList, &player.Base)
		playerList = append(playerList, player)
		cells = append(cells, cell)
	}

	for i, name := range resourcesList {
		resource, cell := newResource(i+1, name)
		resources = append(resources, resource)
		cells = append(cells, cell)
	}

	for _, cell := range cells {
		setPoints(cell, cells, sizeField)
	}

	return Buttlefield{
		PlayerList: playerList,
		BaseList:   baseList,
		Resources:  resources,
		Cells:      cells,
		Size:       sizeField,
	}
}

// Show main buttlefield in console
func (bf *Buttlefield) ShowButtlefield() {

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
			fmt.Printf("%d  | \n", x)
		} else {
			fmt.Printf("%d | \n", x)
		}
		for y := 0; y < bf.Size[1]; y++ {
			// pass
		}
	}
}

func (bf *Buttlefield) getInfo() {
	// log.
}

func (bf *Buttlefield) getPlayerList() {
	log.Println("PLAYERS:", bf.PlayerList)
}
