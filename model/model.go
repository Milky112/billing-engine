package model

import (
	"time"
)

type TransactionScheduleResponse struct {
	Schedule []TransactionSchedule `json:"transaction_schedule"`
	User     UserInfo              `json:"user_info"`
}
type TransactionSchedule struct {
	ScheduleID    int       `db:"id"`
	PaymentDate   time.Time `db:"payment_date"`
	Status        int       `db:"status"`
	User          int       `db:"user_id"`
	TransactionID int       `db:"transaction_id"`
}

type UserInfo struct {
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
	Email  string `db:"email"`
}

type Transaction struct {
	TransactionID int `db:"transaction_id"`
}

type Config struct {
	Database Database `yaml:"database"`
	User     int      `yaml:"user"`
}

type Database struct {
	DatabaseConn     string `yaml:"dabatase_conn"`
	DatabaseUser     string `yaml:"database_user"`
	DatabasePassword string `yaml:"database_password"`
	DatabasePort     string `yaml:"database_port"`
}
