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
	Token string `json:"token,omitempty"`
}

type Logininput struct {
	Mobilenumber string `json:"mobilenumber"`
}

//All of the /login requests will land here to understand first where we are with the login for a particular user
func LoginRouter(Request *http.Request) ([]byte, map[string]string)  {
	var loginjson LoginJSON
    HeaderValues := make(map[string]string)
    authorization := Request.Header.Get("Authorization")
	fmt.Println("Authorization header : ", authorization)

	//Check if authorization code is stored in the session table
	//If authorization code is already existing, return saying "already logged in"
	//loggedin := VerifyAuthorizationCode(authorization)
	//if loggedin == true {
		//Return
	//	return []byte("hello great job"), nil
	//}
	body, err := ioutil.ReadAll(Request.Body)
	fmt.Println(string(body))
	if err != nil {
		//send error json for login
		return nil, nil
	}
	fmt.Println(body)
	err = json.Unmarshal(body, &loginjson)
	fmt.Println("Login Structure: ")
	fmt.Println(loginjson)
	if err != nil {
		//send error json for login
		fmt.Println("Error while unmarshing the json from the client...returning.")
		return nil, nil
	}

	if loginjson.Request_type == cfg.LOGIN_MOBILE_NUMBER {
		if loginjson.Mobile_number != "" {
			fmt.Println(loginjson.Mobile_number)

			resp := VerifyuserIDAndGenerateOTP(loginjson)
			HeaderValues["Content-type"] = "application/json"
			return resp, HeaderValues
		} else {
			return nil, nil
		}
	} else if loginjson.Request_type == cfg.LOGIN_OTP_NUMBER {
		fmt.Println(loginjson.OTP_number)
		otpresp := VerifyuserIDotpsessionidAndLogin(loginjson)
        HeaderValues["Content-type"] = "application/json"

		return []byte(otpresp), HeaderValues
	} else {
		fmt.Println("No request type has been mentioned..")
		return nil, nil
	}
	return nil, nil
}

func VerifyAuthorizationCode(authorization string) bool {
	if false == datamodels.GetSessionByAuthorization(authorization) {
		return false
	}
	return true
}

func VerifyuserIDotpsessionidAndLogin(loginjson LoginJSON) []byte {

	authorization, errs := datamodels.GetSession(loginjson.Mobile_number, loginjson.Session_id, loginjson.OTP_number)

	if errs == false {
		fmt.Println("Failed to get the session")
		return nil
	}

	//TODO: Generate authorization token & Pass it on.  Henceforth all of the requests will be verified against this token
    //TODO: Update the session table with authorization token (Token format : "Bearer<space><token>"
	//if false == datamodels.UpdateSessionwithAuthorizationToken(loginjson.Mobile_number, loginjson.Session_id, loginjson.OTP_number, *authorization) {
		//fmt.Println("Failed to update the authorization token")
		//return nil
	//}
	var loginresponse = LoginResponse{Response_type: cfg.LOGIN_SUCCESSFUL, Mobile_number: loginjson.Mobile_number, Session_id: loginjson.Session_id, Token: *authorization}

	logresp, errss := json.Marshal(&loginresponse)
	if errss != nil {
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
    var OTP string
	OTP = utils.Get6DigitsRandomNumbers()
	var stable = datamodels.Session_table{Session_phone: loginjson.Mobile_number, Session_sessionid: utils.GenerateSecureSessionID(), Session_created: utils.GenerateUnixTimeStamp(), Session_expired: false, Session_loggedin: false, Session_sessionotp: OTP}
	err := datamodels.InsertSession(&stable)
	if err == false {
		return nil
	}
	//Send OTP to the phone as it is established that the user exists and has an associated ORG
    //go utils.SendSMS("+91" + loginjson.Mobile_number,  OTP)
	//Send it over as SMS
	var loginresponse = LoginResponse{Response_type: cfg.LOGIN_OTP_NUMBER, Mobile_number: loginjson.Mobile_number, Session_id: stable.Session_sessionid}

	logresp, errs := json.Marshal(&loginresponse)
	fmt.Println("JSON after Marshaling: ", logresp)
	fmt.Println("Error after Marshaling", errs)
	if errs != nil {
		//send error json for login
		return nil
	}
	return logresp

}
