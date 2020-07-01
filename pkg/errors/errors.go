package errors

import "github.com/pkg/errors"

type HTTPError struct {
	statusCode int
	err        error
}

func (e *HTTPError) Error() string {
	return e.err.Error()
}

func (e *HTTPError) Code() int {
	return e.statusCode
}

func WrapMC(err error, t string, c int) *HTTPError {
	return &HTTPError{
		statusCode: c,
		err:        errors.Wrap(err, t),
	}
}

func WrapC(err error, c int) *HTTPError {
	return &HTTPError{
		statusCode: c,
		err:        err,
	}
}

func NewC(m string, c int) *HTTPError {
	return &HTTPError{
		statusCode: c,
		err:        errors.New(m),
	}
}
