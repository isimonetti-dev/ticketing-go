package httpserver

import (
	"encoding/json"
	"net/http"
)

const contentTypeJSON = "application/json"

type healthResponse struct {
	Status string `json:"status"`
}

func Health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	payload, err := json.Marshal(body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(status)
	_, _ = w.Write(payload)
}
