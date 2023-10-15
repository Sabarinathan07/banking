package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabarinathan07/banking/domain"
	"github.com/sabarinathan07/banking/service"
)

func Start(){

	router := mux.NewRouter()


	//wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}


	//defining routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)


	// starting server
	fmt.Println("Server is running at localhost:8000")
    log.Fatal(http.ListenAndServe("localhost:8000", router))
}