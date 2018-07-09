package versions

import (
	// "io/ioutil"
	"fmt"
	"os"
	"testing"

	"github.com/stevepartridge/go/file"
)

var (
	store *fileStore
)

func TestMain(m *testing.M) {

	fileStoreFilename = ".versions-tests"

	if file.Exists(fileStoreFilename) {
		err := os.Remove(fileStoreFilename)
		if err != nil {
			fmt.Errorf("Error removing file %s", fileStoreFilename)
			os.Exit(1)
		}
	}

	result := m.Run()

	teardown()

	os.Exit(result)
}

func teardown() {
	if file.Exists(fileStoreFilename) {
		err := os.Remove(fileStoreFilename)
		if err != nil {
			fmt.Errorf("Error removing file %s", fileStoreFilename)
			os.Exit(1)
		}
	}
}

func makeNewFileStore(t *testing.T) {

	removeFileStore(t)

	var err error
	store, err = newFileStore()
	if err != nil {
		t.Errorf("Error creating new file store should be nil. %s",
			fileStoreFilename)
	}

}

func removeFileStore(t *testing.T) {
	if file.Exists(fileStoreFilename) {
		err := os.Remove(fileStoreFilename)
		if err != nil {
			t.Errorf("Error removing file %s", fileStoreFilename)
		}
	}
}
