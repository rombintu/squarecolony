package game

import (
	"fmt"
)

var PlayerTypeListNames []string = []string{
	"Larius", "Red", "Polyergus",
	"Myrmica", "Formica", "Pharaoh",
}

type Player struct {
	ID       string
	Name     string
	Password string
	Type     string
	Base     Base
	Bonuses  []Bonus
	isDead   bool
}

func newPlayer(id, name, password string) *Player {
	base := newBase(
		id,
		fmt.Sprintf(
			"[%s#NoneType]",
			name,
		),
	)
	return &Player{
		ID:       id,
		Name:     name,
		Password: password,
		Base:     base,
	}
}

func (p *Player) Dead() {
	p.isDead = true
}

func (p *Player) AddBonus(bonus Bonus) {
	p.Bonuses = append(p.Bonuses, bonus)
}
