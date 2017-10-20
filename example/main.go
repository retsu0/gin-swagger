package main

import (
	"log"

	"github.com/retsu0/gin-swagger/example/restapi"
)

var apiConfig restapi.Config

func main() {
	err := apiConfig.Parse()
	if err != nil {
		log.Fatal(err)
	}

	svc := &ExampleService{Health: false}

	api := restapi.NewAPI(svc, &apiConfig)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}
}
