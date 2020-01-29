package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func startServer() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	apiRoute(e)
	socketRoute(e)

	log.Fatal(e.Start(":9393"))
}

func apiRoute(e *echo.Echo) {
	e.POST("/api/v1/login", Login)
	e.POST("/api/v1/register", Register)
}

func socketRoute(e *echo.Echo) {

	e.GET("/ws", WebSocket)

}
