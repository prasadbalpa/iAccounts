package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"iAccounts/utils"
	_ "strconv"
)

type Supplier_table struct {
	Supplier_requesttype int `json:"requesttype"`
	Supplier_tin     string `json:"tin"`
	Supplier_name    string `json:"supplier"`
	Supplier_city    string `json:"city"`
	Supplier_email   string `json:"email"`
	Supplier_id      string
	Supplier_phone   string `json:"phone"`
}

func AddSuppliersByauthCodeAndSupplier(authCode string, supplier Supplier_table) bool {
	/////Verify if the customer and vehicle exist.
	fmt.Println("AddCustomersByauthCodeAndCustomer: ", supplier)
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		fmt.Println("Either orgid or orgname is NIL")
		return false
	}
	if true == IsValidSupplier(*orgid, supplier.Supplier_name) {
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
	fmt.Println(supplier)
	buffer := "insert into supplierbyname (supplier_id, supplier_name, supplier_tin, supplier_city, supplier_email, supplier_phone) values (" + utils.GenerateSecureSessionID() + ",'" + supplier.Supplier_name + "','" + supplier.Supplier_tin + "','" + supplier.Supplier_city + "','" + supplier.Supplier_email + "','" + supplier.Supplier_phone + "')"
	fmt.Println(buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	fmt.Println("Returning true after insertion...")
	return true
}

func GetSuppliersByAuthCode(authCode string) ([]Supplier_table, *string) {
	var supplierlog []Supplier_table
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

	buffer := "select supplier_name, supplier_id, supplier_city, supplier_email, supplier_phone, supplier_tin from supplierbyname"
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
	var supplier_name, supplier_id, supplier_city, supplier_email, supplier_phone, supplier_tin string


	for iteration.Scan(&supplier_name, &supplier_id, &supplier_city, &supplier_email,  &supplier_phone, &supplier_tin) {

		supptab := Supplier_table{Supplier_name: supplier_name, Supplier_tin: supplier_tin, Supplier_city:supplier_city, Supplier_email:supplier_email, Supplier_id: supplier_id, Supplier_phone: supplier_phone}

		supplierlog = append(supplierlog, supptab)
	}
	fmt.Println("Crossed iteration Scan")
	fmt.Println("Customers: ", supplierlog)
	return supplierlog, orgname
}

func IsValidSupplier(orgid string, supplier string) bool {
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
	buffer := "select supplier_id from supplierbyname where supplier_name='" + supplier +"'"
	var supplierid string
	err := session_handle.Query(buffer).Scan(&supplierid)
	if err != nil {
		return false
	}
	return true
}

func GetSuppliers_Given_ORG_ID(orgid string) []Supplier_table {

	var suppliertable []Supplier_table
	var supptab Supplier_table
	var supplier_tin, supplier_name, supplier_city, supplier_email, supplier_phone string
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
	buffer := "select supplier_tin, supplier_name, supplier_city, supplier_email, supplier_phone from suppliersbyname"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")

	for iteration.Scan(&supplier_tin, &supplier_name, &supplier_city, &supplier_email, &supplier_phone) {
		fmt.Println("Assigning the values to the custtable..")
		supptab = Supplier_table{Supplier_tin: supplier_tin, Supplier_name: supplier_name, Supplier_city: supplier_city, Supplier_email: supplier_email, Supplier_phone: supplier_phone}
		fmt.Println("Appending to the array...")
		suppliertable = append(suppliertable, supptab)
	}
	fmt.Println("Crossed iteration Scan")
	return suppliertable
}
