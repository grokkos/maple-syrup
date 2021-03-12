package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/grokkos/maple-syrup/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

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

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {

	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Batch{}, &models.Roundup{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()
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

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
