package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handle request for `GET /v1/accounts/{username}`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if account, err := h.ar.FindByUsername(ctx, username); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if account == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else {
		json.NewEncoder(w).Encode(account)
	}
}
