package util

import (
	. "fmt"
)

//Errors

type GoLinkError struct {
	cause error
	message string
}

//Returns nested error
func (err GoLinkError) Cause() error {
	return err.cause
}

//Golang stdlib func
func (err GoLinkError) Error() string {
	return err.message
}

//Sugar
func ExitGently(a ...interface{}) {
	panic(GoLinkError{message: Sprintln(a...)})
}
