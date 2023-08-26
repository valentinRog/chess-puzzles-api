package puzzle

import (
	"bufio"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func Populate(col *mongo.Collection, filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	buff := []interface{}{}
	i := 0
	for scanner.Scan() {
		i += 1
		arr := strings.Split(scanner.Text(), ",")
		if arr[3] == "Rating" {
			continue
		}
		fen := arr[1]
		moves := arr[2]
		elo, err := strconv.Atoi(arr[3])
		if err != nil {
			return err
		}
		npieces := 9
		p := Puzzle{
			Fen:     fen,
			Moves:   moves,
			Elo:     elo,
			Npieces: npieces,
			I:       i,
		}
		buff = append(buff, p)
		if len(buff) == 10000 {
			_, err = col.InsertMany(context.Background(), buff)
			if err != nil {
				return err
			}
			buff = []interface{}{}
		}
	}
	_, err = col.InsertMany(context.Background(), buff)
	if err != nil {
		return err
	}
	return nil
}

func RandomPuzzle(col *mongo.Collection, elo_min int, elo_max int) (Puzzle, error) {
	var p Puzzle
	err := col.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{"i": -1})).Decode(&p)
	if err != nil {
		return Puzzle{}, err
	}
	fmt.Println(p)
	i := rand.Intn(p.I)
	if i < p.I/2 {
		query := bson.M{"i": bson.M{"$gte": i}, "elo": bson.M{"$gte": elo_min, "$lte": elo_max}}
		err = col.FindOne(context.Background(), query, options.FindOne().SetSort(bson.M{"i": 1})).Decode(&p)
		if err != nil {
			return Puzzle{}, err
		}
	} else {
		query := bson.M{"i": bson.M{"$lte": i}, "elo": bson.M{"$gte": elo_min, "$lte": elo_max}}
		err = col.FindOne(context.Background(), query, options.FindOne().SetSort(bson.M{"i": -1})).Decode(&p)
		if err != nil {
			return Puzzle{}, err
		}
	}
	return p, nil
}
