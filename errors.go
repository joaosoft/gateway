package auth

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorNotFound             = errors.New(errors.ErrorLevel, int(web.StatusNotFound), "user not found")
	ErrorInvalidBodyParameter = errors.New(errors.ErrorLevel, int(web.StatusBadRequest), "invalid body parameter '%s'")
)
