package puzzle

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Populate(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)

	record, err := r.Read()
	if err != nil {
		return err
	}
	fields := map[string]int{}
	for i, v := range record {
		fields[strings.ToLower(v)] = i
	}

	buff := []interface{}{}
	for i := 0; ; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fen := record[fields["fen"]]
		moves := record[fields["moves"]]
		elo, err := strconv.Atoi(record[fields["rating"]])
		if err != nil {
			return err
		}
		npieces := func() int {
			n := 0
			for _, c := range strings.Split(fen, " ")[0] {
				if unicode.IsLetter(c) {
					n++
				}
			}
			return n
		}()
		p := Puzzle{
			Id:      primitive.NewObjectID(),
			Fen:     fen,
			Moves:   moves,
			Elo:     elo,
			Npieces: npieces,
			I:       i,
		}
		buff = append(buff, p)
		if len(buff) == 10_000 {
			_, err = Col.InsertMany(context.Background(), buff)
			if err != nil {
				return err
			}
			buff = []interface{}{}
		}
	}
	_, err = Col.InsertMany(context.Background(), buff)
	if err != nil {
		return err
	}
	return nil
}

func Shuffle() error {
	var p Puzzle
	Col.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{"i": -1})).Decode(&p)
	for i := 0; i < p.I; i++ {
		i2 := rand.Intn(p.I)
		var p1, p2 Puzzle
		Col.FindOne(context.Background(), bson.M{"i": i}).Decode(&p1)
		Col.FindOne(context.Background(), bson.M{"i": i2}).Decode(&p2)
		Col.UpdateOne(context.Background(), bson.M{"_id": p1.Id}, bson.M{"$set": bson.M{"i": i2}})
		Col.UpdateOne(context.Background(), bson.M{"_id": p2.Id}, bson.M{"$set": bson.M{"i": i}})
		if i%1000 == 0 {
			fmt.Printf("%d\r", i)
		}
	}
	return nil
}

func GetRandom(elo_min int, elo_max int) (Puzzle, error) {
	var p Puzzle
	err := Col.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{"i": -1})).Decode(&p)
	if err != nil {
		return Puzzle{}, err
	}
	fmt.Println(p)
	i := rand.Intn(p.I)
	if i < p.I/2 {
		query := bson.M{"i": bson.M{"$gte": i}, "elo": bson.M{"$gte": elo_min, "$lte": elo_max}}
		err = Col.FindOne(context.Background(), query, options.FindOne().SetSort(bson.M{"i": 1})).Decode(&p)
		if err != nil {
			return Puzzle{}, err
		}
	} else {
		query := bson.M{"i": bson.M{"$lte": i}, "elo": bson.M{"$gte": elo_min, "$lte": elo_max}}
		err = Col.FindOne(context.Background(), query, options.FindOne().SetSort(bson.M{"i": -1})).Decode(&p)
		if err != nil {
			return Puzzle{}, err
		}
	}
	return p, nil
}
