package globals

import (
	"github.com/golangcimri/worker/database"
)

var Variables *GlobalsStruct

type GlobalsStruct struct {
	AccessKey  string
	SecretKey  string
	Region     string
	BucketName string
	Database   *database.Database
}
