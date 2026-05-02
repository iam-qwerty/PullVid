package api

import (
	"encoding/json"
	"net/http"
)

type videoInfoRequest struct {
	URL string `json:"url"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(dst)
}
