package game

import (
	"math"
	"math/rand"
	"time"
)

func Randint(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min) + min
	return r
}

func Pick(arr []string) string {
	pickItem := arr[Randint(0, len(arr))]
	return pickItem
}

// Return array with random element from other array
func RandomSelectFromArr(arr [][]string, size int) []string {
	var newArr []string
	for i := 0; i < size; i++ {
		randn := Randint(0, 100)

		if randn > 25 {
			newArr = append(newArr, Pick(arr[0]))
		} else {
			newArr = append(newArr, Pick(arr[1]))
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
	return int(math.Sqrt(
		math.Pow((float64(f[0]+f[1])), 2) +
			math.Pow(float64((d[0]+d[1])), 2)),
	)
}

// Return distance between two cells
func DistBetweenCells(c1, c2 Cell) int {
	return DistBetweenPoints(c1.points, c2.points)
}

// // Reurn all resource sorted of distance
// func GetAllResourcesSortedOfDist(
// 	base Base,
// 	resources []*Resource,
// ) []Resource {
// 	var sortedResources []Resource
// 	mainCell := base.Cell
// 	buff := make(map[int]Resource)
// 	keys := make([]int, len(resources))
// 	for i := 0; i < len(resources); i++ {
// 		p := DistBetweenCells(mainCell, resources[i].Cell)
// 		buff[p] = *resources[i]
// 		keys[i] = p
// 	}

// 	sort.Ints(keys)

// 	for _, key := range keys {
// 		sortedResources = append(sortedResources, buff[key])
// 	}

// 	return sortedResources
// }

// // Return one resource the closest to base
// func NearestResource(
// 	base Base,
// 	resources []*Resource,
// ) Resource {
// 	return GetAllResourcesSortedOfDist(
// 		base,
// 		resources,
// 	)[0]
// }

// // Return (count) resources the closest to base
// func NearestResources(
// 	base Base,
// 	resources []*Resource,
// 	count int,
// ) []Resource {
// 	arr := GetAllResourcesSortedOfDist(
// 		base,
// 		resources,
// 	)

// 	if count >= len(arr) {
// 		count = len(arr)
// 	} else if count <= 0 {
// 		count = 0
// 	}

// 	return arr[:count]
// }
