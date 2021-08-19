package game

type Player struct {
	Id       int
	Name     string
	TeamName string
	Points   [2]int
	Base     Base
}

func NewPlayer(
	id int,
	name string,
	teamName string,
	points [2]int,
	capacityBase int,
) Player {
	return Player{id, name, teamName, points,
		newBase(id, id, capacityBase, name+"/base"),
	}
}
