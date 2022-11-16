package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/auth-service/database"
	"github.com/rafimuhammad01/auth-service/internal/handler"
	"github.com/rafimuhammad01/auth-service/internal/repository"
	"github.com/rafimuhammad01/auth-service/internal/service"
	"log"
	"os"
	"strconv"
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
	expiredTimeStr := os.Getenv("JWT_EXPIRED_AT")
	expireTimeInt, err := strconv.Atoi(expiredTimeStr)
	if err != nil {
		log.Fatalln("Wrong expired time format")
	}

	mongoDB := database.ConnectMongoDB()
	redis := database.ConnectRedis()

	authRepository := repository.NewUser(mongoDB, redis)
	jwtService := service.NewJWT(os.Getenv("JWT_SECRET"), authRepository, expireTimeInt)
	userHandler := handler.NewUser(jwtService)

	r := NewRoutes(s.Router, userHandler)
	r.Init()
}

func (s *Server) Run(addr string) {
	s.Router.Run(addr)
}
