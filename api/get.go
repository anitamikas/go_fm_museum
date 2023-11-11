package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mikasanita.com/go/fm-museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // by def http expect that we send text, so we let it know its json
	id := r.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {
			json.NewEncoder(w).Encode(data.GetAll()[finalId])
		} else {
			http.Error(w, "Invalid exhibition id", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}
}
