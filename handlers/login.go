package handlers

import (
	"encoding/json"
	"fmt"
	"iAccounts/cfg"
	"iAccounts/datamodels"
	"iAccounts/utils"
	"io/ioutil"
	"net/http"
	"regexp"
)

type LoginJSON struct {
	Request_type  int    `json:"requesttype"`
	Mobile_number string `json:"mobilenumber"`
	OTP_number    string `json:"otpnumber,omitempty"`
	Session_id    string `json:"sessionid,omitempty"`
}

type LoginResponse struct {
	Response_type int    `json:"responsetype"`
	Mobile_number string `json:"mobilenumber"`
	Session_id    string `json:"sessionid"`
}

type Logininput struct {
	Mobilenumber string `json:"mobilenumber"`
}

//All of the /login requests will land here to understand first where we are with the login for a particular user
func LoginRouter(Request *http.Request) []byte {
	var loginjson LoginJSON

	body, err := ioutil.ReadAll(Request.Body)
	fmt.Println(string(body))
	if err != nil {
		//send error json for login
		return nil
	}
	err = json.Unmarshal(body, &loginjson)
	fmt.Println("Login Structure: ")
	fmt.Println(loginjson)
	if err != nil {
		//send error json for login
		fmt.Println("Error while unmarshing the json from the client...returning.")
		return nil
	}

	if loginjson.Request_type == cfg.LOGIN_MOBILE_NUMBER {
		if loginjson.Mobile_number != "" {
			fmt.Println(loginjson.Mobile_number)
			resp := VerifyuserIDAndGenerateOTP(loginjson)
			return resp
		}
	} else if loginjson.Request_type == cfg.LOGIN_OTP_NUMBER {
		fmt.Println(loginjson.OTP_number)
		otpresp := VerifyuserIDotpsessionidAndLogin(loginjson)
		return []byte(otpresp)
	} else {
		return nil
	}
	return nil
}

func VerifyuserIDotpsessionidAndLogin(loginjson LoginJSON) []byte {

	if false == datamodels.GetSession(loginjson.Mobile_number, loginjson.Session_id, loginjson.OTP_number) {
		fmt.Println("Failed to the fetch the session....")
		return nil
	}
	var loginresponse = LoginResponse{Response_type: cfg.LOGIN_SUCCESSFUL, Mobile_number: loginjson.Mobile_number, Session_id: loginjson.Session_id}

	logresp, errs := json.Marshal(&loginresponse)
	if errs != nil {
		//send error json for login
		return nil
	}
	return logresp
}

func VerifyuserIDAndGenerateOTP(loginjson LoginJSON) []byte {

	match, _ := regexp.MatchString("([a-z]+)", loginjson.Mobile_number)
	if match {
		//return error JSON to the view
		return nil
	}

	usermap := datamodels.GetUserMap_Given_User_Phone(loginjson.Mobile_number)
	if usermap == nil {
		return nil
	}

	org := datamodels.GetORG_Given_Name_Tin(usermap.Usermap_orgname, usermap.Usermap_tin)
	if org == nil {
		return nil
	}
	//Enter the session info before OTP is sent out...just to make sure there is no gap between the two.

	var stable = datamodels.Session_table{Session_phone: loginjson.Mobile_number, Session_sessionid: utils.GenerateSecureSessionID(), Session_created: utils.GenerateUnixTimeStamp(), Session_expired: false, Session_loggedin: false, Session_sessionotp: utils.Get6DigitsRandomNumbers()}
	err := datamodels.InsertSession(&stable)
	if err == false {
		return nil
	}
	//Send OTP to the phone as it is established that the user exists and has an associated ORG

	//Send it over as SMS
	var loginresponse = LoginResponse{Response_type: cfg.LOGIN_OTP_NUMBER, Mobile_number: loginjson.Mobile_number, Session_id: stable.Session_sessionid}

	logresp, errs := json.Marshal(&loginresponse)
	if errs != nil {
		//send error json for login
		return nil
	}
	return logresp

}
