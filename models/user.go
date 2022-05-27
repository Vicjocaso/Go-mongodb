package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User Data
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	CreateAt time.Time          `bson:"create_At" json:"create_at"`
	UpdateAt time.Time          `bson:"update_At" json:"update_at,omitempty"`
}

type Users []*User
