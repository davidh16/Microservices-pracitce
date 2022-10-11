package initializers

import (
	"log"
	"logger-service/cmd/models"
)

func init() {
	ConnectToDB()
}

func Migrate() {
	err := DB.AutoMigrate(&models.Log{})
	if err != nil {
		log.Fatalf("could not migrate to the database: %v", err)
	}
}
