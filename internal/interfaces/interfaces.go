package interfaces

import "github.com/nilsonmart/we-exchange/internal/models"

type ActivityRepository interface {
	Migrate() error
	Create(activity models.Activity) (*models.Activity, error)
	All() ([]models.Activity, error)
	GetByID(id int) *models.Activity
	GetByName(name string) (*models.Activity, error)
	Update(id int64, updated models.Activity) (*models.Activity, error)
	Delete(id int64) error
}

type SchemaRepository interface {
	Migrate() error
	Create(activity models.Schema) (*models.Schema, error)
	All() ([]models.Schema, error)
	GetByID(id int) *models.Schema
	GetByName(name string) (*models.Schema, error)
	Update(id int64, updated models.Schema) (*models.Schema, error)
	Delete(id int64) error
}
