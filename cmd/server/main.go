package main

import (
	stdlog "log"

	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/device"
	"github.com/GeniusPRO271/lock-system/pkg/jwt"
	"github.com/GeniusPRO271/lock-system/pkg/log"
	"github.com/GeniusPRO271/lock-system/pkg/role"
	"github.com/GeniusPRO271/lock-system/pkg/space"
	"github.com/GeniusPRO271/lock-system/pkg/user"
	"github.com/GeniusPRO271/lock-system/pkg/whitelist"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.Start_db()

	adminRoutes := router.Group("/admin")
	verifyRoutes := router.Group("")
	adminRoutes.Use(jwt.JWTAuth())
	verifyRoutes.Use(jwt.JWTAuthCustomer())

	// Start User
	userService := &user.UserServiceImpl{Db: db}
	userHandler := user.Controller{
		UserService: userService,
	}

	userHandler.RegisterRoutes(router, adminRoutes)

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

	// Start Whitelist
	whitelistService := &whitelist.WhitelistServiceImpl{Db: db}
	whitelistHandler := whitelist.Controller{
		WhitelistService: whitelistService,
	}
	whitelistHandler.RegisterRoutes(router, adminRoutes)

	// Start Space
	spaceService := &space.SpaceServiceImpl{Db: db}
	spaceHandler := space.Controller{
		SpaceService: spaceService,
	}
	spaceHandler.RegisterRoutes(router, adminRoutes)

	// Start Space
	roleService := &role.RoleServiceImpl{Db: db}
	roleHandler := role.Controller{
		RoleService: roleService,
	}
	roleHandler.RegisterRoutes(router, adminRoutes)

	// Start the server
	stdlog.Printf("Starting server at port 8080")
	err := router.Run(":8080")
	if err != nil {
		stdlog.Fatal("Server failed to start: ", err)
	}
}
