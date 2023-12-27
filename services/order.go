package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const userServiceURL = "http://localhost:9009/users"

type CreateUserOrderRequest struct {
	UserID      uint      `json:"UserID"`
	OrderDate   time.Time `json:"OrderDate"`
	TotalAmount float64   `json:"TotalAmount"`
	Status      string    `json:"Status"`
}

type CreateUserOrderResponse struct {
	Data    UserData `json:"data"`
	Message string   `json:"message"`
	Status  string   `json:"status"`
}

func CreateOrder(userID uint, totalAmount float64) (uint, error) {
	// Prepare the create order request body
	createOrderRequest := CreateUserOrderRequest{
		UserID:      userID,
		TotalAmount: totalAmount,
		OrderDate:   time.Now(),
		Status:      "pending",
	}
	requestBody, err := json.Marshal(createOrderRequest)
	if err != nil {
		return 0, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new POST request to create an order
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/", userServiceURL),
		bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to create order. Status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a CreateUserOrderResponse struct
	var createOrderResponse CreateUserOrderResponse
	err = json.NewDecoder(resp.Body).Decode(&createOrderResponse)
	if err != nil {
		return 0, err
	}

	// Return the order ID from the nested structure
	return createOrderResponse.Data.User.ID, nil
}
