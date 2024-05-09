package user

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	model "github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Controller struct {
	UserService UserService
}

func (c *Controller) RegisterRoutes(router *gin.Engine, adminRoute *gin.RouterGroup) {
	router.POST("/v1/user/register", c.PostRegisterUser)
	router.POST("/v1/user/login", c.PostLogin)
	adminRoute.PUT("/v1/user/:id", c.UpdateUser) // not implemented yet
	adminRoute.GET("/v1/user/:id", c.GetUserbyId)
	adminRoute.GET("/v1/users", c.GetUsers)
}

// PostRegisterUser handles the HTTP POST request to register a new user.
// It takes an HTTP response writer and request as input.
func (c *Controller) PostRegisterUser(ctx *gin.Context) {
	// Implement registration logic here.
	var user Register

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User Created",
	})
}

// PostLogin handles the HTTP POST request for user login.
// It takes an HTTP response writer and request as input.
func (c *Controller) PostLogin(ctx *gin.Context) {
	// Implement login logic here.
	var user Login

	if err := ctx.ShouldBindJSON(&user); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	data, err := c.UserService.VerifyUser(user)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user_data": data.User, "token": data.Token, "message": "Successfully logged in"})
}

// PutEditUser handles the HTTP PUT request to edit user information.
// It takes an HTTP response writer and request as input.
// Not implemented
func (c *Controller) PutEditUser(ctx *gin.Context) {
	// Implement user editing logic here.
	ctx.JSON(200, gin.H{
		"message": "PutEditUser",
	})
}

// GetUsers handles the HTTP GET request to retrieve all users.
// It takes an HTTP response writer and request as input.
func (c *Controller) GetUsers(ctx *gin.Context) {
	// Implement logic to retrieve all users here.
	user, err := c.UserService.GetUsers()

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error at creating new user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user.Users,
	})
}

// GetUserId handles the HTTP GET request to retrieve a user by ID.
// It takes an HTTP response writer and request as input.
func (c *Controller) GetUserbyId(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var user model.User
	err := c.UserService.GetUser(&user, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, UserGetResponse{
		Id:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Role:     utils.GetRoleNameByID(user.RoleID),
	})
}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	//var input model.Update
	var User model.User
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.UserService.GetUser(&User, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.BindJSON(&User)
	err = c.UserService.UpdateUser(&User)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, User)
}
