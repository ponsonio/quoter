package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ponsonio/quoter/mortgage"
	"io"
	"net/http"
)

type api struct {
	router            http.Handler
	CalculatorService *mortgage.CalculatorService
}

type Server interface {
	Router() http.Handler
}

func NewServer(calculatorService *mortgage.CalculatorService) Server {

	a := &api{
		CalculatorService: calculatorService,
	}

	r := mux.NewRouter()

	r.HandleFunc("/mortgage/calculate/", a.calculate).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) calculate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var calc mortgage.Calc
	json.Unmarshal(reqBody, &calc)

	(*a.CalculatorService).Execute(&calc)

	json.NewEncoder(w).Encode(calc)
}

func (a *api) Router() http.Handler {
	return a.router
}
