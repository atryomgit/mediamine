package main

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}
type Response struct {
	Exchanges [][]int `json:"exchanges"`
}

func calculateCombinations(amount int, banknotes []int) [][]int {
	var result [][]int
	for i := 0; i < len(banknotes); i++ {
		for j := 0; j < len(banknotes)-i-1; j++ {
			if banknotes[j] < banknotes[j+1] {
				banknotes[j], banknotes[j+1] = banknotes[j+1], banknotes[j]
			}
		}
	}

	var findCombinations func(int, []int, int)
	findCombinations = func(remaining int, combo []int, start int) {
		if remaining == 0 {
			combination := make([]int, len(combo))
			copy(combination, combo)
			result = append(result, combination)
			return
		}
		for i := start; i < len(banknotes); i++ {
			if banknotes[i] <= remaining {
				findCombinations(remaining-banknotes[i], append(combo, banknotes[i]), i)
			}
		}
	}

	findCombinations(amount, []int{}, 0)
	return result
}

func exchangeHandler(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	combinations := calculateCombinations(req.Amount, req.Banknotes)

	resp := Response{Exchanges: combinations}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/exchange", exchangeHandler)
	http.ListenAndServe(":8080", nil)
}
