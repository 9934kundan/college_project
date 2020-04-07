package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type manufacturer struct {
   ObjectType string `json:"docType"`
   ManufacturerId string `json:"manfacturerid"`
   CompanyName string `json:"companyname"`
 }
 


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type m_car struct {
  ObjectType     string `json:"docType"` //docType is used to distinguish the various types of objects in state database
  CarId          string `json:"carid"`
  ChessisNo      string `json:"chessisno"`
  ModelId        string `json:"modelid"`
  CarOwner       string `json:"carowner"`
  NewOwner       string `json:"newowner"`
  Color          string `json:"color"`
  CompanyName    string `json:"companyname"`
}



