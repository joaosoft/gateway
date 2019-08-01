package acl

import (
	"github.com/joaosoft/dbr"
)

type StoragePostgres struct {
	config *AclConfig
	db     *dbr.Dbr
}

func NewStoragePostgres(config *AclConfig) (*StoragePostgres, error) {
	dbr, err := dbr.New(dbr.WithConfiguration(config.Dbr))
	if err != nil {
		return nil, err
	}

	return &StoragePostgres{
		config: config,
		db:     dbr,
	}, nil
}

func (storage *StoragePostgres) GetResourceCategories(domainKey string) (Categories, error) {
	categories := make(Categories, 0)

	_, err := storage.db.
		Select([]interface{}{
			"rc.name",
			"rc.key",
			"rc.description",
			dbr.As(storage.db.
				Select("key").
				From(dbr.As(aclTableResourceCategory, "parent")).
				Where("parent.id_resource_category = rc.fk_parent_resource_category"), "parent_resource_category_key"),
			"rc.active",
			"rc.created_at",
			"rc.updated_at",
		}...).
		From(dbr.As(aclTableResourceCategory, "rc")).
		Join(dbr.As(aclTableDomain, "d"), "d.id_domain = rc.fk_domain").
		Where("d.key = ?", domainKey).
		Where("d.active").
		Where("rc.active").
		Load(&categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (storage *StoragePostgres) GetResourceCategoryPages(domainKey, resourceCategoryKey string) (Pages, error) {
	pages := make(Pages, 0)

	_, err := storage.db.
		Select([]interface{}{
			"rp.name",
			"rp.key",
			"rp.description",
			dbr.As(storage.db.
				Select("key").
				From(dbr.As(aclTableResourcePage, "parent")).
				Where("parent.id_resource_page = rp.fk_parent_resource_page"), "parent_resource_page_key"),
			"rp.active",
			"rp.created_at",
			"rp.updated_at",
		}...).
		From(dbr.As(aclTableResourcePage, "rp")).
		Join(dbr.As(aclTableDomain, "d"), "d.id_domain = rp.fk_domain").
		Join(dbr.As(aclTableResourceCategory, "rc"), "rc.id_resource_category = rp.fk_resource_category").
		Where("d.key = ?", domainKey).
		Where("rc.key = ?", resourceCategoryKey).
		Where(dbr.IsNull("rp.fk_parent_resource_page")).
		Where("rp.active").
		Where("d.active").
		Where("rc.active").
		Load(&pages)

	if err != nil {
		return nil, err
	}

	return pages, nil
}

func (storage *StoragePostgres) GetResourceCategoryPage(domainKey, resourceCategoryKey, resourcePageKey string) (*Page, error) {
	var page Page

	count, err := storage.db.
		Select([]interface{}{
			"rp.name",
			"rp.key",
			"rp.description",
			dbr.As(storage.db.
				Select("key").
				From(dbr.As(aclTableResourcePage, "parent")).
				Where("parent.id_resource_page = rp.fk_parent_resource_page"), "parent_resource_page_key"),
			"rp.active",
			"rp.created_at",
			"rp.updated_at",
		}...).
		From(dbr.As(aclTableResourcePage, "rp")).
		Join(dbr.As(aclTableDomain, "d"), "d.id_domain = rp.fk_domain").
		Join(dbr.As(aclTableResourceCategory, "rc"), "rc.id_resource_category = rp.fk_resource_category").
		Where("d.key = ?", domainKey).
		Where("rc.key = ?", resourceCategoryKey).
		Where("rp.key = ?", resourcePageKey).
		Where(dbr.IsNull("rp.fk_parent_resource_page")).
		Where("rp.active").
		Where("d.active").
		Where("rc.active").
		Load(&page)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	return &page, nil
}

func (storage *StoragePostgres) GetPageResources(domainKey, roleKey, resourceCategoryKey, resourcePageKey, idUser string) (Resources, error) {
	resources := make(Resources, 0)

	_, err := storage.db.
		Select([]interface{}{
			"rs.name",
			"rs.key",
			dbr.As("rc.key", "resource_category_key"),
			dbr.As("rp.key", "resource_page_key"),
			dbr.As("rt.key", "resource_type_key"),
			"rs.description",
			"rs.active",
			"rs.created_at",
			"rs.updated_at",
		}...).
		From(dbr.As(aclTableResource, "rs")).
		Join(dbr.As(aclTableDomain, "d"), "d.id_domain = rs.fk_domain").
		Join(dbr.As(aclTableResourcePage, "rp"), "rp.id_resource_page = rs.fk_resource_page").
		Join(dbr.As(aclTableResourceCategory, "rc"), "rc.id_resource_category = rp.fk_resource_category").
		Join(dbr.As(aclTableRoleResource, "rr"), "rr.fk_resource = rs.id_resource").
		Join(dbr.As(aclTableRole, "r"), "r.id_role = rr.fk_role").
		Join(dbr.As(aclTableResourceType, "rt"), "rt.id_resource_type = rs.fk_resource_type").
		Where("d.key = ?", domainKey).
		Where("r.key = ?", roleKey).
		Where("rp.key = ?", resourcePageKey).
		Where("rs.active").
		Where("r.active").
		Where("rr.active").
		Where("d.active").
		Where("rp.active").
		Where("rc.active").
		Where("rt.active").
		Where(dbr.NotIn("rs.id_resource",
			storage.db.
				Select("ur.fk_resource").
				From(dbr.As(aclTableUserResource, "ur")).
				Where("ur.fk_resource = rs.id_resource").
				Where("ur.active").
				Where("ur.fk_user = ?", idUser).
				Where("ur.allow = ?", false),
		)).
		Load(&resources)

	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (storage *StoragePostgres) GetPageResourcesByType(domainKey, roleKey, resourceCategoryKey, resourcePageKey, resourceTypeKey, idUser string) (Resources, error) {
	resources := make(Resources, 0)

	_, err := storage.db.
		Select([]interface{}{
			"rs.name",
			"rs.key",
			dbr.As("rc.key", "resource_category_key"),
			dbr.As("rp.key", "resource_page_key"),
			dbr.As("rt.key", "resource_type_key"),
			"rs.description",
			"rs.active",
			"rs.created_at",
			"rs.updated_at",
		}...).
		From(dbr.As(aclTableResource, "rs")).
		Join(dbr.As(aclTableDomain, "d"), "d.id_domain = rs.fk_domain").
		Join(dbr.As(aclTableResourcePage, "rp"), "rp.id_resource_page = rs.fk_resource_page").
		Join(dbr.As(aclTableResourceCategory, "rc"), "rc.id_resource_category = rp.fk_resource_category").
		Join(dbr.As(aclTableRoleResource, "rr"), "rr.fk_resource = rs.id_resource").
		Join(dbr.As(aclTableRole, "r"), "r.id_role = rr.fk_role").
		Join(dbr.As(aclTableResourceType, "rt"), "rt.id_resource_type = rs.fk_resource_type").
		Where("d.key = ?", domainKey).
		Where("r.key = ?", roleKey).
		Where("rp.key = ?", resourcePageKey).
		Where("rt.key = ?", resourceTypeKey).
		Where("rs.active").
		Where("r.active").
		Where("rr.active").
		Where("d.active").
		Where("rp.active").
		Where("rc.active").
		Where("rt.active").
		Where(dbr.NotIn("rs.id_resource",
			storage.db.
				Select("ur.fk_resource").
				From(dbr.As(aclTableUserResource, "ur")).
				Where("ur.fk_resource = rs.id_resource").
				Where("ur.active").
				Where("ur.fk_user = ?", idUser).
				Where("ur.allow = ?", false),
		)).
		Load(&resources)

	if err != nil {
		return nil, err
	}

	return resources, nil
}

func (storage *StoragePostgres) CheckEndpointAccess(domainKey, roleKey, resourceTypeKey, method, endpoint, idUser string) (bool, error) {

	allowed := Allowed{}

	// check the general resources
	count, err := storage.db.
		Select("e.check", dbr.As(dbr.Condition(dbr.Count("1"), dbr.ComparatorBigger, 0), "allow")).
		From(dbr.As(aclTableEndpoint, "e")).
		Join(dbr.As(aclTableEndpointResource, "er"), "er.fk_endpoint = e.id_endpoint").
		Join(dbr.As(aclTableResource, "rs"), "rs.id_resource = er.fk_resource").
		Join(dbr.As(aclTableResourcePage, "rp"), "rp.id_resource_page = rs.fk_resource_page").
		Join(dbr.As(aclTableResourceCategory, "rc"), "rc.id_resource_category = rp.fk_resource_category").
		Join(dbr.As(aclTableDomain, "d"), "d.id_domain = e.fk_domain").
		Join(dbr.As(aclTableRole, "r"), "r.id_role = er.fk_role").
		Join(dbr.As(aclTableResourceType, "rt"), "rt.id_resource_type = rs.fk_resource_type").
		Where("e.method = ?", method).
		Where("e.endpoint = ?", endpoint).
		Where("d.key = ?", domainKey).
		Where("r.key = ?", roleKey).
		Where("rt.key = ?", resourceTypeKey).
		Where("e.active").
		Where("er.active").
		Where("rs.active").
		Where("rc.active").
		Where("rp.active").
		Where("d.active").
		Where("r.active").
		Where("rt.active").
		GroupBy("e.check").
		Load(&allowed)

	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	if !allowed.Check {
		return true, nil
	}

	// check the user if has an override to the resource
	var userAllowed bool
	count, err = storage.db.
		Select("ue.allow").
		From(dbr.As(aclTableEndpoint, "e")).
		Join(dbr.As(aclTableEndpointResource, "er"), "er.fk_endpoint = e.id_endpoint").
		Join(dbr.As(aclTableResource, "rs"), "rs.id_resource = er.fk_resource").
		Join(dbr.As(aclTableUserEndpoint, "ue"), "ue.fk_endpoint = e.id_endpoint").
		Where("e.active").
		Where("er.active").
		Where("rs.active").
		Where("ue.active").
		Where("e.method = ?", method).
		Where("e.endpoint = ?", endpoint).
		Where("ue.fk_user = ?", idUser).
		Load(&userAllowed)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return userAllowed, nil
	}

	return allowed.Allow, nil
}
