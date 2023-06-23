package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jwt_auth/database"
	"github.com/jwt_auth/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
