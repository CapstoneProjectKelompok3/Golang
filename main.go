package main

import (
	"project-capston/app/config"
	"project-capston/app/database"
	"project-capston/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	mysql:=database.InitMysql(cfg)
	database.InitialMigration(mysql)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	router.InitRouter(mysql, e)
	e.Logger.Fatal(e.Start(":80"))

}