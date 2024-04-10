package space

import (
	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	SpaceService SpaceService
}

func (c *Controller) RegisterRoutes(router *gin.Engine, adminRoute *gin.RouterGroup) {
	adminRoute.POST("/v1/space", c.PostSpace)
	adminRoute.GET("/v1/spaces", c.GetAllSpaces)
	adminRoute.GET("/v1/space/:spaceId", c.GetSpaceByID)
}

func (c *Controller) PostSpace(ctx *gin.Context) {
	var space database.Space

	if err := ctx.ShouldBindJSON(&space); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Your request contains invalid syntax. Please review and correct it",
			"error":   err.Error(),
		})
		return
	}

	if err := c.SpaceService.CreateSpace(space); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed to Create Space",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Created Space",
	})
}

func (c *Controller) GetSpaceByID(ctx *gin.Context) {
	spaceId := ctx.Param("spaceId")

	space, err := c.SpaceService.GetSpaceByID(spaceId)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed to get the SpaceID",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": space,
	})

}

func (c *Controller) GetAllSpaces(ctx *gin.Context) {
	spaces, err := c.SpaceService.GetAllSpaces()

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Failed to get Spaces",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": spaces,
	})
}
