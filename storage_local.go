package versions

import (
	"fmt"
	"io/ioutil"
	"os"
	// "strconv"
	"github.com/stevepartridge/go/file"
)

var (
	localStorageDirectory = "storage"
)

type localStorage struct {
	directoryPath string
}

func NewLocalStorage(storagePath ...string) (*localStorage, error) {

	storage := localStorage{
		directoryPath: localStorageDirectory,
	}

	if len(storagePath) > 0 {
		storage.directoryPath = storagePath[0]
	}

	if !file.Exists(storage.directoryPath) {
		os.MkdirAll(storage.directoryPath, 0766)
	}

	return &storage, nil
}

func (s *localStorage) Save(versionId int, download *Download) error {

	if ok, err := download.Validate(); !ok {
		return err
	}

	if download.Id == 0 {
		return ErrDownloadIdMissing
	}

	return ioutil.WriteFile(
		fmt.Sprintf("%s/%s", s.directoryPath, download.SHA1),
		download.Data,
		0666,
	)

}

func (s localStorage) Get(download *Download) ([]byte, error) {

	filename := fmt.Sprintf("%s/%s", s.directoryPath, download.SHA1)

	if !file.Exists(filename) {
		return nil, ErrDownloadFileNotFound
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}
