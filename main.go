package main

import (
	"fmt"

	"github.com/rombintu/square_colony/game"
)

func buildGame() {
	// CONFIG
	var sizeField [2]int = [2]int{-10, 20}
	var numResorce int = 3

	// INIT
	// log := utils.NewLogger()
	playerNameList := [][]string{
		{"Player1", game.ClassTypeList[0]},
		{"Player2", game.ClassTypeList[2]},
	}

	bf := game.NewButtlefield(
		numResorce,
		playerNameList,
		sizeField,
	)

	// log.Debug(bf.Resources)
	// log.Println(baseList)
	// bf.ShowButtlefield()
	for i := 0; i < len(bf.Cells); i++ {
		println(i, fmt.Sprint(bf.Cells[i].Points))
	}

	println(fmt.Sprint(bf.PlayerList[0]))
}

func main() {
	buildGame()
}
