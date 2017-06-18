package handlers

import (
	"encoding/json"
	"fmt"
	"iAccounts/cfg"
	"iAccounts/datamodels"
	"iAccounts/utils"
	"io/ioutil"
	"net/http"
)

func PingServer(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", PingResponse())

}

//Handles login
func Login(w http.ResponseWriter, req *http.Request) {
    //TODO: Check if this user has already logged in and has got all of the values in the JSON.

	fmt.Println("Reached login URL....")
	fmt.Println("Method: " + req.Method)
	//Only act if it is a POST request
	switch req.Method {
	case "POST":
		JSONResponse := LoginRouter(req)
		utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
		utils.SendHttpResponse(w, JSONResponse)
	default:
	}

}

//Handles deliverylog
func DeliveryLog(w http.ResponseWriter, req *http.Request) {
    //TODO: The user should have all of the values of the logged in user as part of Authorization
	fmt.Println("Reached deliverylog URL....")
	fmt.Println("Method: " + req.Method)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request

	switch req.Method {
	case "GET":

		resp := Fetch_Delivery_Logs_Given_orgID("0b929804-8457-44bd-a2c6-a6d70b2a0c68")
		fmt.Println(string(resp))
		utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
		utils.SendHttpResponse(w, resp)
	case "POST":
		var dl datamodels.Delivery_log
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))

		_ = json.Unmarshal(body, &dl)
		fmt.Println("Login Structure: ")
		fmt.Println(dl)
		//Get the orgID using the userID, sessionID combination
		//resp := Insert_Delivery_Logs_Given_orgID(orgID, dl)
	default:

	}

}
