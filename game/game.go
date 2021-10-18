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
		Config: utils.NewConfig(),
	}
}

func (g *Game) ConfigLoger() {
	g.Logger = utils.NewLogger(g.Config.Debug.LogLevel)
}

func (g *Game) Init(p [][]string) {
	g.ConfigLoger()

	g.Logger.Info("INIT PLAYERS")
	g.Battlefield.InitPlayers(p)
	g.Logger.Info("INIT RESOURCES")
	g.Battlefield.InitResources(g.Config.Gameplay.CountResorces)
	g.Logger.Info("INIT POINTS")
	g.Battlefield.InitPoints(g.Config.Gameplay.SizeField)
}
