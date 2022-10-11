package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var err error

func ConnectToDB() {
	dsn := os.Getenv("DSN")
	config := &gorm.Config{}
	//config.Logger = logger.Default.LogMode(logger.Info)

	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
}
