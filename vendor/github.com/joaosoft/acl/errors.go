package acl

import (
	"github.com/joaosoft/errors"
	"github.com/joaosoft/web"
)

var (
	ErrorGeneric         = errors.New(errors.ErrorLevel, int(web.StatusForbidden), "%s")
	ErrorAclAccessDenied = errors.New(errors.ErrorLevel, int(web.StatusForbidden), "acl access denied")
)
