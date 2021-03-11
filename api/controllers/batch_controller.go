package controllers

import (
	"net/http"

	"github.com/grokkos/maple-syrup/api/models"
	"github.com/grokkos/maple-syrup/api/responses"
)

func (server *Server) GetBatches(w http.ResponseWriter, r *http.Request) {

	batch := models.Batch{}

	batches, err := batch.FindAllBatches(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, batches)
}
