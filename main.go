package main

import (
	"github.com/ponsonio/quoter/mortgage"
	"github.com/ponsonio/quoter/server"
	"log"
	"net/http"
)

func main() {
	//this allows to inject a mock or replace the implementation
	calculatorService := mortgage.NewCalculatorService("./mortgage/rules.grl")
	s := server.NewServer(&calculatorService)
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
