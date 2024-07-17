package repository

import (
	"context"
	"database/sql"

	"amartha.com/billing/model"
)

func (r *Repository) GetDatabase() *sql.DB {
	return r.Db
}

func (r *Repository) GetTransactionScheduleByUserID(ctx context.Context, userID int) ([]model.TransactionSchedule, error) {

	var (
		result []model.TransactionSchedule
	)

	rows, err := r.Db.QueryContext(ctx, queryGetScheduleByUserID, userID)
	if rows == nil {
		return result, nil
	}
	defer rows.Close()

	for rows.Next() {
		var row model.TransactionSchedule

		err := rows.Scan(&row.ScheduleID, &row.Status, &row.PaymentDate, &row.TransactionID)
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
