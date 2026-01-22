package api

import (
	"gin_learning/internal/api/handler"
	"gin_learning/internal/api/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	logger, _ := zap.NewProduction()

	// 全局中间件
	r.Use(
		gin.Recovery(),
		middleware.RequestID(),
		middleware.LoggerMiddleware(logger),
	)

	// API版本分组
	v1 := r.Group("/api/v1")
	{
		// 用户模块
		userGroup := v1.Group("/users")
		userHandler := handler.NewUserHandler()
		{
			userGroup.POST("/createUser", userHandler.CreateUser)
		}
	}

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
