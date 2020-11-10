package main

import (
	"github.com/marcelblijleven/SuperMarketApp/basket"
	"net/http"
)

func main() {
	app := newApplication()
	app.basketManager = basket.NewManager(basket.NewInMemoryBasketRepository())

	srv := &http.Server{
		Addr:     ":3000",
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}

	app.LogInfo("starting server")
	err := srv.ListenAndServe()
	app.LogError(err)
}
