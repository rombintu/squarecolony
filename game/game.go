package game

import (
	"fmt"

	"github.com/rombintu/square_colony/utils"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Game struct {
	Battlefield Battlefield
	Logger      *logrus.Logger
	Config      *utils.Config
}

func NewGame() *Game {
	config := utils.NewConfig()
	config.Server.Admin = uuid.NewV4().String()
	return &Game{
		Config: config,
	}
}

// func (g *Game) OpenLogFile() error {
// 	f, err := os.OpenFile(g.Config.Debug.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	g.LogFile = f
// 	return nil
// }

// func (g *Game) CloseLogFile() error {
// 	return g.LogFile.Close()
// }

func (g *Game) ConfigLoger() {
	g.Logger = logrus.New()
	level, err := logrus.ParseLevel(g.Config.Debug.LogLevel)
	if err != nil {
		g.Logger.Fatal(err)
	}
	g.Logger.SetLevel(level)
}

func (g *Game) ConfigGame() {
	g.Battlefield.SizeField = g.Config.Gameplay.SizeField
}

func (g *Game) Init() {
	g.ConfigLoger()
	g.ConfigGame()
	// g.Logger.Info("INIT PLAYERS")
	// g.Battlefield.InitPlayers(p)
	g.Logger.Info("INIT RESOURCES")
	g.Battlefield.InitResources(g.Config.Gameplay.CountResorces)
	g.Logger.Info("INIT POINTS")
	g.Battlefield.RefreshPoints(g.Config.Gameplay.SizeField)
	g.Logger.Info(fmt.Sprintf("Admin: %s", g.Config.Server.Admin))
}
