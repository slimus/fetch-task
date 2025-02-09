package receipts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/slimus/fetch-task/internal/utils"
)

type Item struct {
	ShortDescription string `json:"shortDescription" validate:"required"`
	Price            string `json:"price" validate:"required"`
}

type RecieptProcessRequest struct {
	Retailer     string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"dive,required"`
	Total        string `json:"total" validate:"required"`
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
		utils.JSONError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	err := validate.Struct(request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		if len(validationErrors) > 0 {
			log.Printf("validation failed: %v\n", validationErrors)
			utils.JSONError(w, "The receipt is invalid.", http.StatusBadRequest)
			return
		}
	}

	model, err := ConvertRequestRecieptToDB(request)
	if err != nil {
		log.Printf("converting request to model failed: %v\n", err)
		utils.JSONError(w, "The receipt is invalid.", http.StatusBadRequest)
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
