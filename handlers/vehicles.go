package handlers

import (
	"iAccounts/datamodels"
	"fmt"
	"encoding/json"
)

type Vehicle_response struct {
	Vehicle_id string `json:"id"`
	Vehicle_number string `json:"number"`

}

type vehresponse struct {
	Response_type int                `json:"responsetype"`
	Organization  string             `json:"organization"`
	Veh           []Vehicle_response `json:"vehicles"`
}

func Fetch_Vehicles_Given_AuthorizationCode(authorization string) []byte {
	var vehicle_response []Vehicle_response
	var vehicle_log []datamodels.Vehicle_table
	var orgname *string
	vehicle_log, orgname = datamodels.GetVehiclesByAuthCode(authorization)

	if vehicle_log == nil {
		return nil
	}
	vehicle_response = make([]Vehicle_response, len(vehicle_log))
	for i:=0;i<len(vehicle_log);i++ {
		vehicle_response[i].Vehicle_id = vehicle_log[i].Vehicle_id
		vehicle_response[i].Vehicle_number = vehicle_log[i].Vehicle_number

	}
	var veh vehresponse
	veh.Veh = make([]Vehicle_response, len(vehicle_log))
	for i := 0; i < len(vehicle_log); i++ {
		veh.Veh[i] = vehicle_response[i]
	}
	veh.Response_type = 10
	veh.Organization = *orgname
	fmt.Println("Marshalling to JSON...")

	vehresp, errs := json.Marshal(veh)
	fmt.Println(string(vehresp))
	if errs != nil {
		//send error json for login
		fmt.Println("Failed to Marshal..")
		return nil
	}
	return vehresp
}

