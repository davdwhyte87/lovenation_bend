package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id         primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	FullName   string             `json:"name,omitempty" validate:"required"`
	Location   string             `json:"location,omitempty" validate:"required"`
	Profession string             `json:"title,omitempty" validate:"required"`
	Type       UserType           `bson:"user_type"`
	UserRoleId int                `bson:"user_role_id"`
}

type UserType int

const (
	Citizen UserType = iota
	Admin
)
