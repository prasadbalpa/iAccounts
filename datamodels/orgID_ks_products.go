package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"math/big"
)

type Product_table struct {
	Product_id       string
	Product_discount big.Float
	Product_name     string
	Product_price    big.Float
}

func GetProductsByAuthCode(authCode string ) ([]Product_table, *string) {
	var productlog []Product_table
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

	buffer := "select product_id, product_name from products"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		fmt.Println("Did not get any customers..")
		return nil, nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var product_name, product_id string
	//var product_discount, product_price big.Float
    //&product_discount,&product_discount,
	//Product_discount:product_discount,
	//, Product_price:product_price
	for iteration.Scan(&product_id,  &product_name) {

		prodtab := Product_table{Product_id:product_id,  Product_name:product_name}

		productlog = append(productlog, prodtab)
	}
	fmt.Println("Crossed iteration Scan")
	fmt.Println("Products: ", productlog)
	return productlog, orgname
}


func GetProducts_Given_ORG_ID(orgid string) []Product_table {

	var producttable []Product_table
	var prodtab Product_table
	var product_name string

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
	buffer := "select product_name from products"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")

	for iteration.Scan(&product_name) {
		fmt.Println("Assigning the values to the custtable..")
		prodtab = Product_table{Product_name: product_name}
		fmt.Println("Appending to the array...")
		producttable = append(producttable, prodtab)
	}
	fmt.Println("Crossed iteration Scan")
	return producttable
}
