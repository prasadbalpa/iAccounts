package handlers
//All of the deliverylog handlers
import (
	"encoding/json"
	"fmt"
	"iAccounts/datamodels"
)

type Dlrresponse struct {
	Response_type int                `json:"responsetype"`
	Organization  string             `json:"organization"`
	Dlr           []Deliveryresponse `json:"deliverylog"`
}
type Deliveryresponse struct {
	Delivery_id string `json:"id"`
	Delivery_customer  string `json:"customer"`
	Delivery_quantity  int    `json:"quantity"`
	Delivery_vehicle   string `json:"vehicle"`
	Delivery_timestamp string `json:"timestamp"`
}

func Delete_Delivery_Logs_Given_authorizationCode(authCode string, ids []string) bool {
	if false == datamodels.Delete_Delivery_Logs(authCode, ids) {
		return false
	}
	return true
}

func Add_Delivery_Logs_Given_authorizationCode(authCode string, deliverylog datamodels.Delivery_log) bool {

	fmt.Println("Add_Delivery_Logs_Given_authorizationCode: deliverylog:- ", deliverylog)

	err := datamodels.AddDeliveryLogsByauthCode(authCode, deliverylog)

	return err
}

func Search_Delivery_Logs_Given_authorizationCode(authCode string, customer string) []byte {
	var deliveryresponse []Deliveryresponse
	var delivery_log []datamodels.Delivery_log
	var orgname *string


	delivery_log, orgname = datamodels.GetAllDeliveryLogsByauthCodeAndCustomer(authCode, customer)


		//fmt.Println(delivery_log)
		deliveryresponse = make([]Deliveryresponse, len(delivery_log))
		for i := 0; i < len(delivery_log); i++ {
			deliveryresponse[i].Delivery_id = delivery_log[i].Delivery_id
			deliveryresponse[i].Delivery_customer = delivery_log[i].Delivery_customer
			deliveryresponse[i].Delivery_quantity = delivery_log[i].Delivery_quantity
			deliveryresponse[i].Delivery_timestamp = delivery_log[i].Delivery_timestamp
			deliveryresponse[i].Delivery_vehicle = delivery_log[i].Delivery_vehicle
		}
		var ddr Dlrresponse
		ddr.Dlr = make([]Deliveryresponse, len(delivery_log))
		fmt.Println("Copying the pointers")
		for i := 0; i < len(delivery_log); i++ {

			ddr.Dlr[i] = deliveryresponse[i]
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

func Fetch_Delivery_Logs_Given_authorizationCode(authCode string) []byte {
	var deliveryresponse []Deliveryresponse
	var delivery_log []datamodels.Delivery_log
    var orgname *string


   	delivery_log, orgname = datamodels.GetAllDeliveryLogsByauthCode(authCode)

	if delivery_log == nil {
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
		deliveryresponse = make([]Deliveryresponse, len(delivery_log))
		for i := 0; i < len(delivery_log); i++ {
			deliveryresponse[i].Delivery_id = delivery_log[i].Delivery_id
			deliveryresponse[i].Delivery_customer = delivery_log[i].Delivery_customer
			deliveryresponse[i].Delivery_quantity = delivery_log[i].Delivery_quantity
			deliveryresponse[i].Delivery_timestamp = delivery_log[i].Delivery_timestamp
			deliveryresponse[i].Delivery_vehicle = delivery_log[i].Delivery_vehicle
		}
		var ddr Dlrresponse
		ddr.Dlr = make([]Deliveryresponse, len(delivery_log))
		fmt.Println("Copying the pointers")
		for i := 0; i < len(delivery_log); i++ {

			ddr.Dlr[i] = deliveryresponse[i]
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
func Fetch_Delivery_Logs_Given_authorizationCode_Limit(authCode string, limit int) []byte {
	var deliveryresponse []Deliveryresponse
	var delivery_log []datamodels.Delivery_log
	var orgname *string


	delivery_log, orgname = datamodels.GetAllDeliveryLogsByauthCode_Limit(authCode, limit)

	if delivery_log == nil {
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
		deliveryresponse = make([]Deliveryresponse, len(delivery_log))
		for i := 0; i < len(delivery_log); i++ {
			deliveryresponse[i].Delivery_customer = delivery_log[i].Delivery_customer
			deliveryresponse[i].Delivery_quantity = delivery_log[i].Delivery_quantity
			deliveryresponse[i].Delivery_timestamp = delivery_log[i].Delivery_timestamp
			deliveryresponse[i].Delivery_vehicle = delivery_log[i].Delivery_vehicle
		}
		var ddr Dlrresponse
		ddr.Dlr = make([]Deliveryresponse, len(delivery_log))
		fmt.Println("Copying the pointers")
		for i := 0; i < len(delivery_log); i++ {

			ddr.Dlr[i] = deliveryresponse[i]
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