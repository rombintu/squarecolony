package game

var ClassTypeList []string = []string{
	"Larius", "Red", "Polyergus",
	"Myrmica", "Formica", "Pharaoh",
}

type PlayerClass struct {
	Name string
}

type Player struct {
	Id   int
	Name string
	Base Base
	// Type    Type
}

func newPlayer(
	id int,
	name string,
) (Player, *Cell) {
	base, cell := newBase(id, name+"/Base")
	return Player{id, name, base}, cell
}
