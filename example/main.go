package main

import (
	"log"

	"github.com/retsu0/gin-swagger/example/restapi"
)

func main() {

	var apiConfig restapi.Config
	apiConfig.WithDefaultFlags()

	err := apiConfig.Parse()
	if err != nil {
		log.Fatal(err)
	}

	svc := &ExampleService{Health: false}

	api := restapi.NewServer(svc, &apiConfig)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}
}
