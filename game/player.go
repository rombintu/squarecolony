package game

import (
	"fmt"
)

var PlayerTypeListNames []string = []string{
	"Larius", "Red", "Polyergus",
	"Myrmica", "Formica", "Pharaoh",
}

type Player struct {
	ID      int
	Name    string
	Type    string
	Base    Base
	Bonuses []Bonus
}

func newPlayer(
	id int,
	playerInfo []string,
) Player {
	base := newBase(
		id,
		fmt.Sprintf(
			"[%s/%s]",
			playerInfo[0],
			playerInfo[1],
		),
	)
	return Player{
		ID:   id,
		Name: playerInfo[0],
		Type: playerInfo[1],
		Base: base,
	}
}
