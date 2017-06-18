package main

import (
	"fmt"
	"iAccounts/datamodels"
)

func main() {
	var prodtable []datamodels.Product_table
	prodtable = datamodels.GetProducts_Given_ORG_ID("0b929804-8457-44bd-a2c6-a6d70b2a0c68")
	if prodtable == nil {
		fmt.Println("Failed to fetch products")
		return
	}
	fmt.Println(prodtable[0])
}
