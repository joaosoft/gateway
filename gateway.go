package gateway

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Gateway struct {
	config        *GatewayConfig
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	mux           sync.Mutex
}

// NewGateway ...
func NewGateway(options ...GatewayOption) (*Gateway, error) {
	config, simpleConfig, err := NewConfig()

	service := &Gateway{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("auth", logger.WarnLevel),
		config: config.Gateway,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(logger.Instance))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Gateway != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Gateway.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.Gateway = &GatewayConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	web := service.pm.NewSimpleWebServer(config.Gateway.Host)

	controller, err := NewController(config.Gateway)
	if err != nil {
		return nil, err
	}

	controller.RegisterRoutes(web)

	service.pm.AddWeb("api_web", web)

	return service, nil
}

// Start ...
func (gateway *Gateway) Start() error {
	return gateway.pm.Start()
}

// Stop ...
func (gateway *Gateway) Stop() error {
	return gateway.pm.Stop()
}
