package settings

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	DBClient *pgxpool.Pool // this is variable because we will assign new value once db is connected
)

func BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := DBClient.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	log.Println("Beign Transaction")

	return tx, nil

}

func CommitTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Commit(ctx)
	if err != nil {
		return err
	}
	log.Println("Commit Transaction")

	return nil
}

func RollbackTx(ctx context.Context, tx pgx.Tx) error {
	err := tx.Rollback(ctx)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	log.Println("Rollback Transaction")

	return nil
}
