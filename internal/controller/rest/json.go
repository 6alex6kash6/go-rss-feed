package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"log/slog"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		slog.Info("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		slog.Info("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func ParseBody(r *http.Request, input interface{}) error {
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(input); err != nil {
		return errors.New("Error while decoding body")
	}
	return nil
}
