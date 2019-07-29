package gateway

import (
	"github.com/joaosoft/web"
)

type Controller struct {
	config    *GatewayConfig
	webClient *web.Client
}

func NewController(config *GatewayConfig) (*Controller, error) {

	webClient, err := web.NewClient()
	if err != nil {
		return nil, err
	}

	return &Controller{
		config:    config,
		webClient: webClient,
	}, nil
}

func (c *Controller) Alive(ctx *web.Context) error {
	response := &AliveResponse{Message: "I'm alive!"}
	return ctx.Response.JSON(web.StatusOK, response)
}

func (c *Controller) RedirectAuth(ctx *web.Context) error {
	return ctx.Redirect(c.config.Services.Auth)
}
