package versions

import (
	"testing"
)

var (
	testDownload1 = Download{
		Id:        1,
		VersionId: 1,
		Filename:  "testdownloadfilename",
		UpdatedAt: now,
		CreatedAt: now,
	}
)

func Test_Unit_DownloadValidate_Invalid(t *testing.T) {

	dl := Download{}

	_, err := dl.Validate()
	if err != ErrDownloadStorageTypeMissing {
		t.Errorf("Expected %s but saw %s",
			ErrDownloadStorageTypeMissing.Error(),
			err.Error())
	}

	dl.StorageType = "local"

	_, err = dl.Validate()
	if err != ErrDownloadFilenameMissing {
		t.Errorf("Expected %s but saw %s",
			ErrDownloadFilenameMissing.Error(),
			err.Error())
	}

	dl.Filename = "filename"

	_, err = dl.Validate()
	if err != ErrDownloadDataIsNil {
		t.Errorf("Expected %s but saw %s",
			ErrDownloadDataIsNil.Error(),
			err.Error())
	}

	dl.Data = []byte("1234")

	_, err = dl.Validate()
	if err != nil {
		t.Errorf("Expected to be nil but saw %s",
			err.Error())
	}

}
