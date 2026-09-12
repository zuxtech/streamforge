package auth

import "errors"

var ErrPasswordBreached = errors.New(
	"password has been found in data breaches",
)