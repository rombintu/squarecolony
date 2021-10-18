package game_test

import (
	"fmt"
	"testing"

	"github.com/rombintu/square_colony/game"
)

func TestRand(t *testing.T) {
	g := game.NewGame()
	fmt.Println(g.Config.Gameplay.SizeField)
	fmt.Println(game.Randint(g.Config.Gameplay.SizeField[0], g.Config.Gameplay.SizeField[1]))
	fmt.Println(game.Pick([]string{"a", "b", "c"}))
}
