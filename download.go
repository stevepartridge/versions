package versions

import (
	"time"
)

type Download struct {
	Link     string `json:"link"`
	Format   string `json:"format"`
	Protocol string `json:"protocol"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (d *Download) Validate() (bool, error) {
	return true, nil
}
