package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/grokkos/maple-syrup/api/models"
	"github.com/grokkos/maple-syrup/api/responses"
)

func (server *Server) CreateRoundup(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	roundup := models.Roundup{}
	err = json.Unmarshal(body, &roundup)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	roundupCreated, err := roundup.SaveRoundup(server.DB)

	if err != nil {
		formattedError := responses.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, roundupCreated.ID))
	responses.JSON(w, http.StatusCreated, roundupCreated)
}
