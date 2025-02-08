package receipts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/slimus/fetch-task/internal/utils"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type RecieptProcessRequest struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type RecieptProcessResponse struct {
	ID string `json:"id"`
}

func (a *App) HandleRecieptsProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("method %s is not allowed here\n", r.Method)
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var request RecieptProcessRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("decoding request failed: %v\n", err)
		utils.JSONError(w, "Invalid input", http.StatusBadRequest)
		return
	}
	model, err := ConvertRequestRecieptToDB(request)
	if err != nil {
		log.Printf("converting request to model failed: %v\n", err)
		utils.JSONError(w, "Validation Error", http.StatusBadRequest)
		return
	}

	uuid, err := a.db.SaveReciept(model)

	if err != nil {
		log.Printf("saving reciept failed: %v\n", err)
		utils.JSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("New reciept with id=%s has been added\n", uuid.String())
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(RecieptProcessResponse{ID: uuid.String()})
}
