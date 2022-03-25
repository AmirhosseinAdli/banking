package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{Name: "Amirhossein", City: "Tehran", ZipCode: "021"},
	//	{Name: "Mehran", City: "Shooshtar", ZipCode: "0256"},
	//}

	var customer1 Customer
	customer1.Name = "Amirhossein"
	customer1.City = "Tehran"
	customer1.ZipCode = "021"

	var customer2 Customer
	customer2.Name = "Mehran"
	customer2.City = "Shooshtar"
	customer2.ZipCode = "0256"

	customers := []Customer{customer1, customer2}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
