package main

import (
	"fmt"

	"github.com/rombintu/square_colony/game"
)

func buildGame() {
	// ========= CONFIG ========= //

	// Size buttlefield, recommend = [10, 20]
	var sizeField [2]int = [2]int{10, 20}
	// Count rescource pints
	var countResorces int = 3

	// ========= INIT ========= //

	// LOG := utils.NewLogger()
	playerNameList := [][]string{
		{"Player1", game.PlayerTypeListNames[0]},
		{"Player2", game.PlayerTypeListNames[2]},
	}

	bf := game.NewButtlefield(
		countResorces,
		playerNameList,
		sizeField,
	)

	// log.Debug(bf.Resources)
	// log.Println(baseList)
	// bf.ShowButtlefield()
	for i := 0; i < len(bf.Cells); i++ {
		println(i, fmt.Sprint(bf.Cells[i].Points))
	}

	// for i := 0; i < len(bf.PlayerList); i++ {
	// 	println(i, fmt.Sprint(bf.PlayerList))
	// }

	for i, player := range bf.PlayerList {
		println(i, fmt.Sprint(player))
	}

	for i, res := range bf.Resources {
		println(i, fmt.Sprint(res))
	}
}

func main() {
	buildGame()
}
