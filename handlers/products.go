package handlers

import (
	"iAccounts/datamodels"
	"fmt"
	"encoding/json"
	"math/big"
)

type Products_response struct {
	Product_id string `json:"id"`
	Product_discount big.Float `json:"discount"`
	Product_name string `json:"name"`
	Product_price big.Float `json:"price"`
}

type prodresponse struct {
	Response_type int                `json:"responsetype"`
	Organization  string             `json:"organization"`
	Prod           []Products_response `json:"products"`
}

func Add_Products_Given_AuthorizationCode(authCode string, product datamodels.Product_table) bool {

	fmt.Println("Add_Products_Given_AuthorizationCode: ")
	fmt.Println(product)
	err := datamodels.AddProductsByauthCodeAndProduct(authCode, product)

	return err
}

func Fetch_Products_Given_AuthorizationCode(authorization string) []byte {
	var product_response []Products_response
	var product_log []datamodels.Product_table
	var orgname *string
	product_log, orgname = datamodels.GetProductsByAuthCode(authorization)

	if product_log == nil {
		return nil
	}
	product_response = make([]Products_response, len(product_log))
	for i:=0;i<len(product_log);i++ {
		product_response[i].Product_name = product_log[i].Product_name
		product_response[i].Product_discount = product_log[i].Product_discount
		product_response[i].Product_id = product_log[i].Product_id
		product_response[i].Product_price = product_log[i].Product_price

	}
	var pdr prodresponse
	pdr.Prod = make([]Products_response, len(product_log))
	for i := 0; i < len(product_log); i++ {
		pdr.Prod[i] = product_response[i]
	}
	pdr.Response_type = 10
	pdr.Organization = *orgname
	fmt.Println("Marshalling to JSON...")

	prodresp, errs := json.Marshal(pdr)
	fmt.Println(string(prodresp))
	if errs != nil {
		//send error json for login
		fmt.Println("Failed to Marshal..")
		return nil
	}
	return prodresp
}

