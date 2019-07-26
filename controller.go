package gateway

import (
	"fmt"

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
	return c.redirect(ctx, c.config.Services.Auth)
}

func (c *Controller) redirect(ctx *web.Context, host string) error {
	ctx.Request.Client = c.webClient
	ctx.Request.Address = web.NewAddress(fmt.Sprintf("%s%s", host, ctx.Request.Address.Url))

	response, err := ctx.Request.Send()
	if err != nil {
		return err
	}

	return ctx.Response.Bytes(response.Status, response.ContentType, response.Body)
}
