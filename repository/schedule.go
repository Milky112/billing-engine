package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"time"

	"amartha.com/billing/model"
)

func (r *Repository) GetDatabase() *sql.DB {
	return r.Db
}

func (r *Repository) GetTransactionScheduleByUserID(ctx context.Context, userID int, maxSchedule int) ([]model.TransactionSchedule, error) {

	var (
		result []model.TransactionSchedule
	)

	query := fmt.Sprintf(queryGetScheduleByUserID, maxSchedule)

	rows, err := r.Db.QueryContext(ctx, query, userID)
	if rows == nil {
		return result, nil
	}
	defer rows.Close()

	for rows.Next() {
		var row model.TransactionSchedule

		err := rows.Scan(&row.ScheduleID, &row.Status, &row.PaymentDate, &row.TransactionID, &row.PaymentAmount)
		if err != nil {
			return []model.TransactionSchedule{}, err
		}

		result = append(result, row)
	}

	if err != nil {
		return result, err
	}

	return result, nil

}

func (r *Repository) UpdateScheduleByID(ctx context.Context, scheduleID int) error {

	//prepare the statement
	stmts, _ := r.Db.Prepare(queryUpdateSchedule)
	res, err := stmts.Exec(scheduleID)
	fmt.Println(res, err)
	return nil
}

func (r *Repository) InsertLoanSchedule(ctx context.Context, userID int, transactionID int) error {

	stmt, vals := convertQuery(userID, transactionID)
	stmt = strings.TrimSuffix(stmt, ",")

	//prepare the statement
	stmts, _ := r.Db.Prepare(stmt)

	res, err := stmts.Exec(vals...)
	fmt.Println(res, err)
	return nil
}

func convertQuery(userID int, transactionID int) (string, []interface{}) {

	vals := []interface{}{}
	timeNow := time.Now()
	query := queryInsertLoanScheduleByUserID
	for i := 0; i < 5; i++ {
		query += "(?, ?, ?, ?),"
		vals = append(vals, transactionID, userID, timeNow, 1)
		timeNow = timeNow.AddDate(0, 0, 7)
	}

	return query, vals
}
