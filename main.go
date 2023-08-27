package main

import (
	"github.com/valentinRog/chess-puzzles-api/puzzle"
	"log"
)

func main() {
	disconnect, err := puzzle.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer disconnect()
	err = puzzle.Populate("puzzles.csv")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = puzzle.Shuffle()
		if err != nil {
			log.Fatal(err)
		}
	}()

	
}
