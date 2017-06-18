package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
)

type User_table struct {
	User_phone    string
	User_id       string
	User_passwd   string
	User_role     string
	User_username string
}

func GetUsers_Given_ORG_ID(orgid string) []User_table {

	var usertable []User_table
	var usertab User_table
	var user_phone string
	var user_role string
	var user_username string

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
	buffer := "select user_userphone, user_role, user_username from users"
	//log the buffer
	fmt.Println("executing::" + buffer)
	iteration := session_handle.Query(buffer).Iter()
	fmt.Println(iteration.NumRows())
	if iteration == nil {
		return nil
	}
	fmt.Println(iteration)
	fmt.Println("executed::Iter()")

	for iteration.Scan(&user_phone, &user_role, &user_username) {
		fmt.Println("Assigning the values to the custtable..")
		usertab = User_table{User_phone: user_phone, User_role: user_role, User_username: user_username}
		fmt.Println("Appending to the array...")
		usertable = append(usertable, usertab)
	}
	fmt.Println("Crossed iteration Scan")
	return usertable
}
