package globals

import (
	"github.com/golangcimri/api/database"
)

var Variables *GlobalsStruct

type GlobalsStruct struct {
	ApiPort  string
	ApiToken string
	XToken   string
	Database *database.Database
}
