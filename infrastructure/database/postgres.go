package database

import (
	"fmt"
	"log"

	"github.com/FianGumilar/e-commerce/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbPg *gorm.DB

func GetDbPostgres(conf *config.AppConfig) *gorm.DB {
	var err error

	dsn := fmt.Sprintf(
		"host = %s"+
			"port = %s"+
			"user = %s"+
			"pass = %s"+
			"name = %s"+
			"ssl = %s"+
			conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.User,
		conf.Postgres.Pass,
		conf.Postgres.Name,
		conf.Postgres.Ssl,
	)

	DbPg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed connect to Database: %s", err)
	}
	return DbPg
}
