package datamodels

import (
	_ "math/big"
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"iAccounts/utils"
	"strconv"
)

type Purchase_Table struct {
	Purchase_requesttype int `json:"requesttype,omitempty"`
	Purchase_timestamp string `json:"timestamp,omitempty"`
	Purchase_id        string
	Purchase_supplier  string `json:"supplier,omitempty"`
	Purchase_quantity  int    `json:"quantity,omitempty"`
	Purchase_vehicle   string `json:"vehicle,omitempty"`
	Purchase_price string `json:"price,omitempty"`
	Purchase_orderid string `json:"orderid"`
	Purchase_product string `json:"product"`
	Purchase_org string `json:"Organization,omitempty"`
}

func GetAllPurchasesByauthCode(authCode string) ([]Purchase_Table, *string) {
	var purchaselog []Purchase_Table
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

	buffer := "select * from purchasesbytimestamp"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	//fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil, nil
	}
	//fmt.Println(iteration)
	fmt.Println("executed::Iter()")
	var purchase_timestamp, purchase_supplier, purchase_orderid, purchase_id, purchase_vehicle, purchase_product string
	var purchase_quantity, purchase_price int

	for iteration.Scan(&purchase_timestamp, &purchase_supplier, &purchase_vehicle, &purchase_orderid, &purchase_vehicle, &purchase_product,&purchase_quantity, &purchase_price ) {

		purctab := Purchase_Table{Purchase_org: *orgname, Purchase_id: purchase_id, Purchase_supplier: purchase_supplier, Purchase_quantity: purchase_quantity, Purchase_vehicle: purchase_vehicle, Purchase_timestamp: purchase_timestamp, Purchase_orderid: purchase_orderid}

		purchaselog = append(purchaselog, purctab)
	}
	fmt.Println("Crossed iteration Scan")
	return purchaselog, orgname
}

func AddPurchaseByauthCode(authCode string, purc Purchase_Table) bool {
	/////Verify if the customer and vehicle exist.
	fmt.Println("AddDeliByauthCode: ", purc)
	orgid, orgname := GetORG_Given_authCode(authCode)
	if orgid == nil || orgname == nil {
		fmt.Println("Either orgid or orgname is NIL")
		return false
	}
	if false == IsValidDeliveryVehicle(*orgid, purc.Purchase_vehicle) {
		fmt.Println("Not a valid delivery vehicle")
		return false
	}
	if false == IsValidCustomer(*orgid, purc.Purchase_supplier) {
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
	fmt.Println(purc)
	buffer := "insert into purchasesbytimestamp (purchase_timestamp, purchase_id, purchase_supplier, purchase_orderid, purchase_price, purchase_product, purchase_quantity, purchase_vehicle) values (" + utils.GenerateSecureSessionID() + ",'" + deliverylog.Delivery_timestamp + "','" + deliverylog.Delivery_customer + "'," + strconv.Itoa(deliverylog.Delivery_quantity) + ",'" + deliverylog.Delivery_vehicle + "')"
	fmt.Println(buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	fmt.Println("Returning true after insertion...")
	return true
}