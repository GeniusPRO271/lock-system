package main

import (
	stdlog "log"

	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/device"
	"github.com/GeniusPRO271/lock-system/pkg/jwt"
	"github.com/GeniusPRO271/lock-system/pkg/log"
	"github.com/GeniusPRO271/lock-system/pkg/user"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.Start_db()

	privateRouter := router.Group("/private")
	privateRouter.Use(jwt.AuthMiddleware())
	privateRouter.Use(log.LoggerMiddleware(db))

	// Start User
	userService := &user.UserServiceImpl{Db: db}
	userHandler := user.Controller{
		UserService: userService,
	}

	userHandler.RegisterRoutes(router, privateRouter)

	// Start Log
	logService := &log.LogServiceImpl{}
	logHandler := log.Controller{
		LogService: logService,
	}
	logHandler.RegisterRoutes(router)

	// Start Device
	deviceService := &device.DeviceServiceImpl{Db: db}
	deviceHandler := device.Controller{
		DeviceService: deviceService,
	}
	deviceHandler.RegisterRoutes(router)

	// Start the server
	stdlog.Printf("Starting server at port 8080")
	err := router.Run(":8080")
	if err != nil {
		stdlog.Fatal("Server failed to start: ", err)
	}
}
