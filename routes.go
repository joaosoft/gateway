package gateway

import (
	"github.com/joaosoft/manager"
	"github.com/joaosoft/web"
	"github.com/joaosoft/acl"
)

func (c *Controller) RegisterRoutes(w manager.IWeb) error {
	return w.AddRoutes(

		// gateway
		manager.NewRoute(string(web.MethodGet), "/api/v1/gateway/alive", c.Alive),

		// auth
		manager.NewRoute(string(web.MethodPost), "/api/v1/auth/p/sign-up", c.RedirectAuth),
		manager.NewRoute(string(web.MethodGet), "/api/v1/auth/p/get-session", c.RedirectAuth),
		manager.NewRoute(string(web.MethodPut), "/api/v1/auth/p/refresh-session", c.RedirectAuth),

		manager.NewRoute(string(web.MethodPut), "/api/v1/auth/users/:id_user/deactivate", c.RedirectAuth),

		// profile
		manager.NewRoute(string(web.MethodGet), "/api/v1/profile/sections", c.RedirectProfile),
		manager.NewRoute(string(web.MethodGet), "/api/v1/profile/sections/contents", c.RedirectProfile),
		manager.NewRoute(string(web.MethodGet), "/api/v1/profile/sections/:section_key", c.RedirectProfile),
		manager.NewRoute(string(web.MethodGet), "/api/v1/profile/sections/:section_key/contents", c.RedirectProfile),
	)

	w.AddFilter("*", string(web.PositionBefore), acl.MiddlewareAcl(), string(web.MethodGet))

}
