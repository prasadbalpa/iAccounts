package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"strconv"
	"iAccounts/utils"
)

type Delivery_log struct {
	Delivery_requesttype int `json:"requesttype,omitempty"`
	Delivery_timestamp string `json:"timestamp,omitempty"`
	Delivery_id        string
	Delivery_customer  string `json:"customer,omitempty"`
	Delivery_quantity  int    `json:"quantity,omitempty"`
	Delivery_vehicle   string `json:"vehicle,omitempty"`
	Delivery_org string `json:"Organization,omitempty"`
	Delivery_prod string `json:"product,omitempty"`
}
type Deliverylog_delete struct {
	Response_type int `json:"responsetype"`
	todelete int `json:"todelete"`
	deleteids []string `json:"delete"`
}

func Delete_Delivery_Logs(authCode string, ids []string) bool {
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		fmt.Println("Either orgid or orgname is NIL")
		return false
	}
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(*orgid))
	if cluster_handle == nil {
		return false
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return false
	}
	fmt.Println("acquired session handle")
	var delivery_timestamp, delivery_id, delivery_customer, delivery_vehicle string
	var delivery_quantity int
	fmt.Println("Length of Array: ", len(ids))
	for i:=0;i<len(ids);i++ {
		buffer := "select delivery_customer, delivery_quantity, delivery_timestamp, delivery_vehicle from deliverylogbyid where delivery_id=" + ids[i]
		fmt.Println(buffer)
		delivery_id = ids[i]
		errs := session_handle.Query(buffer).Scan(&delivery_customer, &delivery_quantity, &delivery_timestamp, &delivery_vehicle)
		if errs != nil {
			fmt.Println("Error while deleting....")
			return false
		}
		buffer = "delete from deliverylogbytimestamp where delivery_timestamp='" + delivery_timestamp + "' and delivery_id=" + delivery_id + " and delivery_customer='" + delivery_customer + "'"
		fmt.Println(buffer)
		errs = session_handle.Query(buffer).Exec()
		if errs != nil {
			fmt.Println("Error while deleting....")
			return false
		}
		buffer = "delete from deliverylogbyid where delivery_id=" + delivery_id
		fmt.Println(buffer)
		errs = session_handle.Query(buffer).Exec()
		if errs != nil {
			fmt.Println("Error while deleting....")
			return false
		}
		buffer = "delete from deliverylogbytimestamp where delivery_timestamp='" + delivery_timestamp + "' and delivery_id=" + delivery_id + " and delivery_customer='" + delivery_customer + "'"
		fmt.Println(buffer)
		errs = session_handle.Query(buffer).Exec()
		if errs != nil {
			fmt.Println("Error while deleting....")
			return false
		}

	}

    return true
}

func AddDeliveryLogsByauthCode(authCode string, deliverylog Delivery_log) bool {

	/////Verify if the customer and vehicle exist.
	fmt.Println("AddDeliveryLogsByauthCode: ", deliverylog)
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		fmt.Println("Either orgid or orgname is NIL")
		return false
	}
	if false == IsValidDeliveryVehicle(*orgid, deliverylog.Delivery_vehicle) {
		fmt.Println("Not a valid delivery vehicle")
		return false
	}
	if false == IsValidCustomer(*orgid, deliverylog.Delivery_customer) {
		fmt.Println("Not a valid delivery vehicle")
		return false
	}
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(*orgid))
	if cluster_handle == nil {
		return false
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return false
	}
	fmt.Println("acquired session handle")
	fmt.Println(deliverylog)
	buffer := "insert into deliverylogbytimestamp (delivery_id, delivery_timestamp, delivery_customer, delivery_quantity, delivery_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
    fmt.Println(buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into deliverylogbycustomer (delivery_id, delivery_timestamp, delivery_customer, delivery_quantity, delivery_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
	fmt.Println(buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into deliverylogbytimestampcustomer (delivery_id, delivery_timestamp, delivery_customer, delivery_quantity, delivery_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
	fmt.Println(buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into deliverylogbyid (delivery_id, delivery_timestamp, delivery_customer, delivery_quantity, delivery_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
	fmt.Println(buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {

		return false
	}
	fmt.Println("Returning true after insertion...")
	return true
}

func GetAllDeliveryLogsByauthCode(authCode string) ([]Delivery_log, *string) {
	var deliverylog []Delivery_log
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		return nil, nil
	}
	//***********************
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(*orgid))
	if cluster_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired session handle")

	buffer := "select * from deliverylogbytimestamp"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	//fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil, nil
	}
	//fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var delivery_timestamp, delivery_customer, delivery_vehicle, delivery_id string
	var delivery_quantity int

	for iteration.Scan(&delivery_timestamp, &delivery_id, &delivery_customer, &delivery_quantity, &delivery_vehicle) {

		deltab := Delivery_log{Delivery_org: *orgname, Delivery_id: delivery_id, Delivery_customer: delivery_customer, Delivery_quantity: delivery_quantity, Delivery_vehicle: delivery_vehicle, Delivery_timestamp: delivery_timestamp}

		deliverylog = append(deliverylog, deltab)
	}
	fmt.Println("Crossed iteration Scan")
	return deliverylog, orgname
}
func GetAllDeliveryLogsByauthCode_Limit(authCode string, limit int) ([]Delivery_log, *string) {
	var deliverylog []Delivery_log
	orgid, orgname := GetORG_Given_authCode(authCode)
    if orgid == nil || orgname == nil {
		return nil, nil
	}
	//***********************
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(*orgid))
	if cluster_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired session handle")

	buffer := "select * from deliverylogbytimestamp  limit " + strconv.Itoa(limit)
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	//fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil, nil
	}
	//fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var delivery_timestamp, delivery_customer, delivery_vehicle, delivery_id string
	var delivery_quantity int

	for iteration.Scan(&delivery_timestamp, &delivery_id, &delivery_customer, &delivery_quantity, &delivery_vehicle) {

		deltab := Delivery_log{Delivery_org: *orgname, Delivery_id: delivery_id, Delivery_customer: delivery_customer, Delivery_quantity: delivery_quantity, Delivery_vehicle: delivery_vehicle}

		deliverylog = append(deliverylog, deltab)
	}
	fmt.Println("Crossed iteration Scan")
	return deliverylog, orgname
}
