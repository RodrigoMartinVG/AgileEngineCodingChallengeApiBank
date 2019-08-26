package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rodrigomartinvg/AgileEngineCodingChallengeApiBank/api"
	"github.com/rodrigomartinvg/AgileEngineCodingChallengeApiBank/templates"
)

func main() {
	templates.Init()
	
	r :=  mux.NewRouter().StrictSlash(true)
	
	r.Methods("POST").Path("/transactions").Handler(api.NewEndpoint(api.HandleTransactionPost, api.ResponseJSON))
	r.Methods("GET").Path("/transactions/{id}").Handler(api.NewEndpoint(api.HandleTransactionGet, api.ResponseJSON))
	r.Methods("GET").Path("/transactions").Handler(api.NewEndpoint(api.HandleTransactionList, api.ResponseJSON))
	r.Methods("GET").Path("/").Handler(api.NewEndpoint(api.HandleDefault, api.ResponseJSON))
	r.Methods("GET").Path("/ui").Handler(api.NewEndpoint(api.HandleUI, api.ResponseHTML))
	
	log.Println("ListenAndServe:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	
}
