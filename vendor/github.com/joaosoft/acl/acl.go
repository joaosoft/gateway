package acl

import (
	"sync"

	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
	migration "github.com/joaosoft/migration/services"
)

type Acl struct {
	config        *AclConfig
	isLogExternal bool
	pm            *manager.Manager
	logger        logger.ILogger
	mux           sync.Mutex
}

// NewAcl ...
func NewAcl(options ...AclOption) (*Acl, error) {
	config, simpleConfig, err := NewConfig()

	service := &Acl{
		pm:     manager.NewManager(manager.WithRunInBackground(false)),
		logger: logger.NewLogDefault("acl", logger.WarnLevel),
		config: config.Acl,
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(logger.Instance))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Acl != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Acl.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	} else {
		config.Acl = &AclConfig{
			Host: defaultURL,
		}
	}

	service.Reconfigure(options...)

	// execute migrations
	migrationService, err := migration.NewCmdService(migration.WithCmdConfiguration(service.config.Migration))
	if err != nil {
		return nil, err
	}

	if _, err := migrationService.Execute(migration.OptionUp, 0, migration.ExecutorModeDatabase); err != nil {
		return nil, err
	}

	web := service.pm.NewSimpleWebServer(config.Acl.Host)

	storage, err := NewStoragePostgres(config.Acl)
	if err != nil {
		return nil, err
	}

	interactor := NewInteractor(config.Acl, storage)

	controller := NewController(config.Acl, interactor)
	controller.RegisterRoutes(web)

	service.pm.AddWeb("api_web_acl", web)

	return service, nil
}

// Start ...
func (m *Acl) Start() error {
	return m.pm.Start()
}

// Stop ...
func (m *Acl) Stop() error {
	return m.pm.Stop()
}
