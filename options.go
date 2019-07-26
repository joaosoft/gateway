package gateway

import (
	logger "github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// GatewayOption ...
type GatewayOption func(auth *Gateway)

// Reconfigure ...
func (gateway *Gateway) Reconfigure(options ...GatewayOption) {
	for _, option := range options {
		option(gateway)
	}
}

// WithConfiguration ...
func WithConfiguration(config *GatewayConfig) GatewayOption {
	return func(gateway *Gateway) {
		gateway.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) GatewayOption {
	return func(gateway *Gateway) {
		log = logger
		gateway.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) GatewayOption {
	return func(gateway *Gateway) {
		log.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) GatewayOption {
	return func(gateway *Gateway) {
		gateway.pm = mgr
	}
}
