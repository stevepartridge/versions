package versions

type Store interface {
	Type() string
	Migrations() error

	Create(*Version) error
	Update(*Version) error
	Delete(id int) error
	GetById(id int) (Version, error)

	// Get returns the version(s) for an application ID.
	// options:
	//   opts[0] > 0 = limit
	//   opts[1] > 0 = offset
	GetByApplication(appId string, opts ...int) ([]Version, error)

	Exists(*Version) bool

	// GetLatest returns the latest version for an application ID.
	// options:
	//   opts[0] > 0 = stable only
	GetLatest(appId string, opts ...int) (Version, error)

	ApplicationExists(*Application) bool

	CreateApplication(*Application) error
	UpdateApplication(*Application) error
	DeleteApplication(id string) error
	GetApplicationById(id string) (Application, error)
	GetApplications(opts ...int) ([]Application, error)
}

func NewStore(storeType string) (Store, error) {
	switch storeType {
	// case "postgres", "pg":
	// case "mysql", "maria":
	// case "sqlite":
	case "file":
	}
	return newFileStore()

}
