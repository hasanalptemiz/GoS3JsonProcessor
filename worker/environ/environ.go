package environ

import (
	"os"

	"github.com/golangcimri/worker/database"
	"github.com/golangcimri/worker/globals"
	"github.com/joho/godotenv"
)

func Init(check string) {

	var err error

	// Load the environment variables
	if check == "test" {
		err = godotenv.Load("../.env")
		if err != nil {
			panic(err)
		}
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		panic(err)
	}

	// Connect to MongoDB
	db := database.New()
	err = db.Connect(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DB"))
	if err != nil {
		panic(err)
	}

	// Set the global variables
	globals.Variables = &globals.GlobalsStruct{
		AccessKey:  os.Getenv("ACCESS_KEY"),
		SecretKey:  os.Getenv("SECRET_KEY"),
		Region:     os.Getenv("REGION"),
		BucketName: os.Getenv("BUCKET_NAME"),
		Database:   db,
	}
}
