package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"iAccounts/utils"

)

type Vehicle_table struct {
	Vehicle_requesttype int `json:"requesttype"`
	Vehicle_id     string
	Vehicle_number string `json:"number"`
}

func AddProductsByauthCodeAndVehicle(authCode string, vehicle Vehicle_table) bool {
	/////Verify if the customer and vehicle exist.
	fmt.Println("AddCustomersByauthCodeAndCustomer: ", vehicle)
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		fmt.Println("Either orgid or orgname is NIL")
		return false
	}
	if true == IsValidDeliveryVehicle(*orgid, vehicle.Vehicle_number) {
		fmt.Println("Not a valid vehicle")
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
	fmt.Println(vehicle)
	buffer := "insert into deliveryvehicles (vehicle_id, vehicle_number) values (" + utils.GenerateSecureSessionID() + ",'" + vehicle.Vehicle_number + "')"
	fmt.Println(buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	fmt.Println("Returning true after insertion...")
	return true
}

func GetVehiclesByAuthCode(authCode string) ([]Vehicle_table, *string) {
	var vehiclelog []Vehicle_table
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

	buffer := "select * from deliveryvehicles"

	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	//fmt.Println(iteration.NumRows())
	if iteration == nil {
		fmt.Println("Did not get any vehicles..")
		return nil, nil
	}
	//fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var vehicle_id, vehicle_number string


	for iteration.Scan(&vehicle_id, &vehicle_number) {

		vehtab := Vehicle_table{Vehicle_id: vehicle_id, Vehicle_number: vehicle_number}

		vehiclelog = append(vehiclelog, vehtab)
	}
	fmt.Println("Crossed iteration Scan")
	fmt.Println("Customers: ", vehiclelog)
	return vehiclelog, orgname
}


func IsValidDeliveryVehicle(orgid string, vehicle string) bool {
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(orgid)) //change it to config package value
	fmt.Println("OrgID KS : " + cfg.GetOrgIDKeySpace(orgid))
	if cluster_handle == nil {
		return false
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return false
	}
	fmt.Println("acquired session handle")
	//Not using '*' in the query to avoid any out of sequence results
	buffer := "select vehicle_id from deliveryvehicles where vehicle_number='" + vehicle +"' allow filtering"
	var vehicleid string
	err := session_handle.Query(buffer).Scan(&vehicleid)
	if err != nil {
		return false
	}
	return true
}

func GetVehicles_Given_ORG_ID(orgid string) []Vehicle_table {

	var vehicletable []Vehicle_table
	var vehtab Vehicle_table
	var vehicle_number string
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgIDKeySpace(orgid)) //change it to config package value
	fmt.Println("OrgID KS : " + cfg.GetOrgIDKeySpace(orgid))
	if cluster_handle == nil {
		return nil
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil
	}
	fmt.Println("acquired session handle")
	//Not using '*' in the query to avoid any out of sequence results
	buffer := "select vehicle_number from deliveryvehicles"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")

	for iteration.Scan(&vehicle_number) {
		fmt.Println("Assigning the values to the custtable..")
		vehtab = Vehicle_table{Vehicle_number: vehicle_number}
		fmt.Println("Appending to the array...")
		vehicletable = append(vehicletable, vehtab)
	}
	fmt.Println("Crossed iteration Scan")
	return vehicletable
}
