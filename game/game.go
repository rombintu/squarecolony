package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func (g *Game) ConfigGame() {
	g.Battlefield.SizeField = g.Config.Gameplay.SizeField
}

func (g *Game) HandleInput() {
	for {
		fmt.Print("ADMIN: ")
		in := bufio.NewReader(os.Stdin)
		line, err := in.ReadString('\n')
		if err != nil {
			g.Logger.Error(err)
		}
		g.caseInput(strings.Split(line, " "))
	}
}

func (g *Game) caseInput(input []string) {
	switch strings.TrimSpace(input[0]) {
	case "exit":
		g.Logger.Info("Game exiting...")
	case "help":
		g.Logger.Info("useradd <username> <class>")
	case "map":
		g.Battlefield.ShowBF(g.Config.Gameplay.SizeField)
	case "useradd":
		if len(input) != 3 {
			g.Logger.Error("Less values. Use: useradd <username> <class>")
			return
		}
		if err := g.Battlefield.AddPlayer(input[1], input[2]); err != nil {
			g.Logger.Error(err)
			return
		}
		g.Logger.Info(
			fmt.Sprintf("User %s was add. Password: %s", input[1], input[2]),
		)
	case "users":
		players := g.Battlefield.GetPlayers()
		for i, p := range players {
			g.Logger.Info(
				fmt.Sprintf("%d: %+v", i+1, p),
			)
		}

	case "resourses":
		resourses := g.Battlefield.GetResources()
		for i, r := range resourses {
			g.Logger.Info(
				fmt.Sprintf("%d. %+v", i+1, r),
			)
		}
	}
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

	g.HandleInput()
}
