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

func GetSessionByAuthorization(authorization string ) bool {
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

	buffer := "select session_phone from sessionbyauthorization where session_authorization='" + authorization + "'"

	//log the buffer
	fmt.Println("executing::" + buffer)

	errs := session_handle.Query(buffer).Exec()
	if errs != nil {
		return false
	}

	return true
}

func UpdateSessionwithAuthorizationToken(mobilenumber string, sessionid string, otp string, token string) bool {
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

	buffer := "select session_created from sessionbyauthorization where session_authorization='" +  token + "'"
	//log the buffer
	fmt.Println("executing::" + buffer)
	var session_created string
	errs := session_handle.Query(buffer).Scan(&session_created)
	if errs != nil {
		return false
	}
    buffer = "insert into sessionbyauthorization(session_authorization, session_created) values('" + token + "', '" + session_created + "'"
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()
	if errs != nil {
		return false
	}
	return true
}

//loginjson.Mobile_number, loginjson.Session_id, loginjson.OTP_number
func GetSession(mobilenumber string, sessionid string, otp string) (*string, bool) {
	cluster_handle := cassandra.GetClusterHandle(cfg.GetOrganizationKeySpace()) //change it to config package value
	if cluster_handle == nil {
		return nil, false
	}
	fmt.Println("acquired cluster handle")
	session_handle := cassandra.GetSessionHandle(cluster_handle)
	if session_handle == nil {
		return nil, false
	}
	fmt.Println("acquired session handle")

	buffer := "select session_authorization from sessionbyidnumberotp where session_sessionid=" + sessionid + " and session_phone='" + mobilenumber + "' and session_otp='" + otp + "'"
	//log the buffer
	fmt.Println("executing::" + buffer)
	var session_authorization string
	errs := session_handle.Query(buffer).Scan(&session_authorization)
	if errs != nil {
		return nil, false
	}
	return &session_authorization, true
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
	//Create a random string to make sure I have some initial value
	randstr := utils.GenerateSecureSessionID()

  	buffer := "insert into sessionbysessionid (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone, session_authorization) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "','Bearer " + randstr + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs := session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	buffer = "insert into sessionbyphone (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone, session_authorization) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "','Bearer " + randstr + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	buffer = "insert into sessionbyidnumberotp (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone, session_authorization) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "','Bearer " + randstr + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	buffer = "insert into sessionbysessionidphone (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone, session_authorization) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "','Bearer " + randstr + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}

	buffer = "insert into sessionbyauthorization (session_id, session_sessionid, session_created, session_expired, session_loggedin, session_otp, session_phone, session_authorization) values (" + utils.GenerateSecureSessionID() + "," + sessiontable.Session_sessionid + ",'" + sessiontable.Session_created + "'," + "false" + "," + "false" + ",'" + sessiontable.Session_sessionotp + "','" + sessiontable.Session_phone + "','Bearer " + randstr + "')"
	//log the buffer
	fmt.Println("executing::" + buffer)
	errs = session_handle.Query(buffer).Exec()

	if errs != nil {
		return false
	}
	//return the pointer to the result
	return true
}
