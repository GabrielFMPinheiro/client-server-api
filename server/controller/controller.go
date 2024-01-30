package controller

import (
	"github.com/GabrielFMPinheiro/client-server-api/server/controller/exchangecontroller"
	"github.com/GabrielFMPinheiro/client-server-api/server/repository"
)

// Controllers contains all the controllers
type Controllers struct {
	ExchangeController *exchangecontroller.Controller
}

// InitControllers returns a new Controllers
func InitControllers(repositories *repository.Repositories) *Controllers {
	return &Controllers{
		ExchangeController: exchangecontroller.InitController(repositories.ExchangeRepo),
	}
}
