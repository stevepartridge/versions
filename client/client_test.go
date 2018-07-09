package client

import (
	"testing"
)

var (
	testServiceURI = "versions.local:8000"
)

func Test_Unit_NewClientCreate_Success(t *testing.T) {
	c, err := NewClient(testServiceURI)
	if err != nil {
		t.Errorf("Error should be nil but saw %s", err.Error())
	}

	if c.serviceURI != testServiceURI {
		t.Errorf("Service URI should be %s but saw %s",
			testServiceURI,
			c.serviceURI,
		)
	}
}
