package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/valentinRog/chess-puzzles-api/puzzle"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	client.Database("db").Drop(context.Background())
	col := client.Database("db").Collection("puzzles")
	_, err = col.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.M{
				"elo": 1,
			},
		},
		{
			Keys: bson.M{
				"npieces": 1,
			},
		},
		{
			Keys: bson.M{
				"i": 1,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	err = puzzle.Populate(col, "puzzles.csv")
	if err != nil {
		log.Fatal(err)
	}
}
