package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Choice struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Choice  string             `json:"choice"`
	Ranking int                `json:"ranking"`
}
