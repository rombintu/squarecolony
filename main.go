package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/rombintu/square_colony/game"
	"github.com/rombintu/square_colony/utils"
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
	var conf utils.Config

	// // ========= PARSE CONFIG ========= //
	confFile, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatalf("%v", err)
	}
	if err != nil {
		log.Fatalf("%v", err)
	}
	if _, err := toml.Decode(string(confFile), &conf); err != nil {
		log.Fatalf("%v", err)
	}

	// ========= INIT ========= //
	Logger := utils.NewLogger(conf.Debug.LogLevel)
	playerNameList := [][]string{
		{"Player1", game.PlayerTypeListNames[0]},
		{"Player2", game.PlayerTypeListNames[2]},
	}

	bf := buildGame(
		playerNameList,
		conf.Gameplay.SizeField,
		conf.Gameplay.CountResorces,
	)

	Logger.Debug("Create battlefield")
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
	fmt.Println("NEAREST RESOURCE:", game.NearestResources(base, ress, 2))

}
