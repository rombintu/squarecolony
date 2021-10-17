package game

import (
	"github.com/rombintu/square_colony/utils"
	"github.com/sirupsen/logrus"
)

type Game struct {
	Battlefield Battlefield
	Logger      *logrus.Logger
	Config      *utils.Config
}

func NewGame() *Game {
	return &Game{
		Config: &utils.Config{},
	}
}

func (g *Game) ConfigLoger() {
	g.Logger = utils.NewLogger(g.Config.Debug.LogLevel)
}

func (g *Game) InitPlayers(playersInfoList [][]string) {

	for i, info := range playersInfoList {
		player := newPlayer(i+1, info)
		g.Battlefield.Players = append(g.Battlefield.Players, player)
		g.Battlefield.Metadata.CellsCount++
	}

}

func (g *Game) InitResources() {
	resourcesInfoList := RandomSelectFromArr(
		resourceTypeListNames,
		g.Config.Gameplay.CountResorces,
	)

	for i, info := range resourcesInfoList {
		resource := newResource(i+1, info)
		g.Battlefield.Resources = append(g.Battlefield.Resources, resource)
		g.Battlefield.Metadata.CellsCount++
	}
}

// Generate new point of cells
func (g *Game) InitPoints() {
	var cells []*Cell
	// var points [][2]int

	for i := 0; i < len(g.Battlefield.Players); i++ {
		cells = append(cells, &g.Battlefield.Players[i].Base.Cell)
	}

	for i := 0; i < len(g.Battlefield.Resources); i++ {
		cells = append(cells, &g.Battlefield.Resources[i].Cell)
	}

	for i, cell := range cells {
		cell.SetID(i + 1)
		cell.SetNewPoint(g.Config.Gameplay.SizeField, cells)
		// points = append(points, cell.GetPoint())
	}
}

func (g *Game) Init(p [][]string) {
	g.ConfigLoger()

	g.Logger.Info("INIT PLAYERS")
	g.InitPlayers(p)
	g.Logger.Info("INIT RESOURCES")
	g.InitResources()
	g.Logger.Info("INIT POINTS")
	g.InitPoints()
}

func (g *Game) GetPlayers() []Player {
	var arr []Player
	for _, player := range g.Battlefield.Players {
		arr = append(arr, *player)
	}
	return arr
}

func (g *Game) GetBases() []Base {
	var arr []Base
	for _, player := range g.Battlefield.Players {
		arr = append(arr, player.Base)
	}
	return arr
}

func (g *Game) GetResources() []*Resource {
	return g.Battlefield.Resources
}
