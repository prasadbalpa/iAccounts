package utils

import (
	"fmt"
	"net/http"
)


func SetHttpHeaderValues(response http.ResponseWriter, header string, val string) {
    fmt.Println("Header Values: " + header + " : " + val)
	response.Header().Set(header, val)

}

func SendHttpResponse(response http.ResponseWriter, responsedata []byte) {

	fmt.Fprintf(response, string(responsedata))

}
