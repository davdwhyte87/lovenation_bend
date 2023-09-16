package dao

import (
	"context"
	"lovenation_bend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)


type UserDAO struct {
	Collection *mongo.Collection
	Context context.Context
}

// get the role for a particular user based on id 
func (userDAO *UserDAO) GetUserRole(roleID int) ( models.Role){
	// 
	var role models.Role

	userDAO.Collection.FindOne(userDAO.Context, bson.M{"_id":roleID} ).Decode(&role)
	return role
}

// insert 


// update 

// get 

// get by id

// delete 