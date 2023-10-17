package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/nilsonmart/we-exchange/internal/driver"
	"github.com/nilsonmart/we-exchange/internal/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type SQLiteRepository struct {
	db *sql.DB
}

func queryObjectData(schema, key, value string) ([]byte, error) {
	query, err := driver.QueryObject(schema, &models.DBQuery{Key: key, Value: value})
	if err != nil {
		return nil, err
	}
	return query, err
}

func getUserId() int64 {
	var c *gin.Context
	userId, err := c.Cookie("userid")
	if err != nil {
		panic(err)
	}

	p, err := strconv.ParseInt(userId, 10, 0)
	if err != nil {
		return 0
	}

	return p

}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		db: db,
	}
}

// ACCOUNT
// func (r *SQLiteRepository) MigrateAccount() error {
// 	query := `
//     CREATE TABLE IF NOT EXISTS Account(
//         ID INTEGER PRIMARY KEY AUTOINCREMENT,
//         Name TEXT NOT NULL,
//         Email TEXT NOT NULL,
//         Password NUMERIC NOT NULL
//     );
//     `
// 	_, err := r.db.Exec(query)
// 	return err
// }

func (r *SQLiteRepository) CreateAccount(account models.Account) (*models.Account, error) {
	res, err := r.db.Exec("INSERT INTO Account(Name, Email, Password) values(?,?,?)",
		account.Name, account.Email, account.Password)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	account.ID = id

	return &account, nil
}

func (r *SQLiteRepository) AllAccount() ([]models.Account, error) {
	rows, err := r.db.Query("SELECT * FROM Account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.Account
	for rows.Next() {
		var accounts models.Account
		if err := rows.Scan(&accounts.ID, &accounts.Name, &accounts.Email, &accounts.Password); err != nil {
			return nil, err
		}
		all = append(all, accounts)
	}
	return all, nil
}

func (r *SQLiteRepository) GetAccountByID(userId string) (*models.Account, error) {

	docID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	fmt.Printf("HomePageHandler ObjectID - %v", docID)

	query, err := queryObjectData("schema", "_userid", userId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotExists
	}

	var schema models.Schema
	err = json.Unmarshal(query, &schema)
	if err != nil {
		//Log error
		fmt.Println("Error when unmashalling SchemaDB")
		return nil, err
	}

	row := r.db.QueryRow("SELECT * FROM Account WHERE ID = ?", id)

	var account models.Account
	if err := row.Scan(&account.ID, &account.Name, &account.Email, &account.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &account, nil
}

func (r *SQLiteRepository) GetAccountByEmail(email string) (*models.Account, error) {
	row := r.db.QueryRow("SELECT * FROM Account WHERE Email = ?", email)
	fmt.Println(row)

	var account models.Account
	if err := row.Scan(&account.ID, &account.Name, &account.Email, &account.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println(err)
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &account, nil
}

func (r *SQLiteRepository) ValidateAccount(email, password string) (bool, error) {
	if email == "" || password == "" {
		//TODO Log error
		log.Fatal("Email or Password invalid.")
		return false, errors.New("Email or Password invalid.")
	}

	row := r.db.QueryRow("SELECT * FROM Account WHERE Email = ? AND Password = ?", email, password)

	var account models.Account
	if err := row.Scan(&account.ID, &account.Name, &account.Email, &account.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//TODO Log error
			log.Fatal(ErrNotExists)
			return false, ErrNotExists
		}

		//TODO Log error
		log.Fatal(err)

		return false, err
	}

	return true, ErrNotExists
}

// RequestChange
func (r *SQLiteRepository) MigrateRequestChange() error {
	query := `
    CREATE TABLE IF NOT EXISTS RequestChange(
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        OldDate TEXT NOT NULL,
        NewDate TEXT NOT NULL,
        Paid NUMERIC NOT NULL,
		Approved NUMERIC NOT NULL,
		UserID INTEGER NOT NULL,
		CreationDate TEXT NOT NULL,
		CreationUserID INTEGER NOT NULL,
		UpdateDate TEXT NOT NULL,
		UpdateUserID INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) CreateRequestChange(RequestChange models.RequestChange) (*models.RequestChange, error) {
	res, err := r.db.Exec("INSERT INTO RequestChange(OldDate, NewDate, Paid, Approved, UserID, CreationDate, CreationUserID, UpdateDate, UpdateUserID) values(?,?,?,?,?,?,?,?,?)",
		RequestChange.OldDate, RequestChange.NewDate, RequestChange.Paid, RequestChange.Approved, getUserId(), time.Now(), getUserId(), nil, nil)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	RequestChange.ID = id

	return &RequestChange, nil
}

func (r *SQLiteRepository) AllRequestChange() ([]models.RequestChange, error) {
	rows, err := r.db.Query("SELECT * FROM RequestChange")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.RequestChange
	for rows.Next() {
		var RequestChange models.RequestChange
		if err := rows.Scan(&RequestChange.ID, &RequestChange.OldDate, &RequestChange.NewDate, &RequestChange.Paid, &RequestChange.Approved, &RequestChange.UserID, &RequestChange.CreationDate, &RequestChange.CreationUserID, &RequestChange.UpdateDate, &RequestChange.UpdateUserID); err != nil {
			return nil, err
		}
		all = append(all, RequestChange)
	}
	return all, nil
}

func (r *SQLiteRepository) GetRequestChangeByID(id int64) (*models.RequestChange, error) {
	row := r.db.QueryRow("SELECT * FROM RequestChange WHERE ID = ?", id)

	var RequestChange models.RequestChange
	if err := row.Scan(&RequestChange.ID, &RequestChange.OldDate, &RequestChange.NewDate, &RequestChange.Paid, &RequestChange.Approved, &RequestChange.UserID, &RequestChange.CreationDate, &RequestChange.CreationUserID, &RequestChange.UpdateDate, &RequestChange.UpdateUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &RequestChange, nil
}

func (r *SQLiteRepository) GetRequestChangeByUserID(userId int64) (*models.RequestChange, error) {
	row := r.db.QueryRow("SELECT * FROM RequestChange WHERE UserID = ?", userId)

	var RequestChange models.RequestChange
	if err := row.Scan(&RequestChange.ID, &RequestChange.OldDate, &RequestChange.NewDate, &RequestChange.Paid, &RequestChange.Approved, &RequestChange.UserID, &RequestChange.CreationDate, &RequestChange.CreationUserID, &RequestChange.UpdateDate, &RequestChange.UpdateUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &RequestChange, nil
}

func (r *SQLiteRepository) UpdateRequestChange(id int64, updated models.RequestChange) (*models.RequestChange, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE RequestChange SET OldDate = ?, NewDate = ?, Paid = ?, Approved = ?, UserID = ?, UpdateDate = ?, UpdateUserID = ? WHERE id = ?", updated.OldDate, updated.NewDate, updated.Paid, updated.Approved, updated.UserID, time.Now(), getUserId(), id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) DeleteRequestChange(id int64) error {
	res, err := r.db.Exec("DELETE FROM RequestChange WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}

// SCHEMA
func (r *SQLiteRepository) MigrateSchema() error {
	query := `
    CREATE TABLE IF NOT EXISTS Schema(
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        WeekDay TEXT NOT NULL,
		UserID INTEGER NOT NULL,
		CreationDate INTEGER NOT NULL,
		CreationUserID INTEGER NOT NULL,
		UpdateDate INTEGER NOT NULL,
		UpdateUserID INTEGER NOT NULL
    );
    `

	_, err := r.db.Exec(query)
	return err
}

func (r *SQLiteRepository) CreateSchema(schema models.Schema) (*models.Schema, error) {
	res, err := r.db.Exec("INSERT INTO Schema(WeekDay, UserID, CreationDate, CreationUserID, UpdateDate, UpdateUserID) values(?,?,?,?,?,?)",
		schema.WeekDay, getUserId(), schema.CreationDate, getUserId(), time.Now(), nil, nil)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	schema.ID = id

	return &schema, nil
}

func (r *SQLiteRepository) AllSchema() ([]models.Schema, error) {
	rows, err := r.db.Query("SELECT * FROM Schema")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.Schema
	for rows.Next() {
		var schema models.Schema
		if err := rows.Scan(&schema.ID, &schema.WeekDay, &schema.UserID, &schema.CreationDate, &schema.CreationUserID, &schema.UpdateDate, &schema.UpdateUserID); err != nil {
			return nil, err
		}
		all = append(all, schema)
	}
	return all, nil
}

func (r *SQLiteRepository) GetSchemaByID(id int64) (*models.Schema, error) {
	row := r.db.QueryRow("SELECT * FROM Schema WHERE ID = ?", id)

	var schema models.Schema
	if err := row.Scan(&schema.ID, &schema.WeekDay, &schema.UserID, &schema.CreationDate, &schema.CreationUserID, &schema.UpdateDate, &schema.UpdateUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &schema, nil
}

func (r *SQLiteRepository) GetSchemaByUserID(userId int64) (*models.Schema, error) {
	row := r.db.QueryRow("SELECT * FROM Schema WHERE UserID = ?", userId)

	var schema models.Schema
	if err := row.Scan(&schema.ID, &schema.WeekDay, &schema.UserID, &schema.CreationDate, &schema.CreationUserID, &schema.UpdateDate, &schema.UpdateUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &schema, nil
}

func (r *SQLiteRepository) UpdateSchema(id int64, updated models.Schema) (*models.Schema, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE Schema SET WeekDay = ?, UserID = ?, UpdateDate = ?, UpdateUserID = ? WHERE id = ?", updated.WeekDay, updated.UserID, time.Now(), getUserId(), id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updated, nil
}

func (r *SQLiteRepository) DeleteSchema(id int64) error {
	res, err := r.db.Exec("DELETE FROM Schema WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}
