package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApplicationAnswers struct {
	Id            primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
	ApplicationId primitive.ObjectID `bson:"application_id,omitempty"`
	QuestionId    string             `bson:"question_id,omitempty" validate:"required"`
	YesNoAnswer   bool               `bson:"yes_no_answer,omitempty" `
	TextAnswer    string             `bson:"text_answer,omitempty" `
	CreatedAt     string             `bson:"created_at,omitempty" `
}
