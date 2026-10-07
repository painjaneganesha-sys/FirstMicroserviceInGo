package main

import (
	"log"
	"net/http"

	"FirstMicroserviceInGo/api"
	_ "FirstMicroserviceInGo/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title First Microservice
// @version 1.0
// @description Simple Calculator Microservice
// @host localhost:8080
// @BasePath /
func main() {

	calculatorHandler := &api.CalculatorHandler{}

	http.Handle("/calculate", calculatorHandler)

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}