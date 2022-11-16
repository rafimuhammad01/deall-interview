package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/api-gateway/internal"
	"os"
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
	httpHelper := internal.NewHTTPHelper()
	authHandler := internal.NewAuth(os.Getenv("AUTH_BASE_URL"), httpHelper)
	userHandler := internal.NewUser(os.Getenv("USER_BASE_URL"), httpHelper)
	authMiddleware := internal.NewAuthMiddleware()

	r := NewRoutes(s.Router, authHandler, userHandler, authMiddleware)
	r.Init()
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
