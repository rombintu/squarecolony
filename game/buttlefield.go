package game

import (
	"fmt"
	"log"
)

type Buttlefield struct {
	Id         int
	PlayerList []Player
	BaseList   []Base
	Size       [2]int
	Resources  Resources
}

func NewButtlefield(
	id int,
	playerList []Player,
	baseList []Base,
	sizeField [2]int,
	numResorce, capacity int,
) Buttlefield {
	resources := newResources(numResorce, capacity, sizeField)
	return Buttlefield{
		id,
		playerList,
		baseList,
		sizeField,
		resources,
	}
}

func (bf *Buttlefield) ShowButtlefield() {
	fmt.Println("==============")
}

func (bf *Buttlefield) ShowButtlefieldInfo() {
	log.Println("BF SIZE:", bf.Size)
}

func (bf *Buttlefield) ShowPlayerList() {
	log.Println("PLAYERS:", bf.PlayerList)
}
