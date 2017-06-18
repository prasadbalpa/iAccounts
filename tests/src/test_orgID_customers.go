package main

import (
	"fmt"
	"iAccounts/datamodels"
)

func main() {
	var custtable []datamodels.Customer_table
	custtable = datamodels.GetCustomers_Given_ORG_ID("0b929804-8457-44bd-a2c6-a6d70b2a0c68")
	if custtable == nil {
		fmt.Println("Failed to fetch customers")
		return
	}
	fmt.Println(custtable[0])
}
