package environ

import (
	"os"

	"github.com/golangcimri/api/database"
	"github.com/golangcimri/api/globals"
	"github.com/joho/godotenv"
)

func Init(check string) {

	var err error

	//Load .env file
	if check == "test" {
		err = godotenv.Load("../.env")
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
		ApiPort:  os.Getenv("API_PORT"),
		ApiToken: os.Getenv("API_TOKEN"),
		XToken:   os.Getenv("X_TOKEN"),
		Database: db,
	}
}
