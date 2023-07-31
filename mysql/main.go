package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Bookingdetails struct {
	Bookingid    string           `json:"booking_id"`
	Isaccept     int32            `json:"is_accept"`
	Ncatid       string           `json:"ncat_id"`
	Service_name []Servicedetails `json:"service_name"`
}
type Servicedetails struct {
	Service_id   string `json:"service_id"`
	Service_name string `json:"service_name"`
	Quantity     string `json:"quantity"`
	Price        string `json:"price"`
	Ser_flag     string `json:"ser_flag"`
}

var Servicename string

func main() {

	http.HandleFunc("/", Fetchbookingdetails)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func Fetchbookingdetails(w http.ResponseWriter, _ *http.Request) {
	db, err := sql.Open("mysql", "web_app:!5@uGuST1($7FrEe1Ndi@@tcp(192.168.6.180)/myjd?parseTime=true")
	if err != nil {
		fmt.Println("error-----")
	}
	defer db.Close()
	// Execute the query
	results, err := db.Query("select booking_id, is_accept, ncat_id, service_name from myjd.tbl_pending_booking_details order by create_date desc limit 10")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var bookingresponse []Bookingdetails
	var servicedetails []Servicedetails

	for results.Next() {
		var bookingdetails Bookingdetails
		err = results.Scan(&bookingdetails.Bookingid, &bookingdetails.Isaccept, &bookingdetails.Ncatid, &Servicename)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		json.Unmarshal([]byte(Servicename), &servicedetails)
		for _, servicedata := range servicedetails {
			bookingdetails.Service_name = append(servicedetails, servicedata)
		}
		bookingresponse = append(bookingresponse, bookingdetails)
	}
	response, _ := json.Marshal(bookingresponse)

	fmt.Fprintf(w, string(response))
}
