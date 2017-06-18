package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
)

type Customer_table struct {
	Customer_tin     string
	Customer_name    string
	Customer_city    string
	Customer_created string //timestamp
	Customer_email   string
	Customer_id      string
	Customer_phone   string
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
