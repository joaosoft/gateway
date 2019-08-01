package acl

import "fmt"

func format(schema, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

var (
	aclTableDomain           = format(schemaAcl, "domain")
	aclTableRole             = format(schemaAcl, "role")
	aclTableResourceCategory = format(schemaAcl, "resource_category")
	aclTableResourcePage     = format(schemaAcl, "resource_page")
	aclTableResourceType     = format(schemaAcl, "resource_type")
	aclTableResource         = format(schemaAcl, "resource")
	aclTableRoleResource     = format(schemaAcl, "role_resource")
	aclTableUserResource     = format(schemaAcl, "user_resource")

	aclTableEndpoint         = format(schemaAcl, "endpoint")
	aclTableEndpointResource = format(schemaAcl, "endpoint_resource")
	aclTableUserEndpoint     = format(schemaAcl, "user_endpoint")
)
