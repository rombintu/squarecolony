package game

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func randint(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min) + min
	return r
}

func pick(arr []string) string {
	pickItem := arr[randint(0, len(arr))]
	return pickItem
}

// Return array with random element from other array
func RandomSelectFromArr(arr [][]string, size int) []string {
	var newArr []string
	for i := 0; i < size; i++ {
		randn := randint(0, 100)

		if randn > 25 {
			newArr = append(newArr, pick(arr[0]))
		} else {
			newArr = append(newArr, pick(arr[1]))
		}
	}
	return newArr
}

// func Colorite(messsage interface{}, c string) {
// 	println(color.Ize(c, fmt.Sprint(messsage)))
// }

// // Print cells
// func cellToString(cell *Cell) string {
// 	if cell.IsBase {
// 		return ("[B]")
// 	} else if cell.IsResrc {
// 		return ("[R]")
// 	}
// 	return "[ ]"
// }

// Return distance between two points
func DistBetweenPoints(f, d [2]int) int {
	return int(math.Abs(float64((d[0] + d[1]) - (f[0] + f[1]))))
}

// Return distance between two cells
func DistBetweenCells(c1, c2 Cell) int {
	return DistBetweenPoints(c1.points, c2.points)
}

// Return resource the closest to base
func NearestResource(
	base Base,
	resources []Resource,
) Resource {
	mainCell := base.Cell
	buff := make(map[int]Resource)
	keys := make([]int, len(resources))
	for i := 0; i < len(resources); i++ {
		p := DistBetweenCells(mainCell, resources[i].Cell)
		buff[p] = resources[i]
		keys[i] = p
	}

	sort.Ints(keys)
	fmt.Println(keys)
	return buff[keys[0]]
}

// func NearestResources(
// 	base Base,
// 	resources []Resource,
// 	count int,
// ) []Resource {
// 	if count < len(resources) {
// 		count = len(resources)
// 	}

// 	points := base.Cell.points

// 	var buff []Resource

// 	for i := 0; i < resources; i++ {
// 		//
// 	}
// }
