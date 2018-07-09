package main

import (
	"fmt"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println("err:", err.Error())
		return true
	}
	return false
}
