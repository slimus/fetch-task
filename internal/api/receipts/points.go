package receipts

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/slimus/fetch-task/internal/rewards"
	"github.com/slimus/fetch-task/internal/storage/inmemory"
	"github.com/slimus/fetch-task/internal/utils"
)

type RewardResponse struct {
	Points int `json:"points"`
}

func (a *App) HandleRecieptsPoints(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("method %s is not allowed here\n", r.Method)
		utils.JSONError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	if id == "" {
		log.Println("got empty id")
		utils.JSONError(w, "id is required", http.StatusBadRequest)
		return
	}

	reciept, err := a.db.GetById(id)
	if err != nil {
		if errors.Is(err, inmemory.ErrRecieptNotFound) {
			utils.JSONError(w, "Reciept not found", http.StatusNotFound)
			return
		}
		log.Printf("getting reciept by id failed: %v\n", err)
		utils.JSONError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rewards := rewards.CalculateRewards(reciept)

	json.NewEncoder(w).Encode(RewardResponse{Points: rewards})
}
