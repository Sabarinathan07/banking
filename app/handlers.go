package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabarinathan07/banking/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Sabari", "Chennai", "600122"},
	// 	{"Ram", "rajastan", "511456"},
	// }

	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	// customer, err := ch.service.GetCustomer(id)
	if err != nil {

		writeResponse(w, err.Code, err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, customer)

	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil{
		panic(err)
	}
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

// func getCustomer(w http.ResponseWriter,r *http.Request){
// 	vars := mux.Vars(r)
//  	fmt.Fprint(w, vars["customer_id"])
// }
