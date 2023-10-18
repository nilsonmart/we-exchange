package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBQuery struct {
	Key   string
	Value string
}

type Account struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"_name"`
	Email    string             `bson:"_email"`
	Password string             `bson:"_password"`
}

type RequestChange struct {
	ID             primitive.ObjectID `bson:"_id"`
	OldDate        time.Time          `bson:"_olddate"`
	NewDate        time.Time          `bson:"_newdate"`
	Paid           string             `bson:"_paid"`
	Approved       string             `bson:"_approved"`
	UserID         string             `bson:"_requestedfor"`
	CreationDate   int64              `bson:"_creationdate"`
	CreationUserID string             `bson:"_creationuserid"`
	UpdateDate     int64              `bson:"_updatedate"`
	UpdateUserID   string             `bson:"_updateuserid"`
}

type Schema struct {
	ID             primitive.ObjectID `bson:"_id"`
	WeekDay        string             `bson:"_weekday"`
	CreationDate   int64              `bson:"_creationdate"`
	CreationUserID string             `bson:"_creationuserID"`
	UpdateDate     int64              `bson:"_updatedate"`
	UpdateUserID   string             `bson:"_updateuserid"`
}
