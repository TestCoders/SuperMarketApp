package main

import (
	"encoding/json"
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
	}
}
