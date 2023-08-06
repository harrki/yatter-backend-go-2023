package timelines

import (
	"encoding/json"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/object"
)

// Handle request for `GET /v1/timelines`
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(object.SearchRequest)
	if is_only_media, err := strconv.ParseBool(r.URL.Query().Get("only_media")); err != nil {
		req.OnlyMedia = false
	} else {
		req.OnlyMedia = is_only_media
	}

	if max_id, err := strconv.Atoi(r.URL.Query().Get("max_id")); err != nil {
		req.MaxID = -1
	} else {
		req.MaxID = max_id
	}

	if since_id, err := strconv.Atoi(r.URL.Query().Get("since_id")); err != nil {
		req.SinceID = -1
	} else {
		req.SinceID = since_id
	}

	if limit, err := strconv.Atoi(r.URL.Query().Get("limit")); err != nil {
		req.Limit = 40
	} else {
		if limit > 80 {
			req.Limit = 80
		} else {
			req.Limit = limit
		}
	}

	if timeline, err := h.tr.Search(ctx, req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if timeline == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else {
		json.NewEncoder(w).Encode(timeline)
	}
}
