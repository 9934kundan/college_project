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
  
  type dealer struct {
  ObjectType string `json:"docType"`
  DealerId string `json:"dealerid"`
  AgencyNmae string `json:"agnecyname"`
  }
  
  type insurer struct {
  ObjectType string `json:"objectType"`
  InsurerId string `json:"insurerid"`
  InsuranceName string `json:"insurancename"`
  }
  
  type rta struct {
  ObjectType string `json:"objectType"`
  RTAId string `json:"rtaid"`
  RTAName string `json:"rtaname"`
  }
  
  type manufacturer struct {
  ObjectType string `json:"objectType"`
  ManufacturerId string `json:"manufacturerid"`
  ManufacturerName string `json:"manufacturername"`
  }
  
