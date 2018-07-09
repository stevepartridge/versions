package versions

import (
	"fmt"
	"testing"
	"time"
)

var (
	now = time.Now().UTC()

	testVersion1 = Version{
		Id:            1,
		Major:         "1",
		Minor:         "2",
		Revision:      "3",
		ApplicationId: "1234",
		Name:          "Test Version 1 Name",
		Stable:        true,
		UpdatedAt:     now,
		CreatedAt:     now,
	}

	testVersion2 = Version{
		Id:            2,
		Major:         "1",
		Minor:         "2",
		Revision:      "4",
		ApplicationId: "1234",
		Name:          "Test Version 2 Name",
		Stable:        false,
		UpdatedAt:     now,
		CreatedAt:     now,
	}
)

func Test_Unit_ValidateFail_MissingApplicationId(t *testing.T) {
	ver := testVersion1
	ver.ApplicationId = ""

	valid, err := ver.Validate()
	if valid {
		t.Error("Should not have been valid.")
	}
	if err != ErrVersionMissingApplicationId {
		t.Error(fmt.Sprintf("Error should have been: %s but was %s",
			ErrVersionMissingApplicationId,
			err.Error()),
		)
	}
}

func Test_Unit_ValidateFail_MissingComponents(t *testing.T) {
	ver := testVersion1
	ver.Major = ""
	ver.Minor = ""
	ver.Revision = ""

	valid, err := ver.Validate()
	if valid {
		t.Error("Should not have been valid.")
	}
	if err != ErrVersionMissingComponents {
		t.Error(fmt.Sprintf("Error should have been: %s but was %s",
			ErrVersionMissingApplicationId,
			err.Error()),
		)
	}
}

func Test_Unit_ValidateFail_RevisionEmpty(t *testing.T) {
	ver := testVersion1
	ver.Revision = ""

	valid, err := ver.Validate()
	if !valid {
		t.Error("Should have been valid.")
	}

	if err != nil {
		t.Error("Error should be nil")
	}

	if ver.Revision != "0" {
		t.Error("Revision should have been 0")
	}
}

func Test_Unit_ValidateFail_NameEmpty(t *testing.T) {
	ver := testVersion1
	ver.Name = ""

	valid, err := ver.Validate()
	if !valid {
		t.Error("Should have been valid.")
	}

	if err != nil {
		t.Error("Error should be nil")
	}

	if ver.Name != ver.Full() {
		t.Error("Name should have been ", ver.Full())
	}
}
