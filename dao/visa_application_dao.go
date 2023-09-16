package dao

import (
	"context"
	"fmt"
	"lovenation_bend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)


type VisaApplicationDAO struct {
	Collection *mongo.Collection
	Context context.Context
	
}
func (dao *VisaApplicationDAO )UpdateVisaApplication(visaApplication models.VisaApplication) {
	// Update an existing user
	// docID, _ := primit.ObjectIDFromHex(user.ID.Hex())
	// _, err := dao.Collection.UpdateOne(dao.ctx, bson.M{"_id": docID}, bson.M{"$set": user})
	// return err
}

// insert 


// get all visa applications 
func (dao *VisaApplicationDAO) GetAll (){
	pipe := bson.M{
		"$lookup":bson.M{
			"from":"VisaApplicationAnswers",
			"foreignField": "application_id",
			"localField":"_id",
			"as": "application_answers",
		},
	}
	// unwindStage := bson.D{{
	// 	"$lookup", 
	// 	bson.D{
	// 		{"from", "VisaApplications"},
	// 		{"localField",""}
	// 	}}}
	
	pipeline := []bson.M{pipe}
	var visaApplications []bson.M
	cursor, err := dao.Collection.Aggregate(dao.Context, pipeline )
	if err != nil {
		println(err.Error())
		return
	}
	err = cursor.All(dao.Context, &visaApplications)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%v",visaApplications[0]["application_answers"])
}