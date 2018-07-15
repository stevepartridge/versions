package versions

import (
	"net/http"
	"time"
)

type Download struct {
	Id          int    `json:"id"`
	VersionId   int    `json:"version_id"`
	StorageType string `json:"storage_type"`
	ContentType string `json:"content_type"`
	Filename    string `json:"filename"`
	Ext         string `json:"ext"`
	Format      string `json:"format"`
	Protocol    string `json:"protocol"`
	OS          string `json:"os"`
	Arch        string `json:"arch"`

	SHA1   string `json:"sha1"`
	SHA256 string `json:"sha256"`

	Downloads int `json:"total_downloads"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	Data []byte `json:"-"`
}

func (d *Download) Validate() (bool, error) {

	if d.StorageType == "" {
		return false, ErrDownloadStorageTypeMissing
	}

	if d.Filename == "" {
		return false, ErrDownloadFilenameMissing
	}

	if d.Data == nil {
		return false, ErrDownloadDataIsNil
	}

	if d.ContentType == "" {
		d.ContentType = http.DetectContentType(d.Data)
	}

	if d.SHA1 == "" {
		d.SHA1 = Sha1ToString(d.Data)
	}

	if d.SHA256 == "" {
		d.SHA256 = Sha256ToString(d.Data)
	}

	return true, nil
}

func (d *Download) GetFile() ([]byte, error) {

	storage, err := GetStorage(d.StorageType)
	if err != nil {
		return nil, err
	}

	data, err := storage.Get(d)
	if err != nil {
		return nil, err
	}

	return data, nil
}
