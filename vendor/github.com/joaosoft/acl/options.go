package acl

import (
	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// AclOption ...
type AclOption func(client *Acl)

// Reconfigure ...
func (session *Acl) Reconfigure(options ...AclOption) {
	for _, option := range options {
		option(session)
	}
}

// WithConfiguration ...
func WithConfiguration(config *AclConfig) AclOption {
	return func(session *Acl) {
		session.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) AclOption {
	return func(session *Acl) {
		log = logger
		session.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) AclOption {
	return func(session *Acl) {
		log.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) AclOption {
	return func(session *Acl) {
		session.pm = mgr
	}
}
