package interfaces

type ActivityRepository interface {
	Migrate() error
	Create(activity Activity) (*Activity, error)
	All() ([]Activity, error)
	GetByName(name string) (*Activity, error)
	Update(id int64, updated Activity) (*Activity, error)
	Delete(id int64) error
}
