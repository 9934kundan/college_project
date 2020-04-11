package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// get the car_info

func (t *SimpleChaincode) car_info(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var err error
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"car_info\"}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)
}

// get car manufacturer info

func (t *SimpleChaincode) car_manufacturer_info(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"car_manufacturer_info\"}}")
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(queryResults)

	return shim.Success(valAsbytes)
}

