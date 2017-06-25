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

type Resultresponse struct {
	Result int `json:"result"`
}

func PingServer(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", PingResponse())

}

//Handles login
func Login(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Reached login URL....")
	fmt.Println("Method: " + req.Method)
	//Only act if it is a POST request
	if req.Method == "OPTIONS" {
		fmt.Println("Entered into Options Method..")
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1000")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		fmt.Fprintf(w, "")
		return
	}
	switch req.Method {
	case "POST":
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		JSONResponse, HeaderValues := LoginRouter(req)
		fmt.Println("Length of Header Value", len(HeaderValues))
		if len(HeaderValues) > 0 {
            for key,value := range HeaderValues {
				fmt.Println("Key: ", key, "Value: ", value)
				utils.SetHttpHeaderValues(w, key, value)
			}
		}
		fmt.Println(JSONResponse)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		utils.SendHttpResponse(w, JSONResponse)
	case "OPTIONS":
		fmt.Println("Entered into Options Method..")
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1000")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		fmt.Fprintf(w, "")
		//body, err := ioutil.ReadAll(req.Body)
		//if err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(string(body))
	default:
	}

}

func Products(w http.ResponseWriter, req *http.Request) {
	var authorization string
	fmt.Println("Reached Customers URL....")
	fmt.Println("Method: " + req.Method)
	origin := req.Header.Get("Origin")
	fmt.Println(origin)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request
	//Check the authorization code here, if none, reject the offer
	if req.Method == "OPTIONS" {
		fmt.Println("Entered into Options Method..")
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1000")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		fmt.Fprintf(w, "")
		return
	}
	authorization = req.Header.Get("Authorization")
	fmt.Println("Authorization: ", authorization)

	if authorization == "" {
		fmt.Fprintf(w, "No Authorization")
		return
	}
	switch req.Method {
	case "GET":
		resp := Fetch_Products_Given_AuthorizationCode(authorization)
		utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		utils.SendHttpResponse(w, resp)
	case "POST":
		var prod datamodels.Product_table

		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))

		err := json.Unmarshal(body, &prod)
		if err != nil {
			fmt.Println("Error in unmarshing the add customer request...")
			return
		}
		fmt.Println(prod)
		resp:= Add_Products_Given_AuthorizationCode(authorization, prod)
		var rr Resultresponse
		if resp == true {

			rr= Resultresponse{Result: 1}

		} else {
			rr= Resultresponse{Result: 0}

		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		respon, errs := json.Marshal(rr)
		if errs != nil {
			return
		}
		utils.SendHttpResponse(w, respon)
	}
}

func Vehicles(w http.ResponseWriter, req *http.Request) {
	var authorization string
	fmt.Println("Reached Customers URL....")
	fmt.Println("Method: " + req.Method)
	origin := req.Header.Get("Origin")
	fmt.Println(origin)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request
	//Check the authorization code here, if none, reject the offer
	if req.Method == "OPTIONS" {
		fmt.Println("Entered into Options Method..")
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1000")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		fmt.Fprintf(w, "")
		return
	}
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
		w.Header().Set("Access-Control-Allow-Origin", origin)
		utils.SendHttpResponse(w, resp)
	case "POST":
		var vehl datamodels.Vehicle_table

		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))

		err := json.Unmarshal(body, &vehl)
		if err != nil {
			fmt.Println("Error in unmarshing the add customer request...")
			return
		}
		fmt.Println(vehl)
		resp:= Add_Vehicle_Given_AuthorizationCode(authorization, vehl)
		var rr Resultresponse
		if resp == true {

			rr= Resultresponse{Result: 1}

		} else {
			rr= Resultresponse{Result: 0}

		}
		w.Header().Set("Access-Control-Allow-Origin", origin)
		respon, errs := json.Marshal(rr)
		if errs != nil {
			return
		}
		utils.SendHttpResponse(w, respon)
	}

}
func Customers(w http.ResponseWriter, req *http.Request) {
    var authorization string
	origin := req.Header.Get("Origin")
	fmt.Print(origin)
	fmt.Println("Reached Customers URL....")
	fmt.Println("Method: " + req.Method)
	//Verify user, if not a valid user(valid interms of session & ability to do deliverylog GET, POST), send error
	//If valid user, then use orgID to fulfill the request
	//Check the authorization code here, if none, reject the offer
	if req.Method == "OPTIONS" {
		fmt.Println("Entered into Options Method..")
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1000")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
        fmt.Println("Customers::sending OPTIONS headers...")
		fmt.Fprintf(w, "")
		return
	}
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
		       w.Header().Set("Access-Control-Allow-Origin", origin)
		       utils.SendHttpResponse(w, resp)
	case "POST":
		      var cust datamodels.Customer_table

		      body, _ := ioutil.ReadAll(req.Body)
		      fmt.Println(string(body))

			  err := json.Unmarshal(body, &cust)
			  if err != nil {
				  fmt.Println("Error in unmarshing the add customer request...")
				  return
			  }
		      fmt.Println(cust)
		      resp:= Add_Customers_Given_AuthorizationCode(authorization, cust)
		      var rr Resultresponse
		      if resp == true {

				  rr= Resultresponse{Result: 1}

			  } else {
				  rr= Resultresponse{Result: 0}

			  }
		      w.Header().Set("Access-Control-Allow-Origin", origin)
		      respon, errs := json.Marshal(rr)
		      if errs != nil {
			      return
		      }
		      utils.SendHttpResponse(w, respon)
	}

}

//Handles deliverylog
func DeliveryLog(w http.ResponseWriter, req *http.Request) {
	var authorization string
    //TODO: The user should have all of the values of the logged in user as part of Authorization
	origin := req.Header.Get("Origin")
	fmt.Print(origin)
	fmt.Println("Reached deliverylog URL....")
	fmt.Println("Method: " + req.Method)
	if req.Method == "OPTIONS" {
		fmt.Println("Entered into Options Method..")
		origin := req.Header.Get("Origin")
		fmt.Println(origin)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "1000")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")

		fmt.Fprintf(w, "")
		return
	}
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
		search := req.URL.Query().Get("search")
		customer := req.URL.Query().Get("customer")
		fmt.Println(getall)
		if getall == "true" {
			resp := Fetch_Delivery_Logs_Given_authorizationCode(authorization)
			//mt.Println(string(resp))
			utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
			w.Header().Set("Access-Control-Allow-Origin", origin)
			utils.SendHttpResponse(w, resp)
		} else if getall=="false" {
			resp := Fetch_Delivery_Logs_Given_authorizationCode_Limit(authorization, 100)
			//mt.Println(string(resp))
			utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
			w.Header().Set("Access-Control-Allow-Origin", origin)
			utils.SendHttpResponse(w, resp)
		} else if search == "true" {
			resp := Search_Delivery_Logs_Given_authorizationCode(authorization, customer)
			utils.SetHttpHeaderValues(w, cfg.HTTP_HEADER_CONTENT_TYPE, cfg.HTTP_HEADER_DATATYPE_JSON)
			w.Header().Set("Access-Control-Allow-Origin", origin)
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
				  w.Header().Set("Access-Control-Allow-Origin", origin)
				  var rr = Resultresponse{Result: 0}
				  resp, err := json.Marshal(rr)
				  if err != nil {
					  return
				  }
				  utils.SendHttpResponse(w, resp)

			  } else {
				  w.Header().Set("Access-Control-Allow-Origin", origin)
				  var rr = Resultresponse{Result: 1}
				  resp, err := json.Marshal(rr)
				  if err != nil {
					  return
				  }
				  utils.SendHttpResponse(w, resp)
			  }
		case 5: //search delivery request(always based on customer for now.
			fmt.Println("Delivery log: ", dl)

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
