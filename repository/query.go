package repository

const (
	queryGetScheduleByUserID = `
		SELECT
			id, transaction_id, payment_date, status, payment_loan
			FROM transaction_schedule 
			WHERE user_id = ? 
			ORDER BY payment_date ASC LIMIT %d
	`

	queryGetUserByID = `
	SELECT
		name, email, status
		FROM user 
		WHERE user_id = ?
	`

	queryInsertLoanScheduleByUserID = `
		INSERT INTO transaction_schedule (transaction_id, user_id, payment_date, status) VALUES 
	`

	queryInsertTransaction = `
		INSERT INTO transaction (user_id, loan_total, loan_interest, loan_payment) 
		VALUES(?, ?, ?, ?);
	`

	queryUpdateSchedule = `
		UPDATE transaction_schedule set STATUS = 0 where id = ?
	`

	queryPaymentTransaction = `
		UPDATE transaction set loan_payment = loan_payment + ? where id = ?
	`
)
