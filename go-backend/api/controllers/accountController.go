package controllers

import (
	"encoding/json"
	"gorestapi/api/models"
	"net/http"
)

// CreateAccount godoc
// @Summary Create new account
// @Description Add a new account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.Account true "Account Info"
// @Success 201 {object} models.Account
// @Router /accounts [post]
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount models.Account
	if err := json.NewDecoder(r.Body).Decode(&newAccount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement account creation logic...

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAccount)
}

// GetAccount godoc
// @Summary Get account details
// @Description Retrieve details of an account
// @Tags accounts
// @Produce json
// @Param iban path string true "Account IBAN"
// @Success 200 {object} models.Account
// @Router /accounts/{iban} [get]
func GetAccount(w http.ResponseWriter, r *http.Request) {
	// Dummy implementation, replace with real logic
	account := models.Account{
		IBAN: "DE89 3704 0044 0532 0130 00",
		BIC:  "COBADEFFXXX",
	}

	json.NewEncoder(w).Encode(account)
}
