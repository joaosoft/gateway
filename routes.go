package gateway

import (
	"net/http"

	"github.com/joaosoft/manager"
)

func (c *Controller) RegisterRoutes(web manager.IWeb) error {
	return web.AddRoutes(

		// gateway
		manager.NewRoute(http.MethodGet, "/api/v1/gateway/alive", c.Alive),

		// auth
		manager.NewRoute(http.MethodPost, "/api/v1/auth/p/sign-up", c.RedirectAuth),
		manager.NewRoute(http.MethodGet, "/api/v1/auth/p/get-session", c.RedirectAuth),
		manager.NewRoute(http.MethodPut, "/api/v1/auth/p/refresh-session", c.RedirectAuth),

		manager.NewRoute(http.MethodPut, "/api/v1/auth/users/:id_user/deactivate", c.RedirectAuth),
	)
}
