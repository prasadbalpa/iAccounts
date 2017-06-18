package main

import (
	"fmt"
	"iAccounts/datamodels"
)

func main() {
	p := datamodels.GetORG_Given_Name_Tin("Kamath Traders", "29260092504")
	if p == nil {
		fmt.Println("Test:GetORG_Given_Name_Tin() Failed")
	} else {
		fmt.Println(p.Org_email)
		fmt.Println("Test:GetORG_Given_Name_Tin() Passed")
	}

	q := datamodels.GetSessionID_Given_Client_ID("36128282-8998-4b42-a468-9bf00d3c1367")
	if q == nil {
		fmt.Println("Test:GetSessionID_Given_Client_ID() Failed")
	} else {
		fmt.Println(*q)
		fmt.Println("Test:GetSessionID_Given_Client_ID() Passed")
	}
	r := datamodels.InsertSession("9902016406", "334222", "10.10.10.10")

	if r == nil {
		fmt.Println("Test:InsertSession() Failed")
	} else {
		fmt.Println(r)
		fmt.Println("Test:InsertSession() Passed")
	}

	s := datamodels.DeleteSession_Given_Phone("9902016406")

	if s == false {
		fmt.Println("Test:DeleteSession_Given_Phone() Failed")
	} else {
		fmt.Println("Test:DeleteSession_Given_Phone() Passed")
	}

}
