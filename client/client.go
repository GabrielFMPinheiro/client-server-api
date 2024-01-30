package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/GabrielFMPinheiro//client-server-api/client/dto/exchangedto"
)

func saveBid(bid string) {
	file, err := os.Create("cotacao.txt")

	if err != nil {
		fmt.Println("Error to create file")
	}

	defer file.Close()

	_, err = io.WriteString(file, "Dólar: "+bid)

	if err != nil {
		fmt.Println("Error to write file")
	}
}

func main() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancelCtx()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)

	if err != nil {
		fmt.Println("Error to create request")
	}

	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Timeout while trying to get data from the server")
		}
		fmt.Println("Error to get response")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	fmt.Print(string(body))

	var exchange exchangedto.ExchangeDTO
	json.Unmarshal(body, &exchange)
	saveBid(exchange.USDBRL.Bid)
}
