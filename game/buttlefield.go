package game

import (
	"log"
)

type Buttlefield struct {
	Id         int
	PlayerList []Player
	BaseList   []Base
	Size       [2]int
	Resources  []Resource
}

func NewButtlefield(
	id int,
	playerList []Player,
	baseList []Base,
	sizeField [2]int,
	numResorce, capacity int,
) Buttlefield {
	var resourcesListNames []string = RandomSelectFromArr(resourceTypeList, numResorce)
	var resources []Resource
	for i, resourceName := range resourcesListNames {
		resources = append(
			resources,
			newResources(i, resourceName, sizeField),
		)
	}
	return Buttlefield{
		id,
		playerList,
		baseList,
		sizeField,
		resources,
	}
}

func (bf *Buttlefield) ShowButtlefield() {
	// for i:=0; i < len(bf.Size); i++ {
	// 	// TODO
	// }
}

func (bf *Buttlefield) ShowButtlefieldInfo() {
	log.Println("BF SIZE:", bf.Size)
}

func (bf *Buttlefield) ShowPlayerList() {
	log.Println("PLAYERS:", bf.PlayerList)
}
