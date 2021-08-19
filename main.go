package main

import (
	"log"

	game "github.com/rombintu/square_colony/gamefunc"
)

func main() {
	var playerList []game.Player
	var sizeField [2]int = [2]int{100, 100}

	bf := game.NewButtlefield(playerList, sizeField)
	bf.ShowButtlefield()
	bf.ShowPlayerList()
	log.Println(bf.Resources)
}
