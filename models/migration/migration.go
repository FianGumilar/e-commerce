package migration

import (
	"log"

	"github.com/FianGumilar/e-commerce/infrastructure/database"
	"github.com/FianGumilar/e-commerce/models/entity"
)

func Migration() {
	err := database.DbPg.AutoMigrate(
		&entity.Admin{},
	)

	if err != nil {
		log.Printf("Failed to migrate database: %s", err)
	}
}
