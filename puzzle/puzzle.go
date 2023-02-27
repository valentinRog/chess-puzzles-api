package puzzle

import (
	"github.com/valentinRog/chess-puzzles-api/models"
	// "github.com/notnil/chess"
)

type Puzzle struct {
	Fen     string `json:"fen"`
	Moves   string `json:"moves"`
	Rating  int    `json:"rating"`
	NPieces uint8  `json:"n_pieces"`
}

func FromModel(p models.Puzzle) Puzzle {
	return Puzzle{
		Fen:     p.FEN,
		Moves:   p.Moves,
		Rating:  p.Rating,
		NPieces: p.NPieces,
	}
}
