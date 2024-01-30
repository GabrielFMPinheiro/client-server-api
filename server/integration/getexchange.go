package integration

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/GabrielFMPinheiro/client-server-api/server/dto/exchangedto"
)

func GetExchange() (exchangedto.ExchangeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		return exchangedto.ExchangeDTO{}, err
	}

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout while trying to get data from the API")
		}
		return exchangedto.ExchangeDTO{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return exchangedto.ExchangeDTO{}, err
	}

	var exchange exchangedto.ExchangeDTO

	err = json.Unmarshal(body, &exchange)

	if err != nil {
		return exchangedto.ExchangeDTO{}, err
	}

	return exchange, nil
}
