package repository

import (
	"context"
	"fmt"
	"log"
)

func (r *Repository) InsertTransaction(ctx context.Context, userID int, amount int, interest int) (int, error) {

	var id int

	// err := r.Db.QueryRow(queryInsertTransaction, userID, amount, interest, 0)

	stmt, err := r.Db.Prepare(queryInsertTransaction)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(userID, amount, interest, 0)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the last inserted ID
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	id = int(lastInsertId)

	return id, nil
}

func (r *Repository) UpdateTransactionByID(ctx context.Context, loanPayment, transactionID int) error {

	//prepare the statement
	stmts, _ := r.Db.Prepare(queryPaymentTransaction)
	res, err := stmts.Exec(loanPayment, transactionID)
	fmt.Println(res, err)
	return nil
}
