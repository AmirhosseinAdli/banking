package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/api/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/api/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting servers
	log.Fatal(http.ListenAndServe("127.0.0.1:2000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
