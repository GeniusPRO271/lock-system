package instruction

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	InstructionService InstructionService
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	// router.POST("/v1/device/:deviceId/instruction", c.PostLog)
	// router.GET("/v1/device/:deviceId/instruction", c.GetDeviceInstruction)
	// router.GET("/v1/device/:deviceId/instruction/:id ", c.GetLogById)
}

func (c *Controller) GetDevice(ctx *gin.Context) {

}

func (c *Controller) GetDeviceInstruction(ctx *gin.Context) {

}
