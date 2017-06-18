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
	Delivery_customer  string `json:"customer"`
	Delivery_quantity  int    `json:"quantity"`
	Delivery_vehicle   string `json:"vehicle"`
	Delivery_timestamp string `json:"timestamp"`
}

func Fetch_Delivery_Logs_Given_orgID(orgID string) []byte {
	var deliveryresponse []Deliveryresponse
	var delivery_log []datamodels.Delivery_log

	delivery_log = datamodels.GetAllDeliveryLogsByorgID(orgID)

	if delivery_log == nil {
		return nil
	}
	fmt.Println(delivery_log)
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
		fmt.Println("Assigning the pointer....")
		ddr.Dlr[i] = deliveryresponse[i]
	}
	fmt.Println(deliveryresponse)
	ddr.Response_type = 10
	ddr.Organization = "Kamath Traders"
	fmt.Println("Marshalling to JSON...")
	fmt.Println(ddr)
	deliveryresp, errs := json.Marshal(ddr)
	fmt.Println(string(deliveryresp))
	if errs != nil {
		//send error json for login
		fmt.Println("Failed to Marshal..")
		return nil
	}
	return deliveryresp
}
