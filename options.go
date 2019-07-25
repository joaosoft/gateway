package auth

import (
	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// AuthOption ...
type AuthOption func(auth *Auth)

// Reconfigure ...
func (auth *Auth) Reconfigure(options ...AuthOption) {
	for _, option := range options {
		option(auth)
	}
}

// WithConfiguration ...
func WithConfiguration(config *AuthConfig) AuthOption {
	return func(auth *Auth) {
		auth.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) AuthOption {
	return func(auth *Auth) {
		log = logger
		auth.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) AuthOption {
	return func(auth *Auth) {
		log.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) AuthOption {
	return func(auth *Auth) {
		auth.pm = mgr
	}
}
