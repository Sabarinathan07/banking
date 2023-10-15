package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name       string `json:"full_name" xml:"name"`
	City      string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}
type Student struct {
	Name       string `json:"full_name" xml:"name"`
	Email      string `json:"email" xml:"email"`
	Department string `json:"department" xml:"department"`
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}


func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Sabari", "Chennai", "600122"},
		{"Ram", "rajastan", "511456"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
 	fmt.Fprint(w, vars["customer_id"])

}

func getAllStudents(w http.ResponseWriter, r *http.Request) {
	students := []Student{
		{"Sabari", "sabari@gmail.com", "CSE"},
		{"Ram", "ram@gmail.com", "CSE"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(students)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(students)
	}
}
