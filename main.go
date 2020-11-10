package main

import (
	"github.com/marcelblijleven/SuperMarketApp/basket"
	"github.com/marcelblijleven/SuperMarketApp/inventory"
	"net/http"
)

func main() {
	app := newApplication()
	app.basketManager = basket.NewManager(basket.NewInMemoryBasketRepository())
	app.inventoryManager = inventory.NewManager(inventory.NewInMemoryProductRepository())

	srv := &http.Server{
		Addr:     ":3000",
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}

	app.LogInfo("starting server")
	err := srv.ListenAndServe()
	app.LogError(err)
}
