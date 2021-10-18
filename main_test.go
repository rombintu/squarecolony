package main_test

import (
	"fmt"
	"testing"

	"github.com/rombintu/square_colony/game"
)

func TestBuildGame(t *testing.T) {
	GAME := game.NewGame()
	playerNameList := [][]string{
		{"Player1", game.PlayerTypeListNames[0]},
		{"Player2", game.PlayerTypeListNames[2]},
	}
	fmt.Println(GAME.Config)
	GAME.Init(playerNameList)
	base := GAME.Battlefield.GetBases()[0]
	ress := GAME.Battlefield.GetResources()
	fmt.Println("BASE:", base)
	fmt.Println("RESOURCES:", ress)
	fmt.Println("NEAREST RESOURCE:", base.NearestResources(ress, 2))
}

func TestShowBF(t *testing.T) {
	GAME := game.NewGame()
	playerNameList := [][]string{
		{"Player1", game.PlayerTypeListNames[0]},
		{"Player2", game.PlayerTypeListNames[2]},
	}
	fmt.Println(GAME.Config)
	GAME.Init(playerNameList)
	GAME.Battlefield.ShowBF(GAME.Config.Gameplay.SizeField)
}
