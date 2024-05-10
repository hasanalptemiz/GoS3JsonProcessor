package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database struct
type Database struct {
	Db         *mongo.Database
	Collection struct {
		Records string
	}
}

// New database function
func New() *Database {
	return &Database{
		Collection: struct {
			Records string
		}{
			Records: "records",
		},
	}
}
func (db *Database) Connect(mongoURI, mongoDB string) error {
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	database := client.Database(mongoDB)

	// Compound index
	compoundIndex := mongo.IndexModel{
		Keys: bson.M{
			"id": 1, // index in ascending order
		},
	}
	_, err = database.Collection(db.Collection.Records).Indexes().CreateOne(context.Background(), compoundIndex)
	if err != nil {
		return err
	}

	db.Db = database

	return nil
}

func (db *Database) GetOne(collection string, id int) *mongo.SingleResult {
	filter := bson.M{"id": id}
	c := db.Db.Collection(collection)
	return c.FindOne(context.Background(), filter)
}
