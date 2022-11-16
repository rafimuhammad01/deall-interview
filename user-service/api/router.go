package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/user-service/internal/handler"
)

type Routes struct {
	ginRouter          *gin.Engine
	healthCheckHandler *handler.HealthCheck
	userHandler        *handler.User
}

func NewRoutes(ginRouter *gin.Engine, healthCheckHandler *handler.HealthCheck, userHandler *handler.User) *Routes {
	return &Routes{
		ginRouter:          ginRouter,
		healthCheckHandler: healthCheckHandler,
		userHandler:        userHandler,
	}
}

func (r *Routes) Init() {
	r.ginRouter.GET("/", r.healthCheckHandler.HealthCheck)

	users := r.ginRouter.Group("/api/v1/user")
	users.GET("", r.userHandler.GetAll)
	users.POST("", r.userHandler.Create)
	users.GET("/:id", r.userHandler.GetByID)
	users.PUT("/:id", r.userHandler.Update)
	users.DELETE("/:id", r.userHandler.Delete)
}
