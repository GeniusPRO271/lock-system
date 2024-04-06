package user

import (
	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserService UserService
}

func (c *Controller) RegisterRoutes(router *gin.Engine, privateRouter *gin.RouterGroup) {
	router.POST("/v1/user/register", c.PostRegisterUser)
	router.POST("/v1/user/login", c.PostLogin)
	privateRouter.PUT("/v1/user/edit", c.PutEditUser) // not implemented yet
	privateRouter.GET("/v1/user/:id", c.GetUserId)
	privateRouter.GET("/v1/users", c.GetUsers)
}

// PostRegisterUser handles the HTTP POST request to register a new user.
// It takes an HTTP response writer and request as input.
func (c *Controller) PostRegisterUser(ctx *gin.Context) {
	// Implement registration logic here.
	var user database.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Your request contains invalid syntax. Please review and correct it",
			"error":   err.Error(),
		})
		return
	}

	if err := c.UserService.CreateUser(user); err != nil {
		ctx.JSON(404, gin.H{
			"message": "Error at creating new user",
			"error":   err.Error(),
		})
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
	var user UserLogin

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"message": "Your request contains invalid syntax. Please review and correct it",
			"error":   err.Error(),
		})

		return
	}

	token, err := c.UserService.VerifyUser(user)

	if err != nil {
		ctx.JSON(401, gin.H{
			"message": "Incorrect username or password. Please try again.",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
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
func (c *Controller) GetUserId(ctx *gin.Context) {
	// Implement logic to retrieve a user by ID here.
	userID := ctx.Param("id")

	user, err := c.UserService.GetUserByID(userID)

	if err != nil {
		ctx.JSON(404, gin.H{
			"message": "Could not find a user with that ID",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": user,
	})

}
