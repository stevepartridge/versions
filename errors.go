package versions

import (
	"errors"
	"fmt"
)

var (
	ErrVersionNotFound    = errors.New("versions.not_found")
	ErrVersionExists      = errors.New("versions.exists")
	ErrVersionIsNil       = errors.New("versions.version.is_nil")
	ErrVersionMissingOS   = errors.New("versions.missing.os")
	ErrVersionMissingArch = errors.New("versions.missing.arch")

	ErrVersionInvalidId            = errors.New("versions.invalid.id")
	ErrVersionMissingId            = errors.New("versions.missing.id")
	ErrVersionIdNotEmpty           = errors.New("versions.invalid.id_not_empty")
	ErrVersionMissingApplication   = errors.New("versions.missing.application")
	ErrVersionMissingApplicationId = errors.New("versions.missing.application_id")
	ErrVersionMissingComponents    = errors.New("versions.missing.major_minor_and_revision")

	ErrApplicationExists      = errors.New("versions.application.exists")
	ErrApplicationMissingId   = errors.New("versions.application.missing.id")
	ErrApplicationMissingName = errors.New("versions.application.missing.name")
	ErrApplicationNotFound    = errors.New("versions.application.not_found")
	// ErrApplicationExists    = errors.New("versions.application.exists")

	// ErrApplicationIdContainsSpace = errors.New("versions.application.invalid_id.contains")
)

func ErrApplicationIdInvalidCharacter(char string) error {
	return errors.New(
		fmt.Sprintf("versions.application.invalid_id.contains: %s",
			char),
	)
}
