package tests

import (
	"log"
	"testing"

	"github.com/grokkos/maple-syrup/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestSaveRoundup(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatalf("Error user refreshing table %v\n", err)
	}

	err = populateTables()
	if err != nil {
		log.Fatalf("Error populating  table %v\n", err)
	}

	newRoundup := models.Roundup{
		Amount:         30,
		RoundupBatchID: 1,
		RoundupUserID:  1,
	}
	savedRoundup, err := newRoundup.SaveRoundup(server.DB)
	if err != nil {
		t.Errorf("Error while saving a user: %v\n", err)
		return
	}
	assert.Equal(t, newRoundup.Amount, savedRoundup.Amount)
	assert.Equal(t, newRoundup.RoundupBatchID, savedRoundup.RoundupBatchID)
	assert.Equal(t, newRoundup.RoundupUserID, savedRoundup.RoundupUserID)
}
func TestFindRoundups(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatalf("Error refreshing tables %v\n", err)
	}
	err = populateTables()
	if err != nil {
		log.Fatalf("Error populating  table %v\n", err)
	}

	roundups, err := roundupsInstance.FindAllRoundups(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, len(*roundups), 2)
}
