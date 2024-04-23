package whitelist

import (
	"net/http"

	"github.com/GeniusPRO271/lock-system/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	WhitelistService WhitelistService
}

func (c *Controller) RegisterRoutes(router *gin.Engine, adminRote *gin.RouterGroup) {
	adminRote.POST("/v1/space/add-user-whitelist", c.PostUserToWhitelist)
	adminRote.GET("/v1/space/:spaceId/whitelist", c.GetWhitelist)
	adminRote.GET("/v1/space/whitelist", c.IsUserInWhitelist)
	adminRote.GET("/v1/space/whitelist/:userId/allowed", c.GetUserAllowedSpacesHandler)
	adminRote.DELETE("/v1/space/delete-user-whitelist", c.DeleteUserFromWhitelist)
}

func (c *Controller) GetUserAllowedSpacesHandler(ctx *gin.Context) {
	// Get user ID from the context or request parameters
	userId, err := utils.StringToUint(ctx.Param("userId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call service function to get allowed spaces
	allowedSpaces, err := c.WhitelistService.GetRootSpaces(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get allowed spaces"})
		return
	}

	// Return the list of allowed spaces
	ctx.JSON(http.StatusOK, allowedSpaces)
}

func (c *Controller) PostUserToWhitelist(ctx *gin.Context) {
	var params PostDeleteWhiteListParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	progapageVal := false
	if params.Propagate != nil {
		progapageVal = *params.Propagate
	}

	if err := c.WhitelistService.AddUserToWhitelist(params.UserID, params.SpaceID, progapageVal); err != nil {
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

func (c *Controller) IsUserInWhitelist(ctx *gin.Context) {
	// Get user ID and space ID from the request
	userId, err := utils.StringToUint(ctx.Query("userId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	spaceId, err := utils.StringToUint(ctx.Query("spaceId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if user is in the whitelist for the specified space
	isInWhitelist, err := c.WhitelistService.IsUserInWhitelist(userId, spaceId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"userId": userId, "spaceId": spaceId, "isInWhitelist": isInWhitelist})
}

func (c *Controller) DeleteUserFromWhitelist(ctx *gin.Context) {
	var params PostDeleteWhiteListParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	progapageVal := false
	if params.Propagate != nil {
		progapageVal = *params.Propagate
	}

	err := c.WhitelistService.DeleteUserFromWhitelist(params.UserID, params.SpaceID, progapageVal)
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
