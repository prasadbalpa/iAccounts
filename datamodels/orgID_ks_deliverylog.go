package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
)

type Delivery_log struct {
	Delivery_timestamp string `json:"timestamp"`
	Delivery_id        string
	Delivery_customer  string `json:"customer"`
	Delivery_quantity  int    `json:"quantity"`
	Delivery_vehicle   string `json:"vehicle"`
}

func GetAllDeliveryLogsByorgID(orgID string) []Delivery_log {
	var deliverylog []Delivery_log

	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(orgID))
	if cluster_handle == nil {
		return nil
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil
	}
	fmt.Println("acquired session handle")

	buffer := "select * from deliverylogbytimestamp"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var delivery_timestamp, delivery_customer, delivery_vehicle, delivery_id string
	var delivery_quantity int

	for iteration.Scan(&delivery_timestamp, &delivery_id, &delivery_customer, &delivery_quantity, &delivery_vehicle) {
		fmt.Println("Assigning the values to the custtable..")
		deltab := Delivery_log{Delivery_id: delivery_id, Delivery_customer: delivery_customer, Delivery_quantity: delivery_quantity, Delivery_vehicle: delivery_vehicle}
		fmt.Println("Appending to the array...")
		deliverylog = append(deliverylog, deltab)
	}
	fmt.Println("Crossed iteration Scan")
	return deliverylog
}
