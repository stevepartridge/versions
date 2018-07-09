package versions

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"time"

	// "github.com/rs/zerolog/log"

	"github.com/stevepartridge/go/file"
)

var (
	fileStoreFilename = ".versions"
	localDb           fileDB
)

type fileDB struct {
	Version string `json:"version"`

	Versions     []Version     `json:"versions,omitempty"`
	Applications []Application `json:"applications,omitempty"`
	Downloads    []Download    `json:"downloads,omitempty"`
}

type fileStore struct{}

func newFileStore() (*fileStore, error) {

	localDb = fileDB{}
	localDb.init()

	if localDb.Version == "" {
		localDb.Version = "0.1"
		localDb.save()
	}

	store := fileStore{}

	return &store, nil
}

func (s fileStore) Type() string {
	return "file"
}

func (s fileStore) Migrations() error {
	if !file.Exists(fileStoreFilename) {
		return localDb.init()
	}
	return nil
}

func (s fileStore) Create(version *Version) error {

	if ok, err := version.Validate(); !ok {
		return err
	}

	if version.Id > 0 {
		return ErrVersionIdNotEmpty
	}

	_, err := s.GetApplicationById(version.ApplicationId)
	if err != nil {
		return err
	}

	if s.Exists(version) {
		return ErrVersionExists
	}

	version.Id = localDb.getNextId()

	version.CreatedAt = time.Now().UTC()
	version.UpdatedAt = time.Now().UTC()

	localDb.Versions = append(localDb.Versions, *version)

	return localDb.save()

}

func (s fileStore) Update(version *Version) error {

	if ok, err := version.Validate(); !ok {
		return err
	}

	if version.Id == 0 {
		return ErrVersionMissingId
	}

	v, err := s.GetById(version.Id)
	if err != nil {
		return err
	}

	v.Name = version.Name
	v.ApplicationId = version.ApplicationId
	v.Major = version.Major
	v.Minor = version.Minor
	v.Revision = version.Revision

	v.UpdatedAt = time.Now().UTC()

	for i := range localDb.Versions {
		if localDb.Versions[i].Id == v.Id {
			localDb.Versions[i] = v
			break
		}
	}

	*version = v

	return localDb.save()

}

func (s fileStore) GetById(id int) (Version, error) {

	for i := range localDb.Versions {

		if localDb.Versions[i].Id == id {
			return localDb.Versions[i], nil
		}
	}

	return Version{}, ErrVersionNotFound
}

func (s fileStore) Delete(id int) error {

	for i := range localDb.Versions {

		if localDb.Versions[i].Id == id {

			// if i == len(localDb.Versions)-1 {

			//      copy(localDb.Versions[i:], localDb.Versions[i+1:])
			// 	localDb.Versions[len(localDb.Versions)-1] = Version{}
			// 	localDb.Versions = localDb.Versions[:len(localDb.Versions)-1]

			// 	continue

			// }

			copy(localDb.Versions[i:], localDb.Versions[i+1:])
			localDb.Versions[len(localDb.Versions)-1] = Version{}
			localDb.Versions = localDb.Versions[:len(localDb.Versions)-1]

			localDb.save()

			return nil
		}
	}

	return ErrVersionNotFound
}

func (s fileStore) GetByApplication(appId string, opts ...int) ([]Version, error) {

	limit := 0

	if len(opts) > 0 {
		limit = opts[0]
	}

	result := make([]Version, 0)

	for i := range localDb.Versions {
		if localDb.Versions[i].ApplicationId == appId {
			result = append(result, localDb.Versions[i])
		}
	}

	if limit > 0 {
		return result[:limit], nil
	}

	return result, nil
}

func (s fileStore) GetLatest(appId string, opts ...int) (Version, error) {

	stable := false

	if len(opts) > 0 {
		if opts[0] > 0 {
			stable = true
		}
	}

	latest := Version{
		CreatedAt: time.Time{},
	}

	versions, err := s.GetByApplication(appId)
	if err != nil {
		return latest, err
	}

	for i := range versions {

		if versions[i].CreatedAt.After(latest.CreatedAt) {

			if stable && !versions[i].Stable {
				continue
			}

			latest = versions[i]
		}
	}

	return latest, nil
}

func (s fileStore) Exists(version *Version) bool {

	for i := range localDb.Versions {

		if localDb.Versions[i].ApplicationId == version.ApplicationId {

			if localDb.Versions[i].Full() == version.Full() {
				return true
			}
		}
	}
	return false
}

func (s fileStore) ApplicationExists(application *Application) bool {

	for i := range localDb.Applications {

		if localDb.Applications[i].Id == application.Id {

			return true

		}
	}
	return false
}

func (s fileStore) CreateApplication(application *Application) error {

	if ok, err := application.Validate(); !ok {
		return err
	}

	if s.ApplicationExists(application) {
		return ErrApplicationExists
	}

	application.CreatedAt = time.Now().UTC()
	application.UpdatedAt = time.Now().UTC()

	localDb.Applications = append(localDb.Applications, *application)

	return localDb.save()

}

func (s fileStore) GetApplicationById(id string) (Application, error) {

	for i := range localDb.Applications {

		if localDb.Applications[i].Id == id {
			return localDb.Applications[i], nil
		}
	}

	return Application{}, ErrApplicationNotFound
}

func (s fileStore) UpdateApplication(application *Application) error {

	if ok, err := application.Validate(); !ok {
		return err
	}

	if !s.ApplicationExists(application) {
		return ErrApplicationNotFound
	}

	application.UpdatedAt = time.Now().UTC()

	localDb.Applications = append(localDb.Applications, *application)

	return localDb.save()

}

func (s fileStore) DeleteApplication(id string) error {

	for i := range localDb.Applications {

		if localDb.Applications[i].Id == id {

			copy(localDb.Applications[i:], localDb.Applications[i+1:])
			localDb.Applications[len(localDb.Applications)-1] = Application{}
			localDb.Applications = localDb.Applications[:len(localDb.Applications)-1]

			localDb.save()

			return nil
		}
	}

	return ErrApplicationNotFound
}

func (s fileStore) GetApplications(opts ...int) ([]Application, error) {

	limit := 0

	if len(opts) > 0 {
		limit = opts[0]
	}

	if limit > 0 {
		return localDb.Applications[:limit], nil
	}

	return localDb.Applications, nil
}

func (db *fileDB) getNextId() int {

	id := 1

	for i := range localDb.Versions {

		if localDb.Versions[i].Id >= id {

			id = localDb.Versions[i].Id + 1
		}
	}

	return id

}

func (db *fileDB) load() error {

	if !file.Exists(fileStoreFilename) {
		db.save()
	}

	data, err := ioutil.ReadFile(fileStoreFilename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &db)
	if err != nil {
		return err
	}

	return nil
}

func (db *fileDB) save() error {

	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileStoreFilename, data, 0655)

}

func (db *fileDB) init() error {
	return db.load()
}
