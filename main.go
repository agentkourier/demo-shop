package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type checkoutRequest struct {
	Items       []Item `json:"items"`
	DiscountPct int    `json:"discountPct"`
}

func main() {
	http.HandleFunc("POST /checkout", func(w http.ResponseWriter, r *http.Request) {
		var req checkoutRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]int{"total": Total(req.Items, req.DiscountPct)})
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
