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
	isDead  bool
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

func (p *Player) Dead() {
	p.isDead = true
}

func (p *Player) AddBonus(bonus Bonus) {
	p.Bonuses = append(p.Bonuses, bonus)
}
