package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/auth-service/internal/handler"
)

type Routes struct {
	ginRouter   *gin.Engine
	authHandler *handler.Auth
}

func NewRoutes(ginRouter *gin.Engine, authHandler *handler.Auth) *Routes {
	return &Routes{
		ginRouter:   ginRouter,
		authHandler: authHandler,
	}
}

func (r *Routes) Init() {
	auth := r.ginRouter.Group("/api/v1/auth")
	auth.POST("/login", r.authHandler.Login)
	auth.POST("/refresh", r.authHandler.RefreshToken)
}
