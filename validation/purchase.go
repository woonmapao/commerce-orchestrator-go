package validation

import (
	"errors"

	"github.com/woonmapao/commerce-orchestrator-go/models"
)

func ValidatePurchaseOrderData(purchaseOrder *models.PurchaseOrder) error {

	// Required user id
	if purchaseOrder.UserID == 0 {
		return errors.New("UserID is required")
	}

	if len(purchaseOrder.Products) == 0 {
		return errors.New("Products are required")
	}

	return nil

}
