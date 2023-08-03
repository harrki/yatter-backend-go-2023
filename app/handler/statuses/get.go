package statuses

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if status, err := h.sr.FindByID(ctx, id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if status == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else {
		json.NewEncoder(w).Encode(status)
	}
}
