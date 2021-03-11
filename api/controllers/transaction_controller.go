package controllers

import (
	"net/http"

	"github.com/grokkos/maple-syrup/api/models"
	"github.com/grokkos/maple-syrup/api/responses"
)

func (server *Server) GetTransactions(w http.ResponseWriter, r *http.Request) {

	transaction := models.Transaction{}

	transactions, err := transaction.FindAllTransactions(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, transactions)
}
