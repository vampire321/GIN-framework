package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mithun/gin/internal/handler"
	"github.com/mithun/gin/internal/middleware"
	"github.com/mithun/gin/internal/repository"
	"github.com/mithun/gin/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	r := gin.New()
	r.Use(middleware.Logger(logger))

	UserRepository := repository.NewUserRepository()
	userService := service.NewUserService(UserRepository)
	userHandler := handler.NewUserHandler(userService)
	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		api.POST("/users",userHandler.Create) //connect HTTP route to handler method 
	}
	r.Run(":8080")
}
	