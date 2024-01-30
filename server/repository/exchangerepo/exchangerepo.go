package exchangerepo

import (
	"context"
	"database/sql"
	"log"
	"time"

	entity "github.com/GabrielFMPinheiro/client-server-api/server/entity/exchangeentity"
)

type ExchangeRepo struct {
	db *sql.DB
}

func NewExchangeRepo(db *sql.DB) *ExchangeRepo {
	return &ExchangeRepo{
		db: db,
	}
}

func (repo *ExchangeRepo) CreateExchange(exchange entity.ExchangeEntity) error {

	ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelCtx()
	stmt, err := repo.db.PrepareContext(ctx, "INSERT INTO exchange(code,codein,name,high,low,var_bid,pct_change,bid,ask,timestamp,create_date) VALUES(?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("Timeout while trying to insert data into the database")
		}
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, exchange.Code, exchange.Codein, exchange.Name, exchange.High, exchange.Low, exchange.VarBid, exchange.PctChange, exchange.Bid, exchange.Ask, exchange.Timestamp, exchange.CreateDate)

	if err != nil {
		return err
	}

	return nil
}
