package handlers

import (
	"iAccounts/datamodels"
	"fmt"
	"encoding/json"
)

type Customers_response struct {
	Customer_name string `json:"customer"`
	Customer_tin string `json:"tin"`
	Customer_id string `json:"id""`
	Customer_email string `json:"email"`
	Customer_phone string `json:"phone"`
}

type custresponse struct {
	Response_type int                `json:"responsetype"`
	Organization  string             `json:"organization"`
	Cust           []Customers_response `json:"customers"`
}

func Fetch_Customers_Given_AuthorizationCode(authorization string) []byte {
	var customer_response []Customers_response
	var customer_log []datamodels.Customer_table
	var orgname *string
	customer_log, orgname = datamodels.GetCustomersByAuthCode(authorization)

	if customer_log == nil {
		return nil
	}
	customer_response = make([]Customers_response, len(customer_log))
	for i:=0;i<len(customer_log);i++ {
		customer_response[i].Customer_name = customer_log[i].Customer_name
		customer_response[i].Customer_email = customer_log[i].Customer_email
		customer_response[i].Customer_id = customer_log[i].Customer_id
		customer_response[i].Customer_tin = customer_log[i].Customer_tin
		customer_response[i].Customer_phone = customer_log[i].Customer_phone
	}
    var ccr custresponse
    ccr.Cust = make([]Customers_response, len(customer_log))
	for i := 0; i < len(customer_log); i++ {
		ccr.Cust[i] = customer_response[i]
	}
	ccr.Response_type = 10
	ccr.Organization = *orgname
	fmt.Println("Marshalling to JSON...")

	custresp, errs := json.Marshal(ccr)
	fmt.Println(string(custresp))
	if errs != nil {
		//send error json for login
		fmt.Println("Failed to Marshal..")
		return nil
	}
	return custresp
}
