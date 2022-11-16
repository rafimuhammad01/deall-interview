package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/user-service/database"
	"github.com/rafimuhammad01/user-service/internal/handler"
	"github.com/rafimuhammad01/user-service/internal/repository"
	"github.com/rafimuhammad01/user-service/internal/service"
)

type Server struct {
	Router *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Router: gin.Default(),
	}
}

func (s *Server) Init() {
	mongoDB := database.Connect()

	healthCheckHandler := handler.NewHealthCheck()

	hashAlgoService := service.NewHash()
	UserRepository := repository.NewUser(mongoDB)
	UserService := service.NewUser(UserRepository, hashAlgoService)
	UserHandler := handler.NewUser(UserService)

	r := NewRoutes(s.Router, healthCheckHandler, UserHandler)
	r.Init()
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
