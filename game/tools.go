package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwinProduction/go-color"
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

func Colorite(messsage interface{}, c string) {
	println(color.Ize(c, fmt.Sprint(messsage)))
}

// Print cells
func cellToString(cell *Cell) string {
	if cell.isBase {
		return ("[B]")
	} else if cell.isResrc {
		return ("[R]")
	}
	return "[ ]"
}
