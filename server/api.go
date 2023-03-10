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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (a *api) calculate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var calc mortgage.Calc
	marErr := json.Unmarshal(reqBody, &calc)

	if marErr != nil {
		http.Error(w, marErr.Error(), http.StatusInternalServerError)
		return
	}

	valErr := (*a.CalculatorService).Execute(&calc)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusInternalServerError)
		return
	}

	if !calc.Valid {
		calc = mortgage.Calc{
			Errors: calc.Errors,
		}
	}

	resErr := json.NewEncoder(w).Encode(calc)
	if resErr != nil {
		http.Error(w, resErr.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *api) Router() http.Handler {
	return a.router
}
