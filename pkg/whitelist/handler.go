package whitelist

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	WhitelistService WhitelistService
}

func (c *Controller) RegisterRoutes(router *gin.Engine, adminRote *gin.RouterGroup) {
	adminRote.POST("/v1/space/:spaceId/whitelist/:userId", c.PostUserToWhitelist)
	adminRote.GET("/v1/space/:spaceId/whitelist", c.GetWhitelist)
	adminRote.DELETE("/v1/space/:spaceId/whitelist/:userId", c.DeleteUserFromWhitelist)
}

func (c *Controller) PostUserToWhitelist(ctx *gin.Context) {
	spaceId := ctx.Param("spaceId")
	userId := ctx.Param("userId")

	if err := c.WhitelistService.AddUserToWhitelist(userId, spaceId); err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not add User to Space Whitelists",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User added to Whitelist",
	})
}

func (c *Controller) GetWhitelist(ctx *gin.Context) {
	spaceId := ctx.Param("spaceId")

	users, err := c.WhitelistService.GetUsersFromSpaceWhitelist(spaceId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not get Users from the Whitelist",
			"error":   err.Error(),
		})
		return
	}

	if users != nil {
		ctx.JSON(200, gin.H{
			"message": users,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "empty",
	})
}

func (c *Controller) DeleteUserFromWhitelist(ctx *gin.Context) {
	spaceId := ctx.Param("spaceId")
	userId := ctx.Param("userId")

	err := c.WhitelistService.DeleteUserFromWhitelist(userId, spaceId)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Could not delete user from the whitelist",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User deleted from whitelist",
	})
}
