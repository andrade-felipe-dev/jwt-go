package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEntity struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Email string             `bson:"email"`
	Senha string             `bson:"senha"`
	Nome  string             `bson:"nome"`
	Tipo  string             `bson:"tipo"`
}
