
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApplicationQuestion struct {
    Id       primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	Text string
	Type QuestionType
	CreatedAt string 

}

type QuestionType int

const (
	YesNo QuestionType = iota
	Text 
	MultiChoice
)