package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/marcelblijleven/SuperMarketApp/basket"
	"github.com/marcelblijleven/SuperMarketApp/checkout"
	"github.com/marcelblijleven/SuperMarketApp/inventory"
	"io/ioutil"
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

	b, err := a.basketManager.Get(id)

	if err != nil {
		a.LogInfo(err.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")

	if err = e.Encode(&b); err != nil {
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

func (a *application) updateBasket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	data, err := ioutil.ReadAll(r.Body)

	if err != nil {
		a.serveError(w, err, http.StatusBadRequest)
		return
	}

	var body updateBasketRequest

	if err := json.Unmarshal(data, &body); err != nil {
		a.serveError(w, err, http.StatusBadRequest)
		return
	}

	b, err := a.basketManager.Get(id)

	if err != nil {
		if err == basket.ErrBasketNotFound {
			a.serveError(w, err, http.StatusNotFound)
			return
		}

		a.serveError(w, err, http.StatusInternalServerError)
		return
	}

	// Find product
	p, err := a.inventoryManager.Get(body.Barcode)

	if err != nil {
		if err == inventory.ErrProductNotFound {
			a.serveError(w, err, http.StatusBadRequest)
			return
		}

		a.serveError(w, err, http.StatusInternalServerError)
		return
	}

	line := productToProductLine(p, body.Quantity)
	b.UpdateProductLine(line)

	if err = a.basketManager.Save(b); err != nil {
		a.serveError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	e := json.NewEncoder(w)
	e.SetIndent("", "\t")

	if err = e.Encode(&b); err != nil {
		a.serveError(w, err, http.StatusInternalServerError)
		return
	}
}

type updateBasketRequest struct {
	Barcode  string `json:"barcode"`
	Quantity int    `json:"quantity"`
}

func (a *application) checkout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	b, err := a.basketManager.Get(id)

	if err != nil {
		if err == basket.ErrBasketNotFound {
			a.serveError(w, err, http.StatusNotFound)
			return
		}

		a.serveError(w, err, http.StatusInternalServerError)
	}

	// Retrieve order from basket
	order := checkout.MakePurchase(&b)

	// Remove basket
	if err = a.basketManager.Delete(id); err != nil {
		a.serveError(w, err, http.StatusInternalServerError)
		return
	}

	e := json.NewEncoder(w)
	e.SetIndent("", "\t")
	if err = e.Encode(&order); err != nil {
		a.serveError(w, err, http.StatusInternalServerError)
		return
	}
}
