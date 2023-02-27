package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valentinRog/chess-puzzles-api/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.GetPuzzle)
}
