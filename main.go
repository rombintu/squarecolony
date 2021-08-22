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
	var cells []*game.Cell
	var points [][2]int

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

	for i := 0; i < len(bf.PlayerList); i++ {
		cells = append(cells, &bf.PlayerList[i].Base.Cell)
	}

	for i := 0; i < len(bf.Resources); i++ {
		cells = append(cells, &bf.Resources[i].Cell)
	}

	for i, cell := range cells {
		cell.SetID(i + 1)
		cell.SetNewPoint(sizeField, cells)
		points = append(points, cell.GetPoint())
	}

	fmt.Println(bf)

	// log.Debug(bf.Resources)
	// log.Println(baseList)
	// bf.ShowButtlefield(points)

	// for i, player := range bf.PlayerList {
	// 	println(i, fmt.Sprint(player))
	// }
}

func main() {
	buildGame()
}
