package controllers

import (
	"log"

	"github.com/grokkos/maple-syrup/api/models"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	{
		Log: 12,
	},
	{
		Log: 23,
	},
}
var batches = []models.Batch{
	{
		Dispatched:  true,
		BatchUserID: 1,
	},
	{
		Dispatched:  false,
		BatchUserID: 2,
	},
}
var roundups = []models.Roundup{
	{
		Amount:         30,
		RoundupBatchID: 1,
		RoundupUserID:  1,
	},
	{
		Amount:         15,
		RoundupBatchID: 1,
		RoundupUserID:  2,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Roundup{}, &models.Batch{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Batch{}, &models.Roundup{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {

		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed batches table: %v", err)
		}
	}
	for i, _ := range batches {

		err = db.Debug().Model(&models.Batch{}).Create(&batches[i]).Error
		if err != nil {
			log.Fatalf("cannot seed batches table: %v", err)
		}
	}

	for i, _ := range roundups {

		err = db.Debug().Model(&models.Roundup{}).Create(&roundups[i]).Error
		if err != nil {
			log.Fatalf("cannot seed roundup table: %v", err)
		}
	}

}
