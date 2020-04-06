package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


type dealer struct {
  ObjectType string `json:"docType"`
  DealerId string `json:"dealerid"`
  AgencyNmae string `json:"agnecyname"`
  }
