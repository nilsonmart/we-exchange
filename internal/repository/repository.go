package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
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

// ACTIVITY
func (r *SQLiteRepository) MigrateActivity() error {
	query := `
    CREATE TABLE IF NOT EXISTS Activity(
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

func (r *SQLiteRepository) CreateActivity(activity models.Activity) (*models.Activity, error) {
	res, err := r.db.Exec("INSERT INTO Activity(OldDate, NewDate, Paid, Approved, UserID, CreationDate, CreationUserID, UpdateDate, UpdateUserID) values(?,?,?,?,?,?,?,?,?)",
		activity.OldDate, activity.NewDate, activity.Paid, activity.Approved, getUserId(), time.Now(), getUserId(), nil, nil)
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
	activity.ID = id

	return &activity, nil
}

func (r *SQLiteRepository) AllActivity() ([]models.Activity, error) {
	rows, err := r.db.Query("SELECT * FROM Activity")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []models.Activity
	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(&activity.ID, &activity.OldDate, &activity.NewDate, &activity.Paid, &activity.Approved, &activity.UserID, &activity.CreationDate, &activity.CreationUserID, &activity.UpdateDate, &activity.UpdateUserID); err != nil {
			return nil, err
		}
		all = append(all, activity)
	}
	return all, nil
}

func (r *SQLiteRepository) GetActivityByID(id int64) (*models.Activity, error) {
	row := r.db.QueryRow("SELECT * FROM Activity WHERE ID = ?", id)

	var activity models.Activity
	if err := row.Scan(&activity.ID, &activity.OldDate, &activity.NewDate, &activity.Paid, &activity.Approved, &activity.UserID, &activity.CreationDate, &activity.CreationUserID, &activity.UpdateDate, &activity.UpdateUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &activity, nil
}

func (r *SQLiteRepository) GetActivityByUserID(userId int64) (*models.Activity, error) {
	row := r.db.QueryRow("SELECT * FROM Activity WHERE UserID = ?", userId)

	var activity models.Activity
	if err := row.Scan(&activity.ID, &activity.OldDate, &activity.NewDate, &activity.Paid, &activity.Approved, &activity.UserID, &activity.CreationDate, &activity.CreationUserID, &activity.UpdateDate, &activity.UpdateUserID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExists
		}
		return nil, err
	}
	return &activity, nil
}

func (r *SQLiteRepository) UpdateActivity(id int64, updated models.Activity) (*models.Activity, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	res, err := r.db.Exec("UPDATE Activity SET OldDate = ?, NewDate = ?, Paid = ?, Approved = ?, UserID = ?, UpdateDate = ?, UpdateUserID = ? WHERE id = ?", updated.OldDate, updated.NewDate, updated.Paid, updated.Approved, updated.UserID, time.Now(), getUserId(), id)
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

func (r *SQLiteRepository) DeleteActivity(id int64) error {
	res, err := r.db.Exec("DELETE FROM Activity WHERE id = ?", id)
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
