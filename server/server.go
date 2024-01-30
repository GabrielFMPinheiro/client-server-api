package main

import (
	"net/http"

	"github.com/GabrielFMPinheiro/client-server-api/server/controller"
	"github.com/GabrielFMPinheiro/client-server-api/server/database"
	"github.com/GabrielFMPinheiro/client-server-api/server/repository"
)

func main() {
	db := database.InitDatabase()
	defer db.Close()
	repos := repository.InitRepositories(db)
	controller := controller.InitControllers(repos)

	http.HandleFunc("/cotacao", controller.ExchangeController.GetExchange)

	http.ListenAndServe(":8080", nil)
}
