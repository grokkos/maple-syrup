package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/grokkos/maple-syrup/api/models"
	"github.com/grokkos/maple-syrup/api/responses"
)

func (server *Server) GetAllBatches(w http.ResponseWriter, r *http.Request) {

	batch := models.Batch{}

	batches, err := batch.FindAllBatches(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, batches)
}

func (server *Server) GetBatchesByUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Print("Fail to parse!")
	}

	filter := new(models.Filter) //filter struct to parse the user id value
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		fmt.Print("Fail to decode!")
	}

	fmt.Printf("%+v", filter)
	batch := models.Batch{}
	batches, err := batch.FindBatchesByUserId(server.DB, uint32(filter.UserID))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, batches)
}
