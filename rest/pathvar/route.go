package pathvar

import (
	"context"
	"net/http"
)

var matchedRoutePath = contextKey("matchedRoutePath")

// GetMatchedRoutePath returns the matched route path pattern from the request context.
// For example, if the registered route is "/api/users/:id" and the request is
// "GET /api/users/123", this returns "/api/users/:id".
func GetMatchedRoutePath(r *http.Request) string {
	p, _ := r.Context().Value(matchedRoutePath).(string)
	return p
}

// WithMatchedRoutePath writes the matched route path pattern into the request context.
func WithMatchedRoutePath(r *http.Request, path string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), matchedRoutePath, path))
}
