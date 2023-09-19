package models

import "time"

type Schema struct {
	ID             int64
	WeekDay        string
	CreationDate   time.Time
	CreationUserID int64
	UpdateDate     time.Time
	UpdateUserID   int64
}

type Account struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type Activity struct {
	ID             int64
	OldDate        time.Time
	NewDate        time.Time
	Paid           string
	Approved       string
	UserID         string
	CreationDate   time.Time
	CreationUserID int64
	UpdateDate     time.Time
	UpdateUserID   int64
}
