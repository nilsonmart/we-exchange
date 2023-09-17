package models

import "time"

type Schema struct {
	ID             string
	WeekDay        string
	CreationDate   int64
	CreationUserID string
	UpdateDate     int64
	UpdateUserID   string
}

type Account struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type Activity struct {
	ID             string
	OldDate        time.Time
	NewDate        time.Time
	Paid           string
	Approved       string
	UserID         string
	CreationDate   int64
	CreationUserID string
	UpdateDate     int64
	UpdateUserID   string
}
