package initializers

import (
	"createuser-service/cmd/models"
	"log"
)

func init() {
	ConnectToDB()
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("could not migrate to the database: %v", err)
	}
}
