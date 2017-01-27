// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/alligrader/gradebook-api/designs
// --out=$(GOPATH)/src/github.com/alligrader/gradebook-api/src
// --version=v1.1.0-dirty
//
// API "GradebookAPI": Application Security
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

type (
	// Private type used to store auth handler info in request context
	authMiddlewareKey string
)

// UseJWTMiddleware mounts the jwt auth middleware onto the service.
func UseJWTMiddleware(service *goa.Service, middleware goa.Middleware) {
	service.Context = context.WithValue(service.Context, authMiddlewareKey("jwt"), middleware)
}

// NewJWTSecurity creates a jwt security definition.
func NewJWTSecurity() *goa.JWTSecurity {
	def := goa.JWTSecurity{
		In:       goa.LocHeader,
		Name:     "Authorization",
		TokenURL: "https://example.com/token",
		Scopes: map[string]string{
			"profiles:read":  "Can read profiles",
			"profiles:write": "Can update profiles",
			"projects:read":  "Can read projects",
			"projects:write": "Can update projects",
			"settings:read":  "Can read settings",
			"settings:write": "Can update settings",
		},
	}
	def.Description = "For routes that require an authorization token"
	return &def
}

// handleSecurity creates a handler that runs the auth middleware for the security scheme.
func handleSecurity(schemeName string, h goa.Handler, scopes ...string) goa.Handler {
	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		scheme := ctx.Value(authMiddlewareKey(schemeName))
		am, ok := scheme.(goa.Middleware)
		if !ok {
			return goa.NoAuthMiddleware(schemeName)
		}
		ctx = goa.WithRequiredScopes(ctx, scopes)
		return am(h)(ctx, rw, req)
	}
}
