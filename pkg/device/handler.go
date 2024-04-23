package device

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	DeviceService DeviceService
}

func (c *Controller) RegisterRoutes(router *gin.Engine) {
	router.POST("/v1/devices", c.SyncDevices)
	router.PUT("/v1/device", c.UpdateDeviceSpace)
	router.GET("/v1/devices", c.GetDevices)
	router.GET("/v1/device/:deviceId", c.GetDeviceById)
}

func (c *Controller) SyncDevices(ctx *gin.Context) {
	result, err := c.DeviceService.SyncDeviceList()
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error at syncing new devices",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Devices Syncing",
		"result":  result,
	})
}

func (c *Controller) GetDevices(ctx *gin.Context) {

	resp, err := c.DeviceService.GetDevices()
	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error at getting devices",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"result": resp})
}

func (c *Controller) GetDeviceById(ctx *gin.Context) {
	deviceID := ctx.Param("deviceId")

	device, err := c.DeviceService.GetDeviceById(deviceID)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Could not find a device with that ID, remember to sync if you added a device",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"result": device,
	})

}

type UpdateDeviceSpaceDto struct {
	DeviceId uint `binding:"required"`
	SpaceId  uint `binding:"required"`
}

func (c *Controller) UpdateDeviceSpace(ctx *gin.Context) {
	var toUpdateData UpdateDeviceSpaceDto

	if err := ctx.ShouldBindJSON(&toUpdateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.DeviceService.UpdateDevicesSpace(&toUpdateData); err != nil {
		ctx.JSON(404, gin.H{
			"message": "Could not update device Space",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Device Space Updated",
		"Device":  toUpdateData.DeviceId,
		"Space":   toUpdateData.SpaceId,
	})
}
