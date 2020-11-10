package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (a *application) createBasket(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	basket, err := a.basketManager.Create()

	if err != nil {
		a.LogError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")

	if err = e.Encode(&basket); err != nil {
		a.LogError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *application) deleteBasket(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_ = a.basketManager.Delete(id)

	w.WriteHeader(http.StatusNoContent)
}

func (a *application) getBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	basket, err := a.basketManager.Get(id)

	if err != nil {
		a.LogInfo(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")

	if err = e.Encode(&basket); err != nil {
		a.LogError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (a *application) getBaskets(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	baskets, err := a.basketManager.GetAll()

	if err != nil {
		a.LogError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")

	if err = e.Encode(&baskets); err != nil {
		a.LogError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
