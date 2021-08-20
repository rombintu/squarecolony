package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/TwinProduction/go-color"
)

func randomizer(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min) + min
	return r
}

func pick(arr []string) string {
	pickItem := arr[randomizer(0, len(arr))]
	return pickItem
}

// Функция возвращает новый массив заполненый ресурсами по редкости
func RandomSelectFromArr(arr [][]string, size int) []string {
	var newArr []string
	for i := 0; i < size; i++ {
		randn := randomizer(0, 100)

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
