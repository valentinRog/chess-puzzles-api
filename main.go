package main

import (
	"database/sql"
	"github.com/valentinRog/chess-puzzles-api/puzzle"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "file:dev.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = puzzle.Init(db)
	if err != nil {
		log.Fatal(err)
	}
	err = puzzle.Populate(db, "puzzles.csv")
	if err != nil {
		log.Fatal(err)
	}
}
