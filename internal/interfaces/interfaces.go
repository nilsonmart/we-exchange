package interfaces

type ActivityRepository interface {
	Migrate() error
	Create(activity Activity) (*Website, error)
	All() ([]Website, error)
	GetByName(name string) (*Website, error)
	Update(id int64, updated Website) (*Website, error)
	Delete(id int64) error
}
