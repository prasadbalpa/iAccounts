package handlers

import (
	"iAccounts/datamodels"
	"fmt"
	"encoding/json"
)

type Suppliers_response struct {
	Supplier_name string `json:"supplier"`
	Supplier_tin string `json:"tin"`
	Supplier_id string `json:"id"`
	Supplier_email string `json:"email"`
	Supplier_phone string `json:"phone"`
	Supplier_city string `json:"city,omitempty"`
}

type suppresponse struct {
	Response_type int                `json:"responsetype"`
	Organization  string             `json:"organization"`
	Supp           []Suppliers_response `json:"suppliers"`
}

func Add_Suppliers_Given_AuthorizationCode(authorization string, supplier datamodels.Supplier_table) bool {
	fmt.Println("Add_Customers_Given_AuthorizationCode: ")
	fmt.Println(supplier)
	err := datamodels.AddSuppliersByauthCodeAndSupplier(authorization, supplier)

	return err
}

func Fetch_Suppliers_Given_AuthorizationCode(authorization string) []byte {
	var supplier_response []Suppliers_response
	var supplier_log []datamodels.Supplier_table
	var orgname *string
	supplier_log, orgname = datamodels.GetSuppliersByAuthCode(authorization)

	if supplier_log == nil {
		return nil
	}
	supplier_response = make([]Suppliers_response, len(supplier_log))
	for i:=0;i<len(supplier_log);i++ {
		supplier_response[i].Supplier_name = supplier_log[i].Supplier_name
		supplier_response[i].Supplier_email = supplier_log[i].Supplier_email
		supplier_response[i].Supplier_id = supplier_log[i].Supplier_id
		supplier_response[i].Supplier_tin = supplier_log[i].Supplier_tin
		supplier_response[i].Supplier_phone = supplier_log[i].Supplier_phone
	}
	var ccr suppresponse
	ccr.Supp = make([]Suppliers_response, len(supplier_log))
	for i := 0; i < len(supplier_log); i++ {
		ccr.Supp[i] = supplier_response[i]
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
