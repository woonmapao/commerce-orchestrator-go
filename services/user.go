package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/woonmapao/commerce-orchestrator-go/models"
)

type UserResponse struct {
	Data    UserData `json:"data"`
	Message string   `json:"message"`
	Status  string   `json:"status"`
}

type UserData struct {
	User models.User `json:"user"`
}

func GetUserByID(userID uint) (*models.User, error) {
	// Make a GET request to the user-service API
	resp, err := http.Get(fmt.Sprintf("%s/%d", userServiceURL, userID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user. Status code: %d", resp.StatusCode)
	}

	// Decode the JSON response into a UserResponse struct
	var userResponse UserResponse
	err = json.NewDecoder(resp.Body).Decode(&userResponse)
	if err != nil {
		return nil, err
	}

	// Return the user from the nested structure
	return &userResponse.Data.User, nil
}
