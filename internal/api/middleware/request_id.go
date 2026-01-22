package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := generateUUID() // 示例：b0b3d4f5-ae96-4b3d-97d7-3d61a12e5d1d
		c.Header("X-Request-ID", uuid)
		c.Set("request_id", uuid) // 存入上下文
		c.Next()
	}
}

// UUID v4 生成示例
func generateUUID() string {
	uuid := make([]byte, 16)
	if _, err := rand.Read(uuid); err != nil {
		panic(err)
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10
	return hex.EncodeToString(uuid)
}
