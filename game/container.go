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

type ContainerStatuser interface {
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

// Get container status
func (c *Container) GetStatus() string {
	return c.Status
}

// Set container status
func (c *Container) SetStatus() string {
	return c.Status
}

// Return container statistic
func (c *Container) GetStats() map[string]int {
	return map[string]int{
		"hp":          c.HealthPoints,
		"power":       c.PowerPoints,
		"speed":       c.SpeedPoints,
		"capacity":    c.Capacity,
		"capacityMax": c.CapacityMax,
	}
}

// Return container points
func (c *Container) GetPoints() [2]int {
	return c.Points
}

// Get owner id
func (c *Container) GetOwnerId() int {
	return c.OwnerID
}

// Set owner id
func (c *Container) SetOwnerId(newID int) {
	c.OwnerID = newID
}
