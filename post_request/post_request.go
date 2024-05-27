package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "http://localhost:8080/exchange"
	requestBody, err := json.Marshal(map[string]interface{}{
		"amount": 400,
		"banknotes": []int{
			5000, 2000, 1000, 500, 200, 100, 50,
		},
	})

	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}
	fmt.Println(result)
}
