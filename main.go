package main

import (
	"Tasya_Project-Mini/config"
	"Tasya_Project-Mini/middleware"
	"Tasya_Project-Mini/route"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	e := echo.New()
	middleware.Logmiddleware(e)

	route.NewRoute(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
