package router

import (
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/module"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoutingOrder(t *testing.T) {
	router := NewRouter()
	router.Handle(http.GET, "/api/test", GetTestHandler)
	router.Handle(http.GET, "/api/:id", GetIdHandler)

	var m *module.Module
	res, _ := router.CallHandler(m, http.GET, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(m, http.GET, "/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id with id: abc!"), res)
}

func TestRoutingWildcard(t *testing.T) {
	router := NewRouter()
	router.Handle(http.GET, "/api/test", GetTestHandler)
	router.Handle(http.GET, "/api/*", GetProxyHandler)

	var m *module.Module
	res, _ := router.CallHandler(m, http.GET, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(m, http.GET, "/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/abc proxy!"), res)
}

func TestRouter(t *testing.T) {
	router := NewRouter()
	router.Handle(http.GET, "/api/test", GetTestHandler)
	router.Handle(http.GET, "/api/:id", GetIdHandler)
	router.Handle(http.GET, "/api/:id/test", GetIdTestHandler)
	router.Handle(http.POST, "/api/test", PostTestHandler)

	var m *module.Module
	res, _ := router.CallHandler(m, http.GET, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(m, http.POST, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the POST: /api/test!"), res)

	res, _ = router.CallHandler(m, http.GET, "/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id with id: abc!"), res)

	res, _ = router.CallHandler(m, http.GET, "/api/abc/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id/test with id: abc!"), res)
}

func GetTestHandler(m *module.Module, r *Request) ([]byte, error) {
	return []byte(fmt.Sprintf("Hello, this is the GET: %s!", r.Path)), nil
}

func PostTestHandler(m *module.Module, r *Request) ([]byte, error) {
	return []byte(fmt.Sprintf("Hello, this is the POST: %s!", r.Path)), nil
}

func GetIdHandler(m *module.Module, r *Request) ([]byte, error) {
	if id, ok := r.Params["id"]; ok {
		message := fmt.Sprintf("Hello, this is the GET: %s with id: %s!", r.Pattern, id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func GetIdTestHandler(m *module.Module, r *Request) ([]byte, error) {
	if id, ok := r.Params["id"]; ok {
		message := fmt.Sprintf("Hello, this is the GET: %s with id: %s!", r.Pattern, id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func GetProxyHandler(m *module.Module, r *Request) ([]byte, error) {
	return []byte(fmt.Sprintf("Hello, this is the GET: %s proxy!", r.Path)), nil
}
