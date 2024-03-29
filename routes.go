package gateway

import (
	"github.com/joaosoft/manager"
	"github.com/joaosoft/web"
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

		// acl
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/roles/:role_key/pages/:resource_page_key/resources", c.RedirectAcl),
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/roles/:role_key/resources/types/:resource_type_key", c.RedirectAcl),
	)
}
