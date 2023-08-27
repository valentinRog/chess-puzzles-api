package main

import (
	"github.com/valentinRog/chess-puzzles-api/puzzle"
	"github.com/valentinRog/chess-puzzles-api/server"
	"log"
)

func main() {
	disconnect, err := puzzle.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer disconnect()

	go func() {
		err = puzzle.Populate("puzzles.csv")
		if err != nil {
			log.Fatal(err)
		}
		err = puzzle.Shuffle()
		if err != nil {
			log.Fatal(err)
		}
	}()

	app := server.Setup()
	app.Listen(":3000")
}
