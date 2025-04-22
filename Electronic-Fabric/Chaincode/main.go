package main

import (
	"log"

	"electronics/contract"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	electronicContract := new(contracts.ElectronicsContract)

	chaincode, err := contractapi.NewChaincode(electronicContract)

	if err != nil {
		log.Panicf("Could not create chaincode : %v", err)
	}

	err = chaincode.Start()

	if err != nil {
		log.Panicf("Failed to start chaincode : %v", err)
	}
}