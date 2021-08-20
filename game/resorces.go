package game

import (
	"math/rand"
	"time"
)

var resourceNameList []string = []string{
	"Leaves", "Berry", "Honey",
	"Corpse", "Leftovers",
}

type Resources struct {
	Id int
	// Name string
	Capacity int
	Owner    int
	Points   [][2]int
	isUsed   bool
}

func randomizer(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min) + min
	return r
}

func newResources(n, copacity int, sizeField [2]int) Resources {
	var points [][2]int
	var id int = rand.Intn(100)
	for i := 0; i <= n; i++ {
		var x int = randomizer(sizeField[0], sizeField[1])
		var y int = randomizer(sizeField[0], sizeField[1])
		points = append(points, [2]int{x, y})
	}
	return Resources{
		id,
		copacity,
		0,
		points,
		false,
	}
}
