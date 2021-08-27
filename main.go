package main

import (
	"fmt"

	"github.com/rombintu/square_colony/game"
)

func buildGame(
	players [][]string,
	size [2]int,
	countResorces int,
) *game.Buttlefield {

	bf := game.NewButtlefield(
		countResorces,
		players,
		size,
	)

	bf.Init(size)

	return &bf
}

func main() {
	// ========= CONFIG ========= //

	// Size buttlefield, recommend = [10, 20]
	var sizeField [2]int = [2]int{10, 20}
	// Count rescource points
	var countResorces int = 3

	// ========= INIT ========= //

	// LOG := utils.NewLogger()
	playerNameList := [][]string{
		{"Player1", game.PlayerTypeListNames[0]},
		{"Player2", game.PlayerTypeListNames[2]},
	}

	bf := buildGame(playerNameList, sizeField, countResorces)
	// log.Debug(bf.Resources)
	// log.Println(baseList)
	// bf.ShowButtlefield(points)

	// for i, player := range bf.PlayerList {
	// 	println(i, fmt.Sprint(player))
	// }

	// fmt.Println(bf.PlayerList)
	base := bf.GetBases()[0]
	ress := bf.GetResources()
	fmt.Println("BASE:", base)
	fmt.Println("RESOURCES:", ress)
	fmt.Println("NEAREST RESOURCE:", game.NearestResources(base, ress, -2))

}
