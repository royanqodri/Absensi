package main

import (
	"Absensi-App/app/config"
	"Absensi-App/app/database"
	"Absensi-App/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// api.ApiGetUser()
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	redis := database.InitRedis(cfg)
	database.InitialMigration(mysql)
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(e, mysql, redis)
	e.Logger.Fatal(e.Start(":80"))

}
