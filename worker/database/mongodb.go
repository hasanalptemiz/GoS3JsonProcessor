package database

import (
	"context"

	"github.com/golangcimri/worker/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database struct
type Database struct {
	Db         *mongo.Database
	Collection struct {
		Records string
		Tests   string
	}
}

// New database function
func New() *Database {
	return &Database{
		Collection: struct {
			Records string
			Tests   string
		}{
			Records: "records",
			Tests:   "tests",
		},
	}
}

// Connect function
func (db *Database) Connect(mongoURI, mongoDB string) error {
	// Connect to MongoDB
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

	// Create the index
	_, err = database.Collection(db.Collection.Records).Indexes().CreateOne(context.Background(), compoundIndex)
	if err != nil {
		return err
	}

	db.Db = database

	return nil
}

// InsertRecord function
func (db *Database) BulkInsertRecords(collection string, records []models.Record) error {
	c := db.Db.Collection(collection)

	// Prepare the documents
	var documents []interface{}
	for _, record := range records {
		documents = append(documents, record)
	}

	// Insert the documents
	_, err := c.InsertMany(context.Background(), documents)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) GetOne(collection string, id int) *mongo.SingleResult {
	c := db.Db.Collection(collection)

	filter := bson.M{"id": id}

	return c.FindOne(context.Background(), filter)
}
