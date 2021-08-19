package gamefunc

import (
	"log"
	"math/rand"
)

type Player struct {
	Id   int
	Name string
	Base []int
}

type Buttlefield struct {
	Id         int
	PlayerList []Player
	Size       [2]int
	Resources  Resources
}

type Resources struct {
	Id     int
	Points [][2]int
}

func NewButtlefield(playerList []Player, sizeField [2]int) Buttlefield {
	var id int = rand.Intn(100)
	resources := createResources(5)
	return Buttlefield{id, playerList, sizeField, resources}
}

func createResources(n int) Resources {
	var points [][2]int
	var id int = rand.Intn(100)
	for i := 0; i <= n; i++ {
		var x int = rand.Intn(100)
		var y int = rand.Intn(100)
		points = append(points, [2]int{x, y})
	}
	return Resources{id, points}
}

func (bf *Buttlefield) ShowButtlefield() {
	log.Println("BF SIZE:", bf.Size)
}

func (bf *Buttlefield) ShowPlayerList() {
	log.Println("PLAYERS:", bf.PlayerList)
}
