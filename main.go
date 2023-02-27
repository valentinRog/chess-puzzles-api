package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valentinRog/chess-puzzles-api/database"
	"github.com/valentinRog/chess-puzzles-api/routes"
)

func main() {
	database.Connect()
	database.InitTables()
	go database.Populate()
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":80")
}