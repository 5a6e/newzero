package pathvar

import (
	"context"
	"net/http"
)

var (
	matchedRoutePath = contextKey("matchedRoutePath")
	routeMeta        = contextKey("routeMeta")
)

// GetMatchedRoutePath returns the matched route path pattern from the request context.
// For example, if the registered route is "/api/users/:id" and the request is
// "GET /api/users/123", this returns "/api/users/:id".
func GetMatchedRoutePath(r *http.Request) string {
	p, _ := r.Context().Value(matchedRoutePath).(string)
	return p
}

// GetRouteFromContext returns the matched route path pattern from the context.
func GetRouteFromContext(ctx context.Context) string {
	p, _ := ctx.Value(matchedRoutePath).(string)
	return p
}

// WithMatchedRoutePath writes the matched route path pattern into the request context.
func WithMatchedRoutePath(r *http.Request, path string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), matchedRoutePath, path))
}

// GetMeta returns the route meta from the request context.
func GetMeta(r *http.Request) map[string]string {
	m, _ := r.Context().Value(routeMeta).(map[string]string)
	return m
}

// GetMetaFromContext returns the route meta from the context.
func GetMetaFromContext(ctx context.Context) map[string]string {
	m, _ := ctx.Value(routeMeta).(map[string]string)
	return m
}

// WithMeta writes the route meta into the request context.
func WithMeta(r *http.Request, meta map[string]string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), routeMeta, meta))
}
