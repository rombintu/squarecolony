package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Store struct {
	Path   string
	Driver *sql.DB
}

type Server struct {
	Host   string
	Port   string
	Router *gin.Engine
	Store  *Store
	Logger *logrus.Logger
}

func (s *Store) Open() error {
	db, err := sql.Open("sqlite3", s.Path)
	if err != nil {
		return err
	}
	s.Driver = db
	return nil
}

func NewServer(host, port string) *Server {
	return &Server{
		Host:   host,
		Port:   port,
		Router: gin.Default(),
		Logger: logrus.New(),
	}
}

func (s *Server) Start() error {
	s.Logger.Info("Server starting...")
	s.ConfigureRouter()
	s.ConfigureStore()
	return http.ListenAndServe(
		fmt.Sprintf("%s:%s", s.Host, s.Port),
		s.Router,
	)
}

func (s *Server) ConfigureStore() {
	s.Store = &Store{
		Path: "./store.db",
	}
}

func (s *Server) ConfigureRouter() {
	s.Router.GET("/ping", s.Ping())
}
