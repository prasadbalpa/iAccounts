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

	fmt.Println("Reached login URL....")
	fmt.Println("Method: " + req.Method)
	//Only act if it is a POST request
	switch req.Method {
	case "POST":
		JSONResponse, HeaderValues := LoginRouter(req)
		fmt.Println("Length of Header Value", len(HeaderValues))
		if len(HeaderValues) > 0 {
            for key,value := range HeaderValues {
				fmt.Println("Key: ", key, "Value: ", value)
				utils.SetHttpHeaderValues(w, key, value)
			}
		}
		fmt.Println(JSONResponse)
		utils.SendHttpResponse(w, JSONResponse)
	default:
	}

}
func Vehicles(w http.ResponseWriter, req *http.Request) {
	var authorization string
	fmt.Println("Reached Customers URL....")
	fmt.Println("Method: " + req.Method)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request
	//Check the authorization code here, if none, reject the offer
	authorization = req.Header.Get("Authorization")
	fmt.Println("Authorization: ", authorization)
	if authorization == "" {
		fmt.Fprintf(w, "No Authorization")
		return
	}
	switch req.Method {
	case "GET":
		resp := Fetch_Vehicles_Given_AuthorizationCode(authorization)
		utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
		utils.SendHttpResponse(w, resp)
	}

}
func Customers(w http.ResponseWriter, req *http.Request) {
    var authorization string
	fmt.Println("Reached Customers URL....")
	fmt.Println("Method: " + req.Method)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request
	//Check the authorization code here, if none, reject the offer
	authorization = req.Header.Get("Authorization")
	fmt.Println("Authorization: ", authorization)
	if authorization == "" {
		fmt.Fprintf(w, "No Authorization")
		return
	}
	switch req.Method {
	case "GET":
		       resp := Fetch_Customers_Given_AuthorizationCode(authorization)
		       utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
		       utils.SendHttpResponse(w, resp)
	}

}

//Handles deliverylog
func DeliveryLog(w http.ResponseWriter, req *http.Request) {
	var authorization string
    //TODO: The user should have all of the values of the logged in user as part of Authorization
	fmt.Println("Reached deliverylog URL....")
	fmt.Println("Method: " + req.Method)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request
	//Check the authorization code here, if none, reject the offer
	authorization = req.Header.Get("Authorization")
	fmt.Println("Authorization: ", authorization)
	if authorization == "" {
		fmt.Fprintf(w, "No Authorization")
		return
	}
	switch req.Method {
	case "GET":

		getall := req.URL.Query().Get("getall")
		fmt.Println(getall)
		if getall == "true" {
			resp := Fetch_Delivery_Logs_Given_authorizationCode(authorization)
			//mt.Println(string(resp))
			utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
			utils.SendHttpResponse(w, resp)
		} else {
			resp := Fetch_Delivery_Logs_Given_authorizationCode_Limit(authorization, 100)
			//mt.Println(string(resp))
			utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
			utils.SendHttpResponse(w, resp)
		}
	case "POST":
		var dl datamodels.Delivery_log

		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))

		err := json.Unmarshal(body, &dl)
		if err != nil {
			return
		}
		fmt.Println("Deliverylog Structure: ")
		fmt.Println(dl)
		switch dl.Delivery_requesttype {
		case 4: //add delivery request
			  fmt.Println("Delivery log: ", dl)
			  resp := Add_Delivery_Logs_Given_authorizationCode(authorization, dl)
			  if resp == false {
				  utils.SendHttpResponse(w, []byte("did not work"))
			  } else {
				  utils.SendHttpResponse(w, []byte("It worked!"))
			  }
		}
	case "DELETE":
		//var dldelete datamodels.Deliverylog_delete
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))
		var f interface{}
		err := json.Unmarshal(body, &f)

		if err != nil {
			fmt.Println("Could not unmarshal..")
			return
		}
		var ids []string
		m := f.(map[string]interface{})
		for k, v := range m {
			fmt.Println("Key:", k, "Value:", v)
			switch vv := v.(type) {
			case []interface{}:
								fmt.Println(k, "is an array")
							    for i, u := range vv {
									fmt.Println(i, u)
									l := u.(map[string]interface{})
									ids = append(ids, fmt.Sprintf("%s", l["id"]))
								}
			        			fmt.Println(ids)

			}
		}
		if false == Delete_Delivery_Logs_Given_authorizationCode(authorization, ids) {
			utils.SendHttpResponse(w, []byte("Did not work"))
			return
		}

		utils.SendHttpResponse(w, []byte("It worked!"))
	default:

	}

}
