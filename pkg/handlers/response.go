package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		respondWithError(w, 500, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		return
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func parseJsonToStruct(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}
	defer func() {
		err := r.Body.Close()
		if err != nil {
			return
		}
	}()
	return nil
}
