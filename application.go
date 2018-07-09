package versions

import (
	"strings"
	"time"
	"unicode"
)

type Application struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Application) Validate() (bool, error) {

	if a.Id == "" {
		return false, ErrApplicationMissingId
	}

	if a.Name == "" {
		return false, ErrApplicationMissingName
	}

	for _, r := range a.Id {
		switch {
		case unicode.IsLetter(r), unicode.IsNumber(r):
			continue
		case strings.ContainsAny(string(r), "-_."):
			continue
		case unicode.IsSpace(r):
			return false, ErrApplicationIdInvalidCharacter("space")
		default:
			return false, ErrApplicationIdInvalidCharacter(string(r))
		}
	}

	return true, nil
}
