package dao

import (
	"context"
	"errors"
	"lovenation_bend/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FactoryDAO struct {
	db          *mongo.Client
	ctx         context.Context
	Collections map[string]*mongo.Collection
}
 
// creates and setups a new dao factory
func InitializeFactory(db *mongo.Client, ctx context.Context) *FactoryDAO {
	collectionList := []string{"Users", "VisaApplicationAnswers", "VisaApplications"}
	collections := make(map[string]*mongo.Collection)
	for _, key := range collectionList {
		col := configs.GetCollection(db, key)
		collections[key] = col
		// collections = append(collections[], map[string]*mongo.Collection{"":col})
	}
	return &FactoryDAO{
		db:          db,
		ctx:         context.TODO(),
		Collections: collections,
	}
}

func (dao *FactoryDAO) Insert(key string, data interface{}) error {
	collection, ok := dao.Collections[key]
	if !ok {
		return errors.New("invalid collection")
	}

	c, _ := bson.Marshal(data)

	_, err := collection.InsertOne(dao.ctx, c)
	return err
}

// insert many data into a single collection
func (dao *FactoryDAO) InsertMany(key string, data []interface{}) error {
	collection, ok := dao.Collections[key]
	if !ok {
		return errors.New("invalid collection")
	}

	// c, _ := bson.Marshal(data)

	_, err := collection.InsertMany(dao.ctx, data)
	return err
}


// update


// get 

// get by id

// soft  delete 
