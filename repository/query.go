package repository

const (
	queryGetScheduleByUserID = `
		SELECT
			id, transaction_id, payment_date, status
			FROM transaction_schedule 
			WHERE user_id = ?
	`

	queryGetUserByID = `
	SELECT
		name, email
		FROM user 
		WHERE user_id = ?
`
)
