package pathvar

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMatchedRoutePath(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/api/users/123", nil)
	assert.Nil(t, err)

	r = WithMatchedRoutePath(r, "/api/users/:id")
	assert.Equal(t, "/api/users/:id", GetMatchedRoutePath(r))
}

func TestGetMatchedRoutePathEmpty(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.Nil(t, err)

	assert.Equal(t, "", GetMatchedRoutePath(r))
}

func TestGetMatchedRoutePathStaticRoute(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/api/health", nil)
	assert.Nil(t, err)

	r = WithMatchedRoutePath(r, "/api/health")
	assert.Equal(t, "/api/health", GetMatchedRoutePath(r))
}
