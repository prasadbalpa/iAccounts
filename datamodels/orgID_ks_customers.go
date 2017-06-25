package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"iAccounts/utils"
	_ "strconv"
)

type Customer_table struct {
	Customer_requesttype int `json:"requesttype"`
	Customer_tin     string `json:"tin"`
	Customer_name    string `json:"customer"`
	Customer_city    string `json:"city"`
	Customer_created string
	Customer_email   string `json:"email"`
	Customer_id      string
	Customer_phone   string `json:"phone"`
}

func AddCustomersByauthCodeAndCustomer(authCode string, customer Customer_table) bool {
	/////Verify if the customer and vehicle exist.
	fmt.Println("AddCustomersByauthCodeAndCustomer: ", customer)
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		fmt.Println("Either orgid or orgname is NIL")
		return false
	}
	if true == IsValidCustomer(*orgid, customer.Customer_name) {
		fmt.Println("Not a valid customer")
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
	fmt.Println(customer)
	buffer := "insert into customerbyname (customer_id, customer_name, customer_tin, customer_city, customer_email, customer_phone, customer_created) values (" + utils.GenerateSecureSessionID() + ",'" + customer.Customer_name + "','" + customer.Customer_tin + "','" + customer.Customer_city + "','" + customer.Customer_email + "','" + customer.Customer_phone + "','" + utils.GenerateUnixTimeStamp() + "')"
	fmt.Println(buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into customerbyphone (customer_id, customer_name, customer_tin, customer_city, customer_email, customer_phone, customer_created) values (" + utils.GenerateSecureSessionID() + ",'" + customer.Customer_name + "','" + customer.Customer_tin + "','" + customer.Customer_city + "','" + customer.Customer_email + "','" + customer.Customer_phone + "','" + utils.GenerateUnixTimeStamp() + "')"
	//buffer = "insert into customerbytin (delivery_id, delivery_timestamp, delivery_customer, delivery_quantity, delivery_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
	fmt.Println(buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into customerbytin (customer_id, customer_name, customer_tin, customer_city, customer_email, customer_phone, customer_created) values (" + utils.GenerateSecureSessionID() + ",'" + customer.Customer_name + "','" + customer.Customer_tin + "','" + customer.Customer_city + "','" + customer.Customer_email + "','" + customer.Customer_phone + "','" + utils.GenerateUnixTimeStamp() + "')"
	//buffer = "insert into customerbyphone (delivery_id, delivery_timestamp, delivery_customer, delivery_quantity, delivery_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
	fmt.Println(buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	fmt.Println("Returning true after insertion...")
	return true
}

func GetCustomersByAuthCode(authCode string) ([]Customer_table, *string) {
	var customerlog []Customer_table
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

	buffer := "select * from customerbyname"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	//fmt.Println(iteration.NumRows())
	if iteration == nil {
		fmt.Println("Did not get any customers..")
		return nil, nil
	}
	//fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var customer_name, customer_tin, customer_city, customer_created, customer_email, customer_id, customer_phone string


	for iteration.Scan(&customer_name, &customer_tin, &customer_city, &customer_created, &customer_email, &customer_id, &customer_phone) {

		custab := Customer_table{Customer_name: customer_name, Customer_tin: customer_tin, Customer_city:customer_city, Customer_created:customer_created, Customer_email:customer_email, Customer_id: customer_id, Customer_phone: customer_phone}

		customerlog = append(customerlog, custab)
	}
	fmt.Println("Crossed iteration Scan")
	fmt.Println("Customers: ", customerlog)
	return customerlog, orgname
}

func IsValidCustomer(orgid string, customer string) bool {
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
	buffer := "select customer_id from customerbyname where customer_name='" + customer +"'"
	var customerid string
	err := session_handle.Query(buffer).Scan(&customerid)
	if err != nil {
		return false
	}
	return true
}

func GetCustomers_Given_ORG_ID(orgid string) []Customer_table {

	var customertable []Customer_table
	var custtab Customer_table
	var customer_tin, customer_name, customer_city, customer_email, customer_phone string
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
	buffer := "select customer_tin, customer_name, customer_city, customer_email, customer_phone from customers"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")

	for iteration.Scan(&customer_tin, &customer_name, &customer_city, &customer_email, &customer_phone) {
		fmt.Println("Assigning the values to the custtable..")
		custtab = Customer_table{Customer_tin: customer_tin, Customer_name: customer_name, Customer_city: customer_city, Customer_email: customer_email, Customer_phone: customer_phone}
		fmt.Println("Appending to the array...")
		customertable = append(customertable, custtab)
	}
	fmt.Println("Crossed iteration Scan")
	return customertable
}
