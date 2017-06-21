package utils

import (
	"fmt"
	"net/url"
	"strings"
	"net/http"
)

const (
	TWILIO_ACCOUNT_SID="AC718c4003d481e9dbdc45d9096bb8ddf0"
	TWILIO_TEST_AUTHTOKEN="e53f4e7d62b9697226be5464597fb877"
	TWILIO_NUMBER="+12015754077"

)

func SendSMS(To string, Msg string ) {
	accountSid := TWILIO_ACCOUNT_SID
	authToken := TWILIO_TEST_AUTHTOKEN
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	fmt.Println(urlStr)
	fmt.Println("To: ", To, "Msg: ", Msg)
	// Build out the data for our message
	v := url.Values{}
	v.Set("To",To)
	v.Set("From",TWILIO_NUMBER)
	v.Set("Body","iAccounts Verification Code: " + Msg)
	fmt.Println(v)
	rb := *strings.NewReader(v.Encode())
    fmt.Println(rb)
	// Create client
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

}