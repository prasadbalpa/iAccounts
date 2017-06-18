package main

import (
	"fmt"
	"iAccounts/datamodels"
)

func main() {
	var usertable []datamodels.User_table
	usertable = datamodels.GetUsers_Given_ORG_ID("0b929804-8457-44bd-a2c6-a6d70b2a0c68")
	if usertable == nil {
		fmt.Println("Failed to fetch users")
		return
	}
	fmt.Println(usertable[0])
}
