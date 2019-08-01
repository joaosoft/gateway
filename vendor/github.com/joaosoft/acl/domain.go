package acl

import (
	"time"
)

type Categories []*Category

type Category struct {
	Name                      string    `json:"name" db:"name"`
	Key                       string    `json:"key" db:"key"`
	Description               string    `json:"description" db:"description"`
	ParentResourceCategoryKey *string   `json:"parent_resource_category_key,omitempty" db:"parent_resource_category_key"`
	Active                    bool      `json:"active" db:"active"`
	CreatedAt                 time.Time `json:"created_at" db:"created_at"`
	UpdatedAt                 time.Time `json:"updated_at" db:"updated_at"`
}

type Pages []*Page

type Page struct {
	Name                  string    `json:"name" db:"name"`
	Key                   string    `json:"key" db:"key"`
	Description           string    `json:"description" db:"description"`
	ParentResourcePageKey *string   `json:"parent_resource_page_key,omitempty" db:"parent_resource_page_key"`
	Active                bool      `json:"active" db:"active"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
}

type Resources []*Resource

type Resource struct {
	Name                string    `json:"name" db:"name"`
	Key                 string    `json:"key" db:"key"`
	ResourceCategoryKey string    `json:"resource_category_key" db:"resource_category_key"`
	ResourcePageKey     string    `json:"resource_page_key" db:"resource_page_key"`
	ResourceTypeKey     string    `json:"resource_type_key" db:"resource_type_key"`
	Description         string    `json:"description" db:"description"`
	Active              bool      `json:"active" db:"active"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`
}

type Allowed struct {
	Check bool `db:"check"`
	Allow bool `db:"allow"`
}
