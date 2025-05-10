package main

import (
	"project/database"
	"project/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	println("Starting the application...")
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173","http://localhost",},
        AllowMethods: []string{
            "GET", "POST", "PUT", "DELETE", "OPTIONS",
        },
        AllowHeaders: []string{"Content-Type", "Authorization", "X-Requested-With", "X-Real-IP",},
	}))
	database.ConnectDB()
	database.Migrate()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
