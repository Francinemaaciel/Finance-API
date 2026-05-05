package handler

import (
	"encoding/json"
	"finance-api/internal/service"
	"net/http"
)

func GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	exchange, err := service.GetExchangeRate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchange)
}
