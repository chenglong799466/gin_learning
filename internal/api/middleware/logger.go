package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	// 需要跳过的日志路径
	skipPaths := map[string]bool{
		"/health":  true,
		"/metrics": true,
	}

	return func(c *gin.Context) {
		// 跳过指定路径的日志记录
		if _, ok := skipPaths[c.Request.URL.Path]; ok {
			c.Next()
			return
		}
		// 记录请求开始时间
		start := time.Now()
		// 结构化日志字段
		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int("response_size", c.Writer.Size()),
			zap.String("request_id", c.GetString("request_id")),
			zap.Time("timestamp", start.UTC()),
		}

		// 记录日志
		logger.Info("request", fields...)
		c.Next()
	}
}
