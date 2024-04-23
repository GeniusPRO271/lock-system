package main

import (
	stdlog "log"
	"os"

	"github.com/GeniusPRO271/lock-system/pkg/database"
	"github.com/GeniusPRO271/lock-system/pkg/device"
	"github.com/GeniusPRO271/lock-system/pkg/jwt"
	"github.com/GeniusPRO271/lock-system/pkg/log"
	"github.com/GeniusPRO271/lock-system/pkg/role"
	"github.com/GeniusPRO271/lock-system/pkg/space"
	"github.com/GeniusPRO271/lock-system/pkg/user"
	"github.com/GeniusPRO271/lock-system/pkg/whitelist"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tuya/tuya-connector-go/connector"
	tuyaENV "github.com/tuya/tuya-connector-go/connector/env"
	"github.com/tuya/tuya-connector-go/connector/httplib"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		stdlog.Fatal("Error loading .env file")
	}

	connector.InitWithOptions(tuyaENV.WithApiHost(httplib.URL_EU),
		tuyaENV.WithMsgHost(httplib.MSG_EU),
		tuyaENV.WithAccessID(os.Getenv("TUYA_ACCESS_ID")),
		tuyaENV.WithAccessKey(os.Getenv("TUYA_ACCESS_KEY")),
		tuyaENV.WithAppName(os.Getenv("TUYA_APP_NAME")),
		tuyaENV.WithDebugMode(true))

	router := gin.Default()
	db := database.Start_db()

	adminRoutes := router.Group("/admin")
	verifyRoutes := router.Group("")
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

	// Start Space
	spaceService := &space.SpaceServiceImpl{Db: db}
	spaceHandler := space.Controller{
		SpaceService: spaceService,
	}
	spaceHandler.RegisterRoutes(router, adminRoutes)

	// Start Whitelist
	whitelistService := &whitelist.WhitelistServiceImpl{
		Db:           db,
		SpaceService: spaceService,
	}
	whitelistHandler := whitelist.Controller{
		WhitelistService: whitelistService,
	}
	whitelistHandler.RegisterRoutes(router, adminRoutes)

	// Start Role
	roleService := &role.RoleServiceImpl{Db: db}
	roleHandler := role.Controller{
		RoleService: roleService,
	}
	roleHandler.RegisterRoutes(router, adminRoutes)

	// Start the server
	stdlog.Printf("Starting server at port 8080")

	err = router.Run(":8080")
	if err != nil {
		stdlog.Fatal("Server failed to start: ", err)
	}

}
