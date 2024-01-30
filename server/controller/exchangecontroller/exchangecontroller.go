package exchangecontroller

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielFMPinheiro/client-server-api/server/entity/exchangeentity"
	"github.com/GabrielFMPinheiro/client-server-api/server/integration"
	"github.com/GabrielFMPinheiro/client-server-api/server/repository/exchangerepo"
)

type repository interface {
	CreateExchange(exchange exchangeentity.ExchangeEntity) error
}
type Controller struct {
	service repository
}

func InitController(exchangeRepo *exchangerepo.ExchangeRepo) *Controller {
	return &Controller{
		service: exchangeRepo,
	}
}

func (c *Controller) GetExchange(w http.ResponseWriter, r *http.Request) {
	res, err := integration.GetExchange()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.service.CreateExchange(res.USDBRL)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(res)

}
