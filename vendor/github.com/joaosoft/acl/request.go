package acl

import "github.com/joaosoft/web"

type ErrorResponse struct {
	Code    web.Status `json:"code,omitempty"`
	Message string     `json:"message,omitempty"`
	Cause   string     `json:"cause,omitempty"`
}

type GetResourceCategoriesRequest struct {
	DomainKey string `json:"domain_key" validate:"notzero"`
}

type GetResourceCategoryPagesRequest struct {
	DomainKey           string `json:"domain_key" validate:"notzero"`
	ResourceCategoryKey string `json:"resource_category_key" validate:"notzero"`
}

type GetResourceCategoryPageRequest struct {
	DomainKey           string `json:"domain_key" validate:"notzero"`
	ResourceCategoryKey string `json:"resource_category_key" validate:"notzero"`
	ResourcePageKey     string `json:"resource_page_key" validate:"notzero"`
}

type GetPageResourcesRequest struct {
	UrlParams struct {
		DomainKey           string `json:"domain_key" validate:"notzero"`
		RoleKey             string `json:"role_key" validate:"notzero"`
		ResourceCategoryKey string `json:"resource_category_key" validate:"notzero"`
		ResourcePageKey     string `json:"resource_page_key" validate:"notzero"`
	}
	Params struct {
		User string `json:"user" validate:"notzero"`
	}
}

type GetPageResourcesByTypeRequest struct {
	UrlParams struct {
		DomainKey           string `json:"domain_key" validate:"notzero"`
		RoleKey             string `json:"role_key" validate:"notzero"`
		ResourceCategoryKey string `json:"resource_category_key" validate:"notzero"`
		ResourcePageKey     string `json:"resource_page_key" validate:"notzero"`
		ResourceTypeKey     string `json:"resource_type_key" validate:"notzero"`
	}
	Params struct {
		User string `json:"user" validate:"notzero"`
	}
}

type CheckEndpointAccessRequest struct {
	UrlParams struct {
		DomainKey       string `json:"domain_key" validate:"notzero"`
		RoleKey         string `json:"role_key" validate:"notzero"`
		ResourceTypeKey string `json:"resource_type_key" validate:"notzero"`
	}
	Params struct {
		Method   string `json:"method" validate:"notzero"`
		Endpoint string `json:"endpoint" validate:"notzero"`
		User     string `json:"user" validate:"notzero"`
	}
}

type CheckAclMiddleware struct {
	Method   string `json:"method" validate:"notzero"`
	Endpoint string `json:"endpoint" validate:"notzero"`
	Params   struct {
		DomainKey       string `json:"domain_key" validate:"notzero"`
		RoleKey         string `json:"role_key" validate:"notzero"`
		ResourceTypeKey string `json:"resource_type_key" validate:"notzero"`
		User            string `json:"user" validate:"notzero"`
	}
}

type CheckEndpointAccessResponse struct {
	IsAllowed bool `json:"is_allowed"`
}
