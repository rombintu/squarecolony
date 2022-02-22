package game

import (
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
