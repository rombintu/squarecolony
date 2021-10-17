package main_test

import (
	"testing"

	"github.com/rombintu/square_colony/game"
)

func TestBuildGame(t *testing.T) {
	GAME := game.NewGame()
	playerNameList := [][]string{
		{"Player1", game.PlayerTypeListNames[0]},
		{"Player2", game.PlayerTypeListNames[2]},
	}
	GAME.Init(playerNameList)
	// base := GAME.GetBases()[0]
	// ress := GAME.GetResources()
	// fmt.Println("BASE:", base)
	// fmt.Println("RESOURCES:", ress)
	// fmt.Println("NEAREST RESOURCE:", game.NearestResources(base, ress, 2))
}
