// Package common provides shared types and utilities used across the tuner1 application.
package common

type ExitError struct {
	Code    int
	Message string
}

func (ee ExitError) Error() string {
	return ee.Message
}

func (ee ExitError) String() string {
	return ee.Message
}
