package main

import (
	"fmt"
	"iAccounts/handlers"
	"log"
	"net/http"
	"iAccounts/cfg"
	"time"
)


func main() {

	fmt.Println("Initializing iAccounts Server....Please wait....")

	http.HandleFunc("/api/v1/ping", handlers.PingServer)
	http.HandleFunc("/api/v1/login", handlers.Login)
	http.HandleFunc("/api/v1/deliverylog", handlers.DeliveryLog)
	http.HandleFunc("/api/v1/customers", handlers.Customers)
	http.HandleFunc("/api/v1/products", handlers.Products)
	http.HandleFunc("/api/v1/vehicles", handlers.Vehicles)
	cfg.SetStartTime(time.Now())
	err := http.ListenAndServeTLS(cfg.GetHTTPSServerport(), cfg.GetHTTPSTLSCERTIFICATEPath(), cfg.GetHTTPSTLSKEYPath(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
