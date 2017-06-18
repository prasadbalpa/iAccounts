package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
)

type OrgUserMap_table struct {
	Usermap_orgname   string
	Usermap_username  string
	Usermap_userphone string
	Usermap_id        string
	Usermap_tin       string
}

func GetORG_Given_User_Phone(user_phone string) (*string, *string) {

	var orgusermaptable_orgname, orgusermaptable_orgtin *string

	orgusermaptable_orgname = new(string)
	orgusermaptable_orgtin = new(string)
	if orgusermaptable_orgname == nil {
		return nil, nil
	}

	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgUserMapKeySpace()) //change it to config package value
	if cluster_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired session handle")
	//Not using '*' in the query to avoid any out of sequence results
	buffer := "select usermap_orgname, usermap_orgtin from usermapbyphone where usermap_userphone ='" + user_phone + "' ALLOW FILTERING"
	//log the buffer
	fmt.Println("executing::" + buffer)
	err := session_handle.Query(buffer).Scan(orgusermaptable_orgname, orgusermaptable_orgtin)
	if err != nil {
		return nil, nil
	}
	//return the pointer to the result
	return orgusermaptable_orgname, orgusermaptable_orgtin
}

//Only one entry is expected as these are the PRIMARY KEY/PARTITION KEY combination
func GetUserMap_Given_User_Phone(user_phone string) *OrgUserMap_table {
	var orgusermaptable *OrgUserMap_table

	orgusermaptable = new(OrgUserMap_table)
	if orgusermaptable == nil {
		return nil
	}
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrgUserMapKeySpace())
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
	buffer := "select usermap_orgname, usermap_username, usermap_id, usermap_orgtin  from usermapbyphone where usermap_userphone ='" + user_phone + "' ALLOW FILTERING"
	//log the buffer
	fmt.Println("executing::" + buffer)
	err := session_handle.Query(buffer).Scan(&orgusermaptable.Usermap_orgname, &orgusermaptable.Usermap_username, &orgusermaptable.Usermap_id, &orgusermaptable.Usermap_tin)
	if err != nil {
		return nil
	}
	orgusermaptable.Usermap_userphone = user_phone
	//return the pointer to the result
	return orgusermaptable

}
