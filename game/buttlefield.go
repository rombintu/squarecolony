package game

import (
	"log"
)

type Cell struct {
	ID      int
	isBase  bool
	isResrc bool
	Points  [2]int
}

type Buttlefield struct {
	PlayerList []Player
	BaseList   []Base
	Resources  []Resource
	Cells      []Cell
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
	var baseList []Base
	var resources []Resource
	var cells []Cell

	for i, name := range playerNameList {
		player, cell := newPlayer(i+1, name[0]+"/"+name[1], sizeField)
		baseList = append(baseList, player.Base)
		playerList = append(playerList, player)
		cells = append(cells, cell)
	}

	for i, name := range resourcesList {
		resource, cell := newResource(i+1, name, sizeField)
		resources = append(resources, resource)
		cells = append(cells, cell)
	}

	return Buttlefield{
		playerList,
		baseList,
		resources,
		cells,
		sizeField,
	}
}

// Show main buttlefield in console
func (bf *Buttlefield) ShowButtlefield() {
	//
}

func (bf *Buttlefield) ShowButtlefieldInfo() {
	log.Println("BF SIZE:", bf.Size)
}

func (bf *Buttlefield) ShowPlayerList() {
	log.Println("PLAYERS:", bf.PlayerList)
}
