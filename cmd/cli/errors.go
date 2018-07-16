package main

import (
	"errors"
	"fmt"
)

var (
	ErrDownloadCreateFileNotFound = errors.New("Unable to create download, file not found.")
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("err:", err.Error())
		return true
	}
	return false
}
