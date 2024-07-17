package repository

import (
	"context"
	"database/sql"

	"amartha.com/billing/model"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

type RepositoryInterface interface {
	GetDatabase() *sql.DB

	GetTransactionScheduleByUserID(ctx context.Context, userID int) ([]model.TransactionSchedule, error)
	GetUserInfo(ctx context.Context, userID int) (model.UserInfo, error)
}

func NewRepository(opts NewRepositoryOptions) *Repository {

	db, err := sql.Open("mysql", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db: db,
	}
}
