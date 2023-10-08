package main

import (
	"net/http"
	"project-capston/app/config"
	"project-capston/app/database"
	"project-capston/app/middlewares"
	"project-capston/app/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	database.InitialMigration(mysql)

	redis:=middlewares.CreateRedisClient()


	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"https://api.flattenbot.site"}, // Ganti dengan origin yang diizinkan Anda.
        AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
        AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
    }))
	router.InitRouter(mysql, e,redis)
	e.Logger.Fatal(e.Start(":80"))

}
