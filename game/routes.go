package game

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct {
	Message string `json:"message"`
}

func (s *Server) Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		pong := Ping{Message: "pong"}
		c.JSON(http.StatusOK, pong)
	}
}

func (s *Server) ShowMap() gin.HandlerFunc {
	return func(c *gin.Context) {
		nonePrivateInfo := struct {
			Resources []Resource
			Bases     []Base
			SizeField [2]int
		}{
			Resources: s.Game.Battlefield.GetResources(),
			Bases:     s.Game.Battlefield.GetBases(),
			SizeField: s.Game.Battlefield.SizeField,
		}
		c.JSON(http.StatusOK, nonePrivateInfo)
	}
}

func (s *Server) UserAdd() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")
		pass := c.Query("pass")
		admin := c.Query("admin")
		if name == "" || pass == "" || admin == "" || admin != s.Game.Config.Server.Admin {
			c.JSON(http.StatusBadRequest, Player{})
		}
		if err := s.Game.Battlefield.AddPlayer(name, pass); err != nil {
			s.Logger.Info(err)
			c.JSON(http.StatusBadRequest, Player{})
		}
		s.Logger.Info(
			fmt.Sprintf("User %s was add", name),
			fmt.Sprintf("User %s was add. Password: %s", name, pass),
		)

		c.JSON(http.StatusOK, Player{Name: name, Password: pass})
	}
}
