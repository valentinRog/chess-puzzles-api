package puzzle

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var Col *mongo.Collection

func Init() (func() error, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	disconnect := func() error {
		return client.Disconnect(context.Background())
	}
	client.Database("db").Drop(context.Background())
	Col = client.Database("db").Collection("puzzles")
	_, err = Col.Indexes().CreateMany(context.Background(), []mongo.IndexModel{
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
		return nil, err
	}
	return disconnect, nil
}
