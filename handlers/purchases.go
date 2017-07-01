package handlers

import (
	"math/big"
	"iAccounts/datamodels"
	"encoding/json"
	"fmt"
)

type Purchases_response struct {
	Purchase_id string `json:"id"`
	Purchase_supplier string `json:"supplier"`
	Purchase_product string `json:"product"`
	Purchase_quantity big.Float `json:"quantity"`
	Purchase_price big.Float `json:"price"`
	Purchase_vehicle string `json:"vehicle"`
	Purchase_timestamp string `json:"timestamp"`
	Purchase_orderid string `json:"orderid"`
}

type purcresponse struct {
	Response_type int                `json:"responsetype"`
	Organization  string             `json:"organization"`
	Prod           []Purchases_response `json:"purchases"`
}

func Add_Purchases_Given_AuthorizationCode(authCode string, purc datamodels.Purchase_Table) bool {
	err := datamodels.AddPurchaseByauthCode(authCode, purc)
	return err
}

func Fetch_Purchases_Given_AuthorizationCode(authCode string) []byte {
	var purchaseresponse []Purchases_response
	var purchase_log []datamodels.Purchase_Table
	var orgname *string


	purchase_log, orgname = datamodels.GetAllPurchasesByauthCode(authCode)

	if purchase_log == nil {
		var ddno Dlrresponse
		ddno.Response_type = 10
		ddno.Organization = "Blah Blah Blah"
		deliveryre, errs := json.Marshal(ddno)
		fmt.Println(string(deliveryre))
		if errs != nil {
			//send error json for login
			fmt.Println("Failed to Marshal..")
			return nil
		}
		return deliveryre
	} else {
		//fmt.Println(delivery_log)
		purchaseresponse = make([]Purchases_response, len(purchase_log))
		for i := 0; i < len(purchase_log); i++ {
			purchaseresponse[i].Purchase_id = purchase_log[i].Purchase_id
			purchaseresponse[i].Purchase_supplier = purchase_log[i].Purchase_supplier
			purchaseresponse[i].Purchase_quantity = purchase_log[i].Purchase_quantity
			purchaseresponse[i].Purchase_timestamp = purchase_log[i].Purchase_timestamp
			purchaseresponse[i].Purchase_vehicle = purchase_log[i].Purchase_vehicle
			purchaseresponse[i].Purchase_product = purchase_log[i].Purchase_product
			purchaseresponse[i].Purchase_price = purchase_log[i].Purchase_price
			purchaseresponse[i].Purchase_orderid = purchase_log[i].Purchase_orderid


		}
		var ddr Dlrresponse
		ddr.Dlr = make([]Deliveryresponse, len(purchase_log))
		fmt.Println("Copying the pointers")
		for i := 0; i < len(purchase_log); i++ {

			ddr.Dlr[i] = purchaseresponse[i]
		}
		//fmt.Println(deliveryresponse)
		ddr.Response_type = 10
		ddr.Organization = *orgname
		fmt.Println("Marshalling to JSON...")
		//fmt.Println(ddr)
		deliveryresp, errs := json.Marshal(ddr)
		fmt.Println(string(deliveryresp))
		if errs != nil {
			//send error json for login
			fmt.Println("Failed to Marshal..")
			return nil
		}
		return deliveryresp
	}
}
