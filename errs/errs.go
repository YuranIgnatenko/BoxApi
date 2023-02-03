package errs

import "errors"

var (
	ErrorNotFoundHost = errors.New("Not Found Host in config file")
	ErrorNotFoundPort = errors.New("Not Found Port in config file")
)
