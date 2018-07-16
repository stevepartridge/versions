package versions

import (
	"testing"
)

func Test_Unit_AddStorage_Success(t *testing.T) {
	storage, err := NewLocalStorage("testing-storage")
	if err != nil {
		t.Errorf("Didn't expect error, but saw %s", err.Error())
	}

	err = AddStorage(StorageTypeLocal, storage)
	if err != nil {
		t.Errorf("Didn't expect error, but saw %s", err.Error())
	}

	_, err = GetStorage(StorageTypeLocal)
	if err != nil {
		t.Errorf("Didn't expect error, but saw %s", err.Error())
	}

}

func Test_Unit_AddStorage_Fail(t *testing.T) {
	storage, err := NewLocalStorage("testing-storage")
	if err != nil {
		t.Errorf("Didn't expect error, but saw %s", err.Error())
	}

	err = AddStorage("bogus", storage)
	if err == nil {
		t.Error("Expected error, but saw nil")
	}

	_, err = GetStorage("bogus")
	if err == nil {
		t.Error("Expected error, but saw nil")
	}

}
