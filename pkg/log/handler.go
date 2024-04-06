package log

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	LogService LogService
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/v1/log", c.PostLog)
	router.GET("/v1/log", c.GetLog)
	router.GET("/v1/log/:id", c.GetLogById)
}

func (c *Controller) PostLog(ctx *gin.Context) {

}

func (c *Controller) GetLog(ctx *gin.Context) {

}

func (c *Controller) GetLogById(ctx *gin.Context) {

}
