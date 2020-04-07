package main

import (
	"encoding/json"
	"fmt"
        
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// function to create new manufacturer

fun (t *SimpleChaincode) addManufacturer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   
   var err error
   manufacturerjson := manufacture{}
   err = json.Unmarshal([]byte(args[0]), &manufacturerjson)
   if err != nil {
         fmt.Println(err.Error())
   }
   
   manufacturerid := manufacturerjson.ManufacturerId
   companyname := manufacturerjson.CompanyName

   // check if manufacturer already exist
   
   userAsBytes, err := stub.GetState(manufacturerid)
   if err != nil {
	   return shim.Error(err.Error())
   } else if userAsBytes != nil {
         fmt.Println("Already registered with this id " + err.Error())
         return shim.Error("Already registered with this id " + manufacturerid)
   }
			
  // create new id
			
  objectType := "manufacturer"
  manufacturer := &manufacturer(objectType, manufacturerid, companyname)
  manufacturerJSON, err := json.Marshal(manufacturer)
  if err != nil {
	  return shim.Error(err.Error())
  }
  
  
  // save to the state
  
  err = stub.PutState(manufactutrerid, manufacturerJSON)
  if err != nil {
	  return shim.Error(err.Error())
  }

  return shim.Success("OK")
  
  }

fun (t *SimpleChaincode) addCustomer(stub shim.ChaincodeStubInterface, args []string) pb.response {

	var err error
	userjson := customer{}
	err = json.Unmarshal([]byte(args[0]), &userjson)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	customerid := userjson.CustomerId
	customername := userjson.CustomerName
	
	valAsBytes, err := stub.GetState(customerid)
	if err != nil {
		return shim.Error("Failed to get the state")
	} else if (valAsBytes != nil) {
		return shim.Error("Customer with this Id already registered")
	}
	
	// create customer object and marshall to json
	
	objectType := "customer"
	customer := &customer(objectType, customerid, customername)
	customerJSON, err := json.Marshal(customer)
	if err != nil {
		return shim.Error(err.Error())
	}

	// save customer to state

	err = stub.PutState(customerid, customerJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success("OK")
}

fun (t *SimpleChaincode) addInsurer(stub shim.ChaincodeStubInterface, args []string) pb.response {
	
	addInsurerJson : = addInsurer{}
	err := json.Unmarshal([]byte(args[0]), &addInsurerJson)
	if err != nil {
		return shim.Error(err.Error)
	}

	insurerid := addInsurerJson.InsurerId
	insurername := addInsurerJson.InsurerName

	valAsBytes, err := stub.GetState(insurerid)
	if err != nil {
		return shim.Error(err.Error)
	} else if valAsBytes != nil {
		return shim.Error("This insurer already registered")
	}

	// create user object and marshal to json

	objectType := "insurer"
	insurer := &insurer(objectType, insurerid, insurername)

	insurerJSON, err := json.Marshal(insurer)
	if err != nil {
		return shim.Error(err.Error())
	} 

	// save user to state

	err := stub.PutState(insurerid, insurer)
	if err := nil {
		return shim.Error("Failed to save the state")
	}

	return shim.Success("OK")

} 

fun (t *SimpleChaincode) addDealer(stub shim.ChaincodeStubInterface, args []string) pb.response {

	var err error
	dealerjson := dealer{}
	err := json.Unmarshal([]byte(args[0]), &dealerjson)
	if err != nil {
		return shim.Error("Failed to Unmarshal")
	}

	dealerid := dealerjson.DealerId
	agencyname := dealerjson.AgencyName
	
	valAsBytes, err := stub.GetState(dealerid)
	if err != nil {
		return shim.Error("Failed to get the state")
	} else if valAsBytes != nil {
		shim.Error("Dealer already registered with this ID")
	}

	objectType := "dealer"

	dealer := &dealer(objectType, dealerid, agencyname)
	userJSON, err = json.Marshal(dealer)
	if 	err != nil {
		return shim.Error(err.Error())
	}
	
	err = stub.PutState(dealerid, userJSON)
	if err != nil {
		return shim.Error("Failed to put the state")
	}

	return shim.Success("OK")

}


	

   
   
   
   
   
   
