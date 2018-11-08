package main

import (
	
	"fmt"
	"net/http"

	"dpa/route"
	"dpa/models"

	
)


func main() {
	fmt.Println("Starting dpa microservice...")
	router := route.NewRouter()
	http.ListenAndServe(models.CurrentConfiguration.ServiceUrl,router)
	

}

