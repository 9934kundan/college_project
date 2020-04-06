package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type insurer struct {
  ObjectType string `json:"objectType"`
  InsurerId string `json:"insurerid"`
  InsuranceName string `json:"insurancename"`
  }
  
