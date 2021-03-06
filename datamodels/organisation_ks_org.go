package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
)

type Org_table struct {
	Org_id      string //uuid
	Org_name    string
	Org_tin     string
	Org_email   string
	Org_phone   string
	Org_city    string
	Org_created string
}

func GetORG_Given_authCode(authCode string) (*string, *string) {
	fmt.Println("Auth code: ", authCode)
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrganizationKeySpace())
	if cluster_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired session handle")
	buffer := "select session_phone from sessionbyauthorization where session_authorization='" + authCode + "'"
	var session_phone string
	errs := session_handle.Query(buffer).Scan(&session_phone)
	if errs != nil {
		return nil, nil
	}

	cluster_handle = cassandra.GetClusterHandle(cfg.GetOrgUserMapKeySpace())
	if cluster_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired cluster handle")
	session_handle = cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired session handle")
	buffer = "select usermap_orgname, usermap_orgtin from usermapbyphone where usermap_userphone='" + session_phone + "'"
	fmt.Println(buffer)
	var usermaporgname, usermaporgtin string
	errs = session_handle.Query(buffer).Scan(&usermaporgname, &usermaporgtin)
	fmt.Println("Organization: ", usermaporgname, "TIN: ", usermaporgtin)
	if errs != nil {
		return nil, nil
	}
	cluster_handle = cassandra.GetClusterHandle(cfg.GetOrganizationKeySpace())
	if cluster_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired cluster handle")
	session_handle = cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, nil
	}
	fmt.Println("acquired session handle")
	buffer = "select org_id from org where org_name='" + usermaporgname + "' and org_tin='" + usermaporgtin + "'"

	var orgid string
	errs = session_handle.Query(buffer).Scan(&orgid)
	if errs != nil {
		return nil, nil
	}
	return &orgid, &usermaporgname
}

//Only one entry is expected as these are the PRIMARY KEY/PARTITION KEY combination
func GetORG_Given_Name_Tin(org_name string, org_tin string) *Org_table {
	var orgtable *Org_table

	orgtable = new(Org_table)
	if orgtable == nil {
		return nil
	}
	fmt.Println("created org table")
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrganizationKeySpace()) //change it to config package value
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
	buffer := "select org_id, org_name, org_tin, org_email, org_phone, org_city from org where org_name='" + org_name + "' and org_tin='" + org_tin + "'"
	//log the buffer
	fmt.Println("executing::" + buffer)
	err := session_handle.Query(buffer).Scan(&orgtable.Org_id, &orgtable.Org_name, &orgtable.Org_tin, &orgtable.Org_email, &orgtable.Org_phone, &orgtable.Org_city)
	fmt.Println(err)
	if err != nil {
		return nil
	}
	//return the pointer to the result
	return orgtable
}
