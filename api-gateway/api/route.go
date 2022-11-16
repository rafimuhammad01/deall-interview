package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/api-gateway/internal"
)

type Routes struct {
	ginRouter      *gin.Engine
	authHandler    *internal.Auth
	userHandler    *internal.User
	authMiddleware *internal.AuthMiddleware
}

func NewRoutes(ginRouter *gin.Engine, authHandler *internal.Auth, userHandler *internal.User, authMiddleware *internal.AuthMiddleware) *Routes {
	return &Routes{
		ginRouter:      ginRouter,
		authHandler:    authHandler,
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
	}
}

func (r *Routes) Init() {
	auth := r.ginRouter.Group("/api/v1/auth")
	auth.POST("/login", r.authHandler.Login)
	auth.POST("/refresh", r.authHandler.RefreshToken)

	users := r.ginRouter.Group("/api/v1/user")
	users.GET("", r.authMiddleware.TokenAuthMiddleware(internal.RoleAdmin, internal.RoleUser), r.userHandler.GetAll)
	users.POST("", r.authMiddleware.TokenAuthMiddleware(internal.RoleAdmin), r.userHandler.Create)
	users.GET("/:id", r.authMiddleware.TokenAuthMiddleware(internal.RoleAdmin, internal.RoleUser), r.userHandler.GetByID)
	users.PUT("/:id", r.authMiddleware.TokenAuthMiddleware(internal.RoleAdmin), r.userHandler.Update)
	users.DELETE("/:id", r.authMiddleware.TokenAuthMiddleware(internal.RoleAdmin), r.userHandler.Delete)
}
