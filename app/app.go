package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()

	// define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/api/customers", getAllCustomers)

	// starting servers
	log.Fatal(http.ListenAndServe("127.0.0.1:2000", mux))
}
