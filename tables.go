package auth

import "fmt"

func format(schema, table string) string {
	return fmt.Sprintf("%s.%s", schema, table)
}

var (
	authTableUser = format(schemaAuth, "user")
)
