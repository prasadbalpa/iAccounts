package main

import (
	"fmt"
	"iAccounts/handlers"
	"log"
	"net/http"
	"iAccounts/cfg"
	"time"
        "os"
)


func main() {

	fmt.Println("Initializing iAccounts Web Server....Please wait....")
        fmt.Println("Asking me to run at..." + os.Args[1])
	http.HandleFunc("/api/v1/ping", handlers.PingServer)
	http.HandleFunc("/api/v1/login", handlers.Login)
	http.HandleFunc("/api/v1/deliverylog", handlers.DeliveryLog)
	http.HandleFunc("/api/v1/customers", handlers.Customers)
	http.HandleFunc("/api/v1/products", handlers.Products)
	http.HandleFunc("/api/v1/vehicles", handlers.Vehicles)
	cfg.SetStartTime(time.Now())
	err := http.ListenAndServeTLS(":" + os.Args[1] /*cfg.GetHTTPSServerport()*/, cfg.GetHTTPSTLSCERTIFICATEPath(), cfg.GetHTTPSTLSKEYPath(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
