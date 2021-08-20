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
	sizeField [2]int,
) (Player, Cell) {
	base, cell := newBase(id, name+"/Base", sizeField)
	return Player{id, name, base}, cell
}
