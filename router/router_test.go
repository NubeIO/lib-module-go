package router

import (
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/shared"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoutingOrder(t *testing.T) {
	router := NewRouter()
	router.Handle(http.GET, "/api/test", GetTestHandler)
	router.Handle(http.GET, "/api/:id", GetIdHandler)

	var module *shared.Module
	res, _ := router.CallHandler(module, http.GET, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(module, http.GET, "/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id with id: abc!"), res)
}

func TestRoutingWildcard(t *testing.T) {
	router := NewRouter()
	router.Handle(http.GET, "/api/test", GetTestHandler)
	router.Handle(http.GET, "/api/*", GetProxyHandler)

	var module *shared.Module
	res, _ := router.CallHandler(module, http.GET, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(module, http.GET, "/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/abc proxy!"), res)
}

func TestRouter(t *testing.T) {
	router := NewRouter()
	router.Handle(http.GET, "/api/test", GetTestHandler)
	router.Handle(http.GET, "/api/:id", GetIdHandler)
	router.Handle(http.GET, "/api/:id/test", GetIdTestHandler)
	router.Handle(http.POST, "/api/test", PostTestHandler)

	var module *shared.Module
	res, _ := router.CallHandler(module, http.GET, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(module, http.POST, "/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the POST: /api/test!"), res)

	res, _ = router.CallHandler(module, http.GET, "/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id with id: abc!"), res)

	res, _ = router.CallHandler(module, http.GET, "/api/abc/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id/test with id: abc!"), res)
}

func GetTestHandler(m *shared.Module, path string, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return []byte("Hello, this is the GET: /api/test!"), nil
}

func PostTestHandler(m *shared.Module, path string, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return []byte("Hello, this is the POST: /api/test!"), nil
}

func GetIdHandler(m *shared.Module, path string, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	if id, ok := params["id"]; ok {
		message := fmt.Sprintf("Hello, this is the GET: /api/:id with id: %s!", id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func GetIdTestHandler(m *shared.Module, path string, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	if id, ok := params["id"]; ok {
		message := fmt.Sprintf("Hello, this is the GET: /api/:id/test with id: %s!", id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func GetProxyHandler(m *shared.Module, path string, params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return []byte(fmt.Sprintf("Hello, this is the GET: %s proxy!", path)), nil
}
