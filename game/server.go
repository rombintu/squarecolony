package game

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
	// LogFile *os.File
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

// func (s *Server) OpenLogFile() error {
// 	f, err := os.OpenFile("./server.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	s.LogFile = f
// 	return nil
// }

// func (s *Server) CloseLogFile() error {
// 	return s.LogFile.Close()
// }

func (s *Server) Start() error {
	// s.OpenLogFile()
	// defer s.CloseLogFile()
	s.ConfigureLogger()
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

func (s *Server) ConfigureLogger() {
	s.Logger.SetLevel(logrus.DebugLevel)
	// s.Logger.SetOutput(s.LogFile)
	s.Logger.SetFormatter(s.Logger.Formatter)
}

func (s *Server) ConfigureRouter() {
	// gin.DisableConsoleColor()
	// gin.SetMode(gin.ReleaseMode)
	// gin.DefaultWriter = io.MultiWriter(s.LogFile)
	// s.Router.Use(gin.Recovery())
	s.Router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
