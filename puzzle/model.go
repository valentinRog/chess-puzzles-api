package puzzle

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Puzzle struct {
	Id      primitive.ObjectID `bson:"_id"`
	Moves   string             `bson:"moves"`
	Fen     string             `bson:"fen"`
	Elo     int                `bson:"elo"`
	Npieces int                `bson:"npieces"`
	I       int                `bson:"i"`
}
