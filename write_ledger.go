package main

import (
	"encoding/json"
	"fmt"
        
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

fun (t *SimpleChaincode) addCarOnLedger(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	  var err error
	  carjson := m_car{}
	  err = json.Unmarshal([]byte(args[0]), &carjson)
	  if err != nil {
		  return shim.Error("Failed to unmarshall")
	  }

	  carid := carjson.CarId
	  chessisno := carjson.ChessisNo
	  modelid := carjson.ModelId
	  carowner := carjson.CarOwner
	  newowner := carjson.NewOnwer
	  color := carjson.Color
	  companyname := carjson.CompanyName

	  valAsBytes, err := stub.GetState(carid)
	  if err != nil {
		  return shim.Error("Failed to get the state")
	  }	else if valAsBytes != nil {
		  return shim.Error("Car already registerd with this id")
	  }

	  objectType := "m_car"
	  m_car := &m_car(objectType, carid, chessisno, modelid, carowner, newowner, color, companyname)
	  userJSON, err := json.Marshal(m_car)
	  if err != nil {
		  return shim.Error("Failed to Marshal")
	  } 
	  
	 // save the user state
	 err = stub.PutState(userJSON)
	 if err != nil {
		 return shim.Error("Failed to put the state")
	 }
	 
	 return shim.Success("OK")
  }
