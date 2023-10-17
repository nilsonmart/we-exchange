package models

import "time"

type DBQuery struct {
	Key   string
	Value string
}

type Account struct {
	ID       string `json:"_id"`
	Name     string `json:"_name"`
	Email    string `json:"_email"`
	Password string `json:"_password"`
}

type RequestChange struct {
	ID             string    `json:"_id"`
	OldDate        time.Time `json:"_olddate"`
	NewDate        time.Time `json:"_newdate"`
	Paid           string    `json:"_paid"`
	Approved       string    `json:"_approved"`
	UserID         string    `json:"_requestedfor"`
	CreationDate   int64     `json:"_creationdate"`
	CreationUserID string    `json:"_creationuserid"`
	UpdateDate     int64     `json:"_updatedate"`
	UpdateUserID   string    `json:"_updateuserid"`
}

type Schema struct {
	ID             string `json:"_id"`
	WeekDay        string `json:"_weekday"`
	CreationDate   int64  `json:"_creationdate"`
	CreationUserID string `json:"_creationuserID"`
	UpdateDate     int64  `json:"_updatedate"`
	UpdateUserID   string `json:"_updateuserid"`
}
