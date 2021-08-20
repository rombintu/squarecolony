package game

var ClassNameList []string = []string{
	"Larius", "Red", "Polyergus",
	"Myrmica", "Formica", "Pharaoh",
}

type PlayerClass struct {
	Name string
}

type Player struct {
	Id       int
	Name     string
	TeamName string
	Points   [2]int
	Base     Base
	// Class    Class
}

func NewPlayer(
	id int,
	name, teamName string,
	points [2]int,
	capacityBase int,
) Player {
	return Player{id, name, teamName, points,
		newBase(id, id, capacityBase, name+"/base"),
	}
}
