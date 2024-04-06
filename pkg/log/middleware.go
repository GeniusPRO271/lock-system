package log

import (
	"log"
	"strconv"
	"time"

	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LoggerMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		user := c.MustGet("userID").(string)

		userId, err := strconv.ParseUint(user, 10, 32)
		if err != nil {
			// Handle error
			log.Println("Error:", err)
			return
		}

		logToSave := database.Log{
			UserID: uint(userId),
			// Device
			// Instruction
		}

		if err := db.Create(&logToSave).Error; err != nil {
			c.JSON(500, gin.H{"error": "Error at creating log"})
			c.Abort()
			return
		}

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Log request details
		host := c.Request.Host
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		log.Printf("[GIN] %s %d %s %s %v %s", host, statusCode, method, path, latency, user)
	}
}
