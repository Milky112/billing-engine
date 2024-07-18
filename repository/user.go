package repository

import (
	"context"
	"database/sql"
	"errors"

	"amartha.com/billing/model"
)

func (r *Repository) GetUserInfo(ctx context.Context, userID int) (model.UserInfo, error) {

	var (
		result model.UserInfo
	)
	result.UserID = userID
	err := r.Db.QueryRowContext(ctx, queryGetUserByID, userID).Scan(&result.Name, &result.Email, &result.Status)

	if err == sql.ErrNoRows {
		return result, errors.New("USER NOT FOUND")
	}

	return result, nil

}
