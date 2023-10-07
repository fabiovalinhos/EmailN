package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {

	dsn := "host=localhost user=emailn_dev password=cebola dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contacts{})

	return db
}
