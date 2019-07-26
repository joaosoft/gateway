package auth

import (
	"net/http"

	"github.com/joaosoft/manager"
)

func (c *Controller) RegisterRoutes(web manager.IWeb) error {
	return web.AddRoutes(

		// public routes
		manager.NewRoute(http.MethodPost, "/api/v1/auth/p/sign-up", c.SignUpHandler),
		manager.NewRoute(http.MethodGet, "/api/v1/auth/p/get-session", c.GetSessionHandler),
		manager.NewRoute(http.MethodPut, "/api/v1/auth/p/refresh-session", c.RefreshSessionHandler),

		// internal routes
		manager.NewRoute(http.MethodPut, "/api/v1/auth/users/:id_user/deactivate", c.DeactivateUserHandler),
	)
}
