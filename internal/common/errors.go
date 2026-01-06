package common

type ExitError struct {
	Code    int
	Message string
}

func (ee ExitError) Error() string {
	return ee.Message
}
