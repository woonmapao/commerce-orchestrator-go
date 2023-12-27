package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/woonmapao/commerce-orchestrator-go/models"
)

const (
	productServiceURL = "http://localhost:2002/products"
	stockUpdateURL    = "http://localhost:2002/products/update-stock"
)

type ProductResponse struct {
	Data    ProductData `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

type ProductData struct {
	Product models.Product `json:"product"`
}

func GetProductByID(productID int) (*models.Product, error) {
	// Make a GET request to the product-service API
	resp, err := http.Get(fmt.Sprintf("%s/%d", productServiceURL, productID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch product. Status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a ProductResponse struct
	var productResponse ProductResponse
	err = json.NewDecoder(resp.Body).Decode(&productResponse)
	if err != nil {
		return nil, err
	}

	// Return the product from the nested structure
	return &productResponse.Data.Product, nil
}

type UpdateStockRequest struct {
	Quantity int `json:"quantity"`
}

type UpdateStockResponse struct {
	Data    ProductData `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

func UpdateStock(productID, quantity int) (*models.Product, error) {

	// Prepare request body
	updateRequest := UpdateStockRequest{Quantity: quantity}
	requestBody, err := json.Marshal(updateRequest)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create PUT request
	req, err := http.NewRequest("PUT",
		fmt.Sprintf("%s/%d", stockUpdateURL, productID),
		bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to update stock. Status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into an UpdateStockResponse struct
	var updateResponse UpdateStockResponse
	err = json.NewDecoder(resp.Body).Decode(&updateResponse)
	if err != nil {
		return nil, err
	}

	// Return the updated product from the nested structure
	return &updateResponse.Data.Product, nil
}
