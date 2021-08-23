package game

import "time"

var ContainerNameList []string = []string{
	"Queen", "Worker", "Warrior", "Scout", "Warlord",
}

type Message struct {
	Data      []byte
	Timestamp time.Time
}

// Container | Squad
type Container struct {
	Name         string
	Type         string
	Status       string
	ID           int
	CapacityMax  int
	Capacity     int // default: 0
	OwnerID      int
	PowerPoints  int
	HealthPoints int
	SpeedPoints  int
	Points       [2]int
}

type ContainerStatus interface {
	Build() (Message, error)
	Prepare() (Message, error)
	Travel(from, dest [2]int) (Message, error)
	Mine() (Message, error)
	Defend() (Message, error)
	Chill() (Message, error)
}

func NewContainer(
	name, containerType string,
	id, capacity, ownerid int,
	hp, power, speed int,
	points [2]int,
) Container {
	return Container{
		Name:         name,
		Type:         containerType,
		Status:       "Building",
		ID:           id,
		CapacityMax:  capacity,
		OwnerID:      ownerid,
		PowerPoints:  power,
		HealthPoints: hp,
		SpeedPoints:  speed,
		Points:       points,
	}
}
