package main

import (
	"fmt"
	"iAccounts/datamodels"
)

func main() {
	var vehtable []datamodels.Vehicle_table
	vehtable = datamodels.GetVehicles_Given_ORG_ID("0b929804-8457-44bd-a2c6-a6d70b2a0c68")
	if vehtable == nil {
		fmt.Println("Failed to fetch vehicles")
		return
	}
	fmt.Println(vehtable[0])
}
