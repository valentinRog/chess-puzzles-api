package puzzle

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func Init(db *sql.DB) error {
	_, err := db.Exec(`
DROP TABLE IF EXISTS puzzles;
CREATE TABLE puzzles(
	id INTEGER PRIMARY KEY,
	moves TEXT,
	fen TEXT,
	elo INTEGER,
	npieces INTEGER
);
`)
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX puzzles_id_elo_npieces ON puzzles(id, elo, npieces);")
	if err != nil {
		return err
	}
	return nil
}

func Populate(db *sql.DB, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const query = "INSERT INTO puzzles(moves, fen, elo, npieces) VALUES(?, ?, ?, ?);"
	const BatchSize = 10000

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	i := 0
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), ",")
		fen := arr[1]
		moves := arr[2]
		elo := arr[3]
		npieces := 9
		_, err = stmt.Exec(moves, fen, elo, npieces)
		if err != nil {
			return err
		}
		if i%BatchSize == 0 {
			fmt.Printf("%d\r", i)
			err = tx.Commit()
			if err != nil {
				return err
			}
			tx, err = db.Begin()
			if err != nil {
				return err
			}
			stmt.Close()
			stmt, err = tx.Prepare(query)
			if err != nil {
				return err
			}
			defer stmt.Close()
		}
		i += 1
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
