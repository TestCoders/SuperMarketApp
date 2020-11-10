package main

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func (a *application) routes() http.Handler {
	standardMiddleware := alice.New(a.logRequest)
	r := mux.NewRouter()
	r.HandleFunc("/basket", a.createBasket).Methods("POST")
	r.HandleFunc("/basket/{id}", a.deleteBasket).Methods("DELETE")
	r.HandleFunc("/basket/{id}", a.getBasket).Methods("GET")
	r.HandleFunc("/baskets", a.getBaskets).Methods("GET")
	return standardMiddleware.Then(r)
}
