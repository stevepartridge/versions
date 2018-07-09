package versions

import (
	"fmt"
	"time"
)

type Version struct {
	Id       int    `json:"id"`
	Major    string `json:"major"`
	Minor    string `json:"minor"`
	Revision string `json:"revision"`

	ApplicationId string `json:"application_id"`
	Name          string `json:"name"`
	Stable        bool   `json:"stable"`

	Downloads []Download `json:"downloads"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (v *Version) Full() string {
	return fmt.Sprintf("%s.%s.%s", v.Major, v.Minor, v.Revision)
}

// Validate confirms the Version is sufficient for entry into the store
//  returns:
//    - valid/ok (bool)
//    - error (error)
func (v *Version) Validate() (bool, error) {

	if v.ApplicationId == "" {
		return false, ErrVersionMissingApplicationId
	}

	if v.Major == "" && v.Minor == "" && v.Revision == "" {
		return false, ErrVersionMissingComponents
	}

	if v.Revision == "" {
		v.Revision = "0"
	}

	if v.Name == "" {
		v.Name = v.Full()
	}

	return true, nil
}
