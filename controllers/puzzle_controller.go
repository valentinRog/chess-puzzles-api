package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/valentinRog/chess-puzzles-api/database"
	"github.com/valentinRog/chess-puzzles-api/models"
	"github.com/valentinRog/chess-puzzles-api/puzzle"
	"strconv"
)

func getRandomPuzzle(ratingMin int, ratingMax int, nPiecesMin uint8, nPiecesMax uint8) models.Puzzle {
	var p models.Puzzle
	database.DB.Exec("SET SESSION query_cache_type = OFF")
	database.DB.Raw("SELECT * FROM puzzles WHERE id >= ((SELECT MAX(id) FROM puzzles)-(SELECT MIN(id) FROM puzzles)) * RAND() + (SELECT MIN(id) FROM puzzles) AND rating >= ? AND rating <= ? AND n_pieces >= ? AND n_pieces <= ? LIMIT 1", ratingMin, ratingMax, nPiecesMin, nPiecesMax).Scan(&p)
	return p
}

func GetPuzzle(c *fiber.Ctx) error {
	ratingMin, _ := strconv.Atoi(c.Query("ratingMin", "0"))
	ratingMax, _ := strconv.Atoi(c.Query("ratingMax", "4000"))
	nPiecesMin, _ := strconv.Atoi(c.Query("nPiecesMin", "0"))
	nPiecesMax, _ := strconv.Atoi(c.Query("nPiecesMax", "32"))
	p := getRandomPuzzle(ratingMin, ratingMax, uint8(nPiecesMin), uint8(nPiecesMax))
	return c.JSON(puzzle.FromModel(p))
}
