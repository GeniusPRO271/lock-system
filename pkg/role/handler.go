package role

import (
	"errors"
	"net/http"
	"strconv"

	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	RoleService RoleService
}

func (c *Controller) RegisterRoutes(router *gin.Engine, adminRoute *gin.RouterGroup) {
	adminRoute.POST("v1/user/role", c.CreateRole)
	adminRoute.GET("v1/user/roles", c.GetRoles)
	adminRoute.PUT("v1/user/role/:id", c.UpdateRole)
}

// create Role
func (c *Controller) CreateRole(ctx *gin.Context) {
	var Role model.Role
	ctx.BindJSON(&Role)
	err := c.RoleService.CreateRole(&Role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, Role)
}

// get Roles
func (c *Controller) GetRoles(ctx *gin.Context) {
	var Role []model.Role
	err := c.RoleService.GetRoles(&Role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, Role)
}

// get Role by id
func (c *Controller) GetRole(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var Role model.Role
	err := c.RoleService.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, Role)
}

// update Role
func (c *Controller) UpdateRole(ctx *gin.Context) {
	var Role model.Role
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.RoleService.GetRole(&Role, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.BindJSON(&Role)
	err = c.RoleService.UpdateRole(&Role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, Role)
}
