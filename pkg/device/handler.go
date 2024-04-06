package device

import (
	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	DeviceService DeviceService
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/v1/device", c.PostDevice)
	router.GET("/v1/devices", c.GetDevices)
	router.GET("/v1/device/:id", c.GetDeviceById)
}

func (c *Controller) PostDevice(ctx *gin.Context) {
	var device database.Device

	if err := ctx.ShouldBindJSON(&device); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Your request contains invalid syntax. Please review and correct it",
			"error":   err.Error(),
		})
		return
	}

	if err := c.DeviceService.CreateDevice(device); err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error at creating new device",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Device Created",
	})
}

func (c *Controller) GetDevices(ctx *gin.Context) {

	devices, err := c.DeviceService.GetDevices()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error at geting devices data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": devices,
	})
}

func (c *Controller) GetDeviceById(ctx *gin.Context) {
	deviceID := ctx.Param("id")

	user, err := c.DeviceService.GetDeviceById(deviceID)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Could not find a device with that ID",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})

}
