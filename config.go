package gateway

import (
	"fmt"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Gateway *GatewayConfig `json:"gateway"`
}

// GatewayConfig ...
type GatewayConfig struct {
	Host     string    `json:"host"`
	Services *Services `json:"services"`
	TokenKey string    `json:"token_key"`
	Log      struct {
		Level string `json:"level"`
	} `json:"log"`
}

type Services struct {
	Auth string `json:"auth"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}
