package game

import "time"

var squadTypeList []string = []string{
	"Queen", "Worker", "Warrior", "Scout", "Warlord",
}

type Message struct {
	Data      []byte
	Timestamp time.Time
}

type SquadTypeCore struct {
	Name        string
	CapacityMax int
	Capacity    int
}

type SquadTypeStatus interface {
	Build() ([]Message, error)
	Prepare() ([]Message, error)
	Travel(pointsFrom, pointsDest [2]int) ([]Message, error)
	Mine() ([]Message, error)
	Defend() ([]Message, error)
	Chill() ([]Message, error)
}
