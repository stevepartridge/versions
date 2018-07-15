package versions

const (
	StorageTypeLocal  = "local"
	StorageTypeGithub = "github"
	StorageTypeS3     = "s3"
)

var storages = map[string]Storage{}

type Storage interface {
	Save(int, *Download) error
	Get(*Download) ([]byte, error)
}

func AddStorage(storageType string, storage Storage) error {

	switch storageType {
	case StorageTypeLocal:
		storages[StorageTypeLocal] = storage
		return nil
	}
	return ErrStorageInvalidType
}

func GetStorage(storageType string) (Storage, error) {
	if storages[storageType] == nil {
		return nil, ErrStorageInvalidType
	}

	return storages[storageType], nil
}
