package repository

import (
	"database/sql"

	exchangerepo "github.com/GabrielFMPinheiro/client-server-api/server/repository/exchangerepo"
)

type Repositories struct {
	ExchangeRepo *exchangerepo.ExchangeRepo
}

func InitRepositories(db *sql.DB) *Repositories {
	exchangeRepo := exchangerepo.NewExchangeRepo(db)
	return &Repositories{ExchangeRepo: exchangeRepo}
}
