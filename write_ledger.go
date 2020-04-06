package main

import (
	"encoding/json"
	"fmt"
        
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// function to create new manufacturer

fun (t *SimpleChaincode) init_manufacturer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   
   var err error
   manufacturerjson := manufacture{}
   err = json.Unmarshal([]byte(args[0], &manufacturer)
   if err != nil {
         fmt.Println(err.Error())
   }
   
   manufacturerid := manufacturerjson.ManufacturerId
   manufacturername := manufacturerjson.ManufacturerName
   
   
   // check if user already exist
   
   userAsBytes, err := stub.GetState(manufacturerid)
   if err != nil {
         fmt.Println("Already registered with this id " + err.Error())
         return shim.Error("Already registered with this id " + manufacturerid)
   }
   
   
   
   
   
