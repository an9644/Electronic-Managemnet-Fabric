package contracts

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// ElectronicsContract manages CRUD for electronic items
type ElectronicsContract struct {
	contractapi.Contract
}

type ElectronicItem struct {
	AssetType   string `json:"assetType"`
	ItemID      string `json:"itemId"`
	Name        string `json:"name"`
	Model       string `json:"model"`
	Color       string `json:"color"`
	Manufacturer string `json:"manufacturer"`
	Status      string `json:"status"`
}

// ItemExists checks if an item exists in the world state
func (c *ElectronicsContract) ItemExists(ctx contractapi.TransactionContextInterface, itemID string) (bool, error) {
	data, err := ctx.GetStub().GetState(itemID)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return data != nil, nil
}

// CreateItem allows only the manufacturer to create a new electronic item
func (c *ElectronicsContract) CreateItem(ctx contractapi.TransactionContextInterface, itemID, name, model, color, manufacturer string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "manufacturer-electronics-com" {

		exists, err := c.ItemExists(ctx, itemID)
		if err != nil {
			return "", err
		}
		if exists {
			return "", fmt.Errorf("item %s already exists", itemID)
		}

		item := ElectronicItem{
			AssetType:   "electronic",
			ItemID:      itemID,
			Name:        name,
			Model:       model,
			Color:       color,
			Manufacturer: manufacturer,
			Status:      "Manufactured",
		}

		bytes, _ := json.Marshal(item)
		err = ctx.GetStub().PutState(itemID, bytes)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("successfully added item %s", itemID), nil

	} else {
		return "", fmt.Errorf("unauthorized: %s cannot perform this action", clientOrgID)
	}
}

// ReadItem returns the electronic item by ID
func (c *ElectronicsContract) ReadItem(ctx contractapi.TransactionContextInterface, itemID string) (*ElectronicItem, error) {
	bytes, err := ctx.GetStub().GetState(itemID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if bytes == nil {
		return nil, fmt.Errorf("item %s does not exist", itemID)
	}

	var item ElectronicItem
	err = json.Unmarshal(bytes, &item)
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %v", err)
	}

	return &item, nil
}

// DeleteItem allows only the manufacturer to delete an electronic item
func (c *ElectronicsContract) DeleteItem(ctx contractapi.TransactionContextInterface, itemID string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID == "manufacturer-electronics-com" {
		exists, err := c.ItemExists(ctx, itemID)
		if err != nil {
			return "", err
		}
		if !exists {
			return "", fmt.Errorf("item %s does not exist", itemID)
		}

		err = ctx.GetStub().DelState(itemID)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("item %s deleted", itemID), nil
	} else {
		return "", fmt.Errorf("unauthorized: %s cannot perform this action", clientOrgID)
	}
}
