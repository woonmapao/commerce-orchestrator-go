package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/woonmapao/commerce-orchestrator-go/models"
)

const orderDetailServiceURL = "http://localhost:4040/order-details"

type CreateOrderDetailRequest struct {
	OrderID   int     `json:"orderId"`
	ProductID int     `json:"productId"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type CreateOrderDetailResponse struct {
	Data    OrderDetailData `json:"data"`
	Message string          `json:"message"`
	Status  string          `json:"status"`
}

type OrderDetailData struct {
	OrderDetail models.OrderDetail `json:"order_detail"`
}

func CreateOrderDetail(orderID, productID, quantity int, subtotal float64) error {
	// Prepare the create order detail request body
	createOrderDetailRequest := CreateOrderDetailRequest{
		OrderID:   orderID,
		ProductID: productID,
		Quantity:  quantity,
		Subtotal:  subtotal,
	}
	requestBody, err := json.Marshal(createOrderDetailRequest)
	if err != nil {
		return err
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new POST request to create an order detail
	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/", orderDetailServiceURL),
		bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set the request header
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create order detail. Status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a CreateOrderDetailResponse struct
	var createOrderDetailResponse CreateOrderDetailResponse
	err = json.NewDecoder(resp.Body).Decode(&createOrderDetailResponse)
	if err != nil {
		return err
	}

	return nil
}
