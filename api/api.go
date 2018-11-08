package api

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var (
	e *echo.Echo
)

func StartAPP() {
	HTTPAddr := ":8000"
	e.Logger.Fatal(e.Start(HTTPAddr))
}

func InitApi() {
	e = echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// web
	e.File("/", "web/dist/index.html")
	e.Static("/static", "web/dist/static")

	// api
	srv := e.Group("/api")

	srv.POST("/verify", verify)
}
