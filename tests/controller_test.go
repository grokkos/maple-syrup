package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grokkos/maple-syrup/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetRoundups(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatalf("Error refreshing tables %v\n", err)
	}

	err = populateTables()
	if err != nil {
		log.Fatalf("Error populating tables %v\n", err)
	}

	req, err := http.NewRequest("GET", "/roundups", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetRoundups)
	handler.ServeHTTP(rr, req)

	var roundups []models.Roundup
	err = json.Unmarshal([]byte(rr.Body.String()), &roundups)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(roundups), 2)
}

func TestCreateRoundups(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatal(err)
	}

	err = populateTables()
	if err != nil {
		log.Fatalf("Error populating tables %v\n", err)
	}
	samples := []struct {
		inputJSON     string
		statusCode    int
		errorMessage  string
		amount        int
		roundupUserID int
	}{
		{

			inputJSON:  `{"amount": 101, "roundup_user_id": 1}`,
			statusCode: 201,

			amount:        101,
			roundupUserID: 1,
		},
		{

			inputJSON:  `{"amount": 301, "roundup_user_id": 1}`,
			statusCode: 201,

			amount:        301,
			roundupUserID: 1,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/roundups", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateRoundup)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetBatchesGeneratedAfterPostingRoundups(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatal(err)
	}
	err = populateTables()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON     string
		amount        int
		roundupUserID int
	}{
		{

			inputJSON:     `{"amount": 101, "roundup_user_id": 1}`,
			amount:        101,
			roundupUserID: 1,
		},
		{

			inputJSON:     `{"amount": 301, "roundup_user_id": 1}`,
			amount:        301,
			roundupUserID: 1,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/roundups", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateRoundup)
		handler.ServeHTTP(rr, req)
	}

	req, err := http.NewRequest("GET", "/batches", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetAllBatches)
	handler.ServeHTTP(rr, req)

	var batches []models.Batch
	err = json.Unmarshal([]byte(rr.Body.String()), &batches)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(batches), 3)
}

func TestGetTransactionsGeneratedAfterPostingRoundups(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatal(err)
	}
	err = populateTables()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON     string
		amount        int
		roundupUserID int
	}{
		{

			inputJSON:     `{"amount": 151, "roundup_user_id": 1}`,
			amount:        151,
			roundupUserID: 1,
		},
		{

			inputJSON:     `{"amount": 21, "roundup_user_id": 2}`,
			amount:        21,
			roundupUserID: 2,
		},
		{

			inputJSON:     `{"amount": 90, "roundup_user_id": 2}`,
			amount:        90,
			roundupUserID: 2,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/roundups", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateRoundup)
		handler.ServeHTTP(rr, req)
	}

	req, err := http.NewRequest("GET", "/transactions", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetTransactions)
	handler.ServeHTTP(rr, req)

	var transactions []models.Transaction
	err = json.Unmarshal([]byte(rr.Body.String()), &transactions)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(transactions), 2)
}

func TestGetBatchesByUserId(t *testing.T) {

	err := refreshTables()
	if err != nil {
		log.Fatal(err)
	}
	err = populateTables()
	if err != nil {
		log.Fatal(err)
	}
	samples := []struct {
		inputJSON     string
		amount        int
		roundupUserID int
	}{
		{

			inputJSON:     `{"amount": 151, "roundup_user_id": 1}`,
			amount:        151,
			roundupUserID: 1,
		},
		{

			inputJSON:     `{"amount": 211, "roundup_user_id": 1}`,
			amount:        211,
			roundupUserID: 1,
		},
		{

			inputJSON:     `{"amount": 90, "roundup_user_id": 2}`,
			amount:        90,
			roundupUserID: 2,
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/roundups", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateRoundup)
		handler.ServeHTTP(rr, req)
	}

	req, err := http.NewRequest("GET", "/batchlist?id=1", nil)
	if err != nil {
		t.Errorf("this is the error: %v\n", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetBatchesByUser)
	handler.ServeHTTP(rr, req)

	var batches []models.Batch
	err = json.Unmarshal([]byte(rr.Body.String()), &batches)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v\n", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(batches), 2)
}
