package repository

import (
	"context"
	"database/sql"

	"amartha.com/billing/model"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	Db *sql.DB
}

type NewRepositoryOptions struct {
	Dsn string
}

type RepositoryInterface interface {
	GetDatabase() *sql.DB

	GetTransactionScheduleByUserID(ctx context.Context, userID int, maxSchedule int) ([]model.TransactionSchedule, error)
	GetUserInfo(ctx context.Context, userID int) (model.UserInfo, error)
	InsertLoanSchedule(ctx context.Context, userID int, transactionID int) error
	InsertTransaction(ctx context.Context, userID int, amount int, interest int) (int, error)
	UpdateTransactionByID(ctx context.Context, loanPayment, transactionID int) error
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
