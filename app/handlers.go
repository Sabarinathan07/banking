package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Student struct {
	Name       string `json:"full_name" xml:"name"`
	Email      string `json:"email" xml:"email"`
	Department string `json:"department" xml:"department"`
}


func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
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

	//

}
