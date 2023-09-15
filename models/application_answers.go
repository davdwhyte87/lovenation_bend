
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApplicationAnswers struct {
    Id       primitive.ObjectID `bson:"_id"  json:"id,omitempty"`
    ApplicationId     string             `json:"application_id,omitempty"`
    QuestionId string             `json:"question_id,omitempty" validate:"required"`
    YesNoAnswer    bool             `json:"yes_no_answer,omitempty" `
	TextAnswer    string             `json:"text_answer,omitempty" `
	CreatedAt    string             `json:"created_at,omitempty" `
}
