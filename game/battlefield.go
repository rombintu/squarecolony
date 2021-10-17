package game

type Metadata struct {
	CellsCount int
	Size       [2]int
}
type Battlefield struct {
	Players   []*Player
	Resources []*Resource
	Metadata  Metadata
}

func (bf *Battlefield) InitPlayers(playersInfoList [][]string) {

	for i, info := range playersInfoList {
		player := newPlayer(i+1, info)
		bf.Players = append(bf.Players, player)
		bf.Metadata.CellsCount++
	}

}

func (bf *Battlefield) InitResources() {
	resourcesInfoList := RandomSelectFromArr(
		resourceTypeListNames,
		bf.Config.Gameplay.CountResorces,
	)

	for i, info := range resourcesInfoList {
		resource := newResource(i+1, info)
		bf.Data.Resources = append(bf.Data.Resources, resource)
		bf.Data.Metadata.CellsCount++
	}
}

// Generate new point of cells
func (bf *Battlefield) InitPoints() {
	var cells []*Cell
	// var points [][2]int

	for i := 0; i < len(bf.Data.Players); i++ {
		cells = append(cells, &bf.Data.Players[i].Base.Cell)
	}

	for i := 0; i < len(bf.Data.Resources); i++ {
		cells = append(cells, &bf.Data.Resources[i].Cell)
	}

	for i, cell := range cells {
		cell.SetID(i + 1)
		cell.SetNewPoint(bf.Config.Gameplay.SizeField, cells)
		// points = append(points, cell.GetPoint())
	}
}

func (bf *Battlefield) Init(p [][]string) {
	bf.InitPlayers(p)
	bf.InitResources()
	bf.InitPoints()
}

// Show main battlefield in console
// func (bf *Battlefield) ShowBattlefield(points [][2]int) {

// 	print("0  | ")

// 	for y0 := 1; y0 < bf.size[1]+1; y0++ {
// 		if y0 < 10 {
// 			fmt.Printf(" %d ", y0)
// 		} else {
// 			fmt.Printf("%d ", y0)
// 		}
// 	}

// 	print("\n")
// 	fmt.Print(strings.Repeat("-+-", bf.size[1]+2))
// 	print("\n")
// 	for x := 1; x < bf.size[0]+1; x++ {
// 		if x < 10 {
// 			fmt.Printf("%d  | ", x)
// 		} else {
// 			fmt.Printf("%d | ", x)
// 		}
// 		for y := 1; y < bf.size[1]+1; y++ {
// 			// TODO
// 			// for _, p := range points {
// 			// 	if reflect.DeepEqual([2]int{x, y}, p) {
// 			// 		fmt.Print("*")
// 			// 	}
// 			// }
// 		}
// 	}
// }

func (bf *Battlefield) GetInfo() {
	// log.
}

func (bf *Battlefield) GetPlayers() []*Player {
	return bf.Data.Players
}

func (bf *Battlefield) GetBases() []Base {
	var arr []Base
	for _, player := range bf.Data.Players {
		arr = append(arr, player.Base)
	}
	return arr
}

func (bf *Battlefield) GetResources() []*Resource {
	return bf.Data.Resources
}
