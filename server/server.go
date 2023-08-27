package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valentinRog/chess-puzzles-api/puzzle"
	"strconv"
)

func Setup() *fiber.App {
	app := fiber.New()

	app.Get("/puzzle", func(c *fiber.Ctx) error {
		var eloMin int
		var eloMax int
		var npiecesMin int
		var npiecesMax int
		if c.Query("elo_min") == "" {
			p, _ := puzzle.Min("elo")
			eloMin = p.Elo
		} else {
			eloMin, _ = strconv.Atoi(c.Query("elo_min"))
		}
		if c.Query("elo_max") == "" {
			p, _ := puzzle.Max("elo")
			eloMax = p.Elo
		} else {
			eloMax, _ = strconv.Atoi(c.Query("elo_max"))
		}
		if c.Query("npieces_min") == "" {
			p, _ := puzzle.Min("npieces")
			npiecesMin = p.Npieces
		} else {
			npiecesMin, _ = strconv.Atoi(c.Query("npieces_min"))
		}
		if c.Query("npieces_max") == "" {
			p, _ := puzzle.Max("npieces")
			npiecesMax = p.Npieces
		} else {
			npiecesMax, _ = strconv.Atoi(c.Query("npieces_max"))
		}
		p, _ := puzzle.GetRandom(eloMin, eloMax, npiecesMin, npiecesMax)
		return c.JSON(p)
	})

	return app
}
