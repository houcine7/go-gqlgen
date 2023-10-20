//go:build tools
// +build tools

package tools

import (
	_ "database/sql"

	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	_ "golang.org/x/crypto/bcrypt"
)


