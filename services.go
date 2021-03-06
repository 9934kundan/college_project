package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


// Main
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting the chaincode: %s", err)
	}
}

// Init initializes chaincode

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "addManufacturer" { //create a new user
		return t.addmanufacturer(stub, args)
	} else if function == "addUser" { //read a user
		return t.adduser(stub, args)
	} else if function == "addInsurer" { //initialize a course for a user
		return t.addinsurer(stub, args)
	} else if function =="addDealer {
           return t.addDealer(stub, args)
        }else if function =="addRTA {
           return t.addDealer(stub, args)
        }  else if function == "car_info" {
	   return t.car_info(stub, args)
        } else if function == "car_manufacturer_info"
           return t.car_manufacturer_info(stub, args)
        }else if function == "addCartoLedger"
           return t.addCartoLedger(stub, args)
        }else if function == "transferToDealer"
           return t.transferToDealer(stub, args)
        }else if function == "sellToCustomer"
           return t.sellToCustomer(stub, args)
        }else if function == "sellCar"
           return t.sellCar(stub, args)
        }else if function == "insuranceCar"
           return t.insuranceCar(stub, args)
        }else if function == "rtaApprove"
           return t.rtaApprove(stub, args)
        }else if function == "sellToAnotherCustomer"
           return t.rtaApprove(stub, args)
        }
   fmt.Println("invoke did not find function: " + function) //error
	return shim.Error("Received unknown function call")
  
  }
