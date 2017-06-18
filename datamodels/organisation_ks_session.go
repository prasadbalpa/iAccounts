package datamodels

import (
	"fmt"
	"iAccounts/cassandra"
	"iAccounts/cfg"
	"iAccounts/utils"
)

type Session_table struct {
	Session_clientid   string
	Session_created    string
	Session_clientip   string
	Session_expired    bool
	Session_sessionid  string
	Session_sessionotp string
	Session_phone      string
	Session_loggedin   bool
}

//loginjson.Mobile_number, loginjson.Session_id, loginjson.OTP_number
func GetSession(mobilenumber string, sessionid string, otp string) bool {
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrganizationKeySpace()) //change it to config package value
	if cluster_handle == nil {
		return false
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return false
	}
	fmt.Println("acquired session handle")

	buffer := "select * from sessionbyidnumberotp where session_sessionid=" + sessionid + " and session_phone='" + mobilenumber + "' and session_otp='" + otp + "'"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs := session_handle.Query(buffer).Exec()
	if errs != nil {
		return false
	}
	return true
}

func InsertSession(sessiontable *Session_table) bool {

	if sessiontable == nil {
		return false
	}
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrganizationKeySpace()) //change it to config package value
	if cluster_handle == nil {
		return false
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return false
	}
	fmt.Println("acquired session handle")

	buffer := "insert into sessionbysessionid (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into sessionbyphone (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	buffer = "insert into sessionbyidnumberopt (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	buffer = "insert into sessionbysessionidphone (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	//return the pointer to the result
	return true
}
