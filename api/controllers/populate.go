package controllers

import (
	"log"

	"github.com/grokkos/maple-syrup/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Name: "Ger",
	},
}
var batches = []models.Batch{
	{
		BatchUserID: 1,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Roundup{}, &models.Transaction{}, &models.Batch{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Batch{}, &models.Transaction{}, &models.Roundup{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.User{}).Create(&users[0]).Error
	if err != nil {
		log.Fatalf("cannot seed batches table: %v", err)
	}

	err = db.Debug().Model(&models.Batch{}).Create(&batches[0]).Error
	if err != nil {
		log.Fatalf("cannot seed batches table: %v", err)
	}

}
