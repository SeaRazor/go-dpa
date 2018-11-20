package main

import (
	"fmt"
	"net/http"

	"dpa/models"
	"dpa/route"
)



func main() {
	fmt.Printf("Starting DPA microservice at %v...", models.CurrentConfiguration.ServiceUrl)
	router := route.NewRouter()
	http.ListenAndServe(models.CurrentConfiguration.ServiceUrl,router)
	fmt.Println("DPA microservice up and running")

	

}

