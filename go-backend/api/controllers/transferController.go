package controllers

import (
	"encoding/json"
	"gorestapi/api/models"
	"net/http"
)

// InitiateTransfer godoc
// @Summary Initiate a SEPA transfer
// @Description Perform a SEPA transfer from one account to another
// @Tags transfers
// @Accept json
// @Produce json
// @Param transfer body models.Transfer true "Transfer Info"
// @Success 200 {string} string "Transfer initiated successfully"
// @Router /transfers [post]
func InitiateTransfer(w http.ResponseWriter, r *http.Request) {
	var newTransfer models.Transfer
	if err := json.NewDecoder(r.Body).Decode(&newTransfer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Implement transfer logic...

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransfer)
}
