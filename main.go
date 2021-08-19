package main

import (
	"log"

	"github.com/rombintu/square_colony/game"
)

func buildGame() {
	// CONFIG
	var sizeField [2]int = [2]int{-100, 100}
	var numResorce, capacityResource, capacityBase int = 5, 100, 300

	// INIT
	var id_bf, id_pl1 int = 1, 1

	playerOne := game.NewPlayer(
		id_pl1, "name1", "TeamName",
		[2]int{1, 1}, capacityBase,
	)
	var playerList []game.Player = []game.Player{playerOne}
	var baseList []game.Base

	for _, player := range playerList {
		baseList = append(baseList, player.Base)
	}

	bf := game.NewButtlefield(id_bf, playerList, baseList, sizeField, numResorce, capacityResource)

	bf.ShowButtlefieldInfo()
	bf.ShowPlayerList()
	log.Println(bf.Resources)
	log.Println(baseList)
}

func main() {
	buildGame()
}
