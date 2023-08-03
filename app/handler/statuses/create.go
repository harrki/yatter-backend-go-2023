package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

type Media struct {
	ID          int
	Description string
}

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Status string
	Medias []Media
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// status := new(object.Status)
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)
	status.AccountID = auth.AccountOf(r).ID
	status.Content = req.Status

	if err := h.sr.CreateStatus(ctx, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
