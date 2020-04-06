package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


type customer struct {
  ObjectType string `json:"objectType"`
  CustomerId string `json:"customerid"`
  CustomerName string `json:"customername"`
  }
