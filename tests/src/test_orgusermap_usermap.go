package main

import (
	"fmt"
	"iAccounts/datamodels"
)

func main() {
	p := datamodels.GetUserMap_Given_User_Phone("9902016406")
	if p == nil {
		fmt.Println("Test:GetUserMap_Given_User_Phone() : Failed")
	} else {
		fmt.Println(p.Usermap_orgname)
		fmt.Println("Test:GetUserMap_Given_User_Phone() : Passed")
	}
	q := datamodels.GetORG_Given_User_Phone("9902016400")
	if q == nil {
		fmt.Println("Test:GetORG_Given_User_Phone() : Failed")
	} else {
		fmt.Println(*q)
		fmt.Println("Test:GetORG_Given_User_Phone() : Passed")
	}
}
