package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/grokkos/maple-syrup/api/controllers"
	"github.com/grokkos/maple-syrup/api/models"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}
var roundupsInstance = models.Roundup{}

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error obtaining .env %v\n", err)
	}
	Database()
	os.Exit(m.Run())
}

func Database() {

	var err error
	server.Initialize(os.Getenv("TEST_DB_DRIVER"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_PASSWORD"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_NAME"))

	if err != nil {
		log.Fatal("This is the log:", err)
	} else {
		fmt.Printf("We are connected to the database")
	}
}

func refreshTables() error { //refreshing the tables before the tests
	err := server.DB.Debug().DropTableIfExists(&models.Roundup{}, &models.Transaction{}, &models.Batch{}, &models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.Debug().AutoMigrate(&models.User{}, &models.Batch{}, &models.Transaction{}, &models.Roundup{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	log.Printf("refreshUserTable routine OK !!!")
	return nil
}

func populateTables() error { //seeding the tables with data to assert in the tests
	users := []models.User{
		models.User{
			Name: "Tina",
		},
		models.User{
			Name: "John",
		},
	}

	for i := range users {
		err := server.DB.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	log.Printf("populate users table successfully")

	batches := []models.Batch{
		models.Batch{
			Dispatched:  false,
			BatchUserID: 1,
		},
	}
	for i := range batches {
		err := server.DB.Debug().Model(&models.Batch{}).Create(&batches[i]).Error
		if err != nil {
			return err
		}
	}
	log.Printf("populate users table successfully")

	roundups := []models.Roundup{
		models.Roundup{
			Amount:         50,
			RoundupBatchID: 1,
			RoundupUserID:  2,
		},
		models.Roundup{
			Amount:         10,
			RoundupBatchID: 1,
			RoundupUserID:  2,
		},
	}
	for i := range roundups {
		err := server.DB.Debug().Model(&models.Roundup{}).Create(&roundups[i]).Error
		if err != nil {
			return err
		}
	}
	log.Printf("populate roundups table successfully")
	return nil

}
