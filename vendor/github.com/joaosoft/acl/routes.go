package acl

import (
	"encoding/json"

	"github.com/joaosoft/manager"
	"github.com/joaosoft/validator"
	"github.com/joaosoft/web"
)

func (c *Controller) RegisterRoutes(w manager.IWeb) error {
	err := w.AddRoutes(
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/domains/:domain_key/categories", c.GetResourceCategoriesHandler),
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/domains/:domain_key/categories/:resource_category_key/pages", c.GetResourceCategoryPagesHandler),
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/domains/:domain_key/categories/:resource_category_key/pages/:resource_page_key", c.GetResourceCategoryPageHandler),
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/domains/:domain_key/roles/:role_key/categories/:resource_category_key/pages/:resource_page_key/resources", c.GetResourcesHandler),
		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/domains/:domain_key/roles/:role_key/categories/:resource_category_key/pages/:resource_page_key/resources/types/:resource_type_key", c.GetResourcesByTypeHandler),

		manager.NewRoute(string(web.MethodGet), "/api/v1/acl/domains/:domain_key/roles/:role_key/resources/types/:resource_type_key", c.CheckEndpointAccessHandler),

		manager.NewRoute(string(web.MethodGet), "/api/v1/dummy", c.DummyHandler),
	)

	if err != nil {
		return err
	}

	w.AddFilter("/api/v1/dummy", string(web.PositionBefore), c.MiddlewareAcl(), string(web.MethodGet))

	return nil
}

func (c *Controller) MiddlewareAcl() web.MiddlewareFunc {
	return func(next web.HandlerFunc) web.HandlerFunc {
		return func(ctx *web.Context) error {

			if err := c.checkAcl(ctx); err != nil {
				ctx.Response.Status = web.StatusForbidden
				return nil
			}

			return next(ctx)
		}
	}
}

func (c *Controller) checkAcl(ctx *web.Context) error {
	request := &CheckAclMiddleware{
		Method:   string(ctx.Request.Method),
		Endpoint: ctx.Request.Address.Url,
	}

	if err := ctx.Request.BindParams(&request.Params); err != nil {
		return err
	}

	if errs := validator.Validate(request); len(errs) > 0 {
		bytes, err := json.Marshal(errs)
		if err != nil {
			return err
		}

		return ErrorGeneric.Format(string(bytes))
	}

	isAllowed, err := c.interactor.CheckAcl(request)
	if err != nil {
		return err
	}

	if !isAllowed {
		return ErrorAclAccessDenied
	}

	return nil
}
