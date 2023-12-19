package router

import (
	"fmt"
	"github.com/NubeIO/lib-module-go/module"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRoutingOrder(t *testing.T) {
	router := NewRouter()
	router.Handle(nhttp.GET, "/api/test", GetTestHandler)
	router.Handle(nhttp.GET, "/api/:id", GetIdHandler)

	var m *module.Module
	headers := http.Header{}
	headers["Authorization"] = []string{"Bearer abc"}
	res, _ := router.CallHandler(m, nhttp.GET, "/api/test?abc=test", headers, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(m, nhttp.GET, "/api/abc", nil, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id with id: abc!"), res)
}

func TestRoutingWildcard(t *testing.T) {
	router := NewRouter()
	router.Handle(nhttp.GET, "/api/test", GetTestHandler)
	router.Handle(nhttp.GET, "/api/*", GetProxyHandler)

	var m *module.Module
	res, _ := router.CallHandler(m, nhttp.GET, "/api/test", nil, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(m, nhttp.GET, "/api/abc", nil, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/abc proxy!"), res)
}

func TestRouter(t *testing.T) {
	router := NewRouter()
	router.Handle(nhttp.GET, "/api/test", GetTestHandler)
	router.Handle(nhttp.GET, "/api/:id", GetIdHandler)
	router.Handle(nhttp.GET, "/api/:id/test", GetIdTestHandler)
	router.Handle(nhttp.POST, "/api/test", PostTestHandler)

	var m *module.Module
	res, _ := router.CallHandler(m, nhttp.GET, "/api/test", nil, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/test!"), res)

	res, _ = router.CallHandler(m, nhttp.POST, "/api/test", nil, nil)
	assert.Equal(t, []byte("Hello, this is the POST: /api/test!"), res)

	res, _ = router.CallHandler(m, nhttp.GET, "/api/abc", nil, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id with id: abc!"), res)

	res, _ = router.CallHandler(m, nhttp.GET, "/api/abc/test", nil, nil)
	assert.Equal(t, []byte("Hello, this is the GET: /api/:id/test with id: abc!"), res)
}

func GetTestHandler(m *module.Module, r *Request) ([]byte, error) {
	fmt.Printf("Query params: abc = %s\n", r.QueryParams.Get("abc"))
	fmt.Printf("Header Authorization = %s\n", r.Headers.Get("Authorization"))
	return []byte(fmt.Sprintf("Hello, this is the GET: %s!", r.Path)), nil
}

func PostTestHandler(m *module.Module, r *Request) ([]byte, error) {
	return []byte(fmt.Sprintf("Hello, this is the POST: %s!", r.Path)), nil
}

func GetIdHandler(m *module.Module, r *Request) ([]byte, error) {
	if id, ok := r.PathParams["id"]; ok {
		message := fmt.Sprintf("Hello, this is the GET: %s with id: %s!", r.Pattern, id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func GetIdTestHandler(m *module.Module, r *Request) ([]byte, error) {
	if id, ok := r.PathParams["id"]; ok {
		message := fmt.Sprintf("Hello, this is the GET: %s with id: %s!", r.Pattern, id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func GetProxyHandler(m *module.Module, r *Request) ([]byte, error) {
	return []byte(fmt.Sprintf("Hello, this is the GET: %s proxy!", r.Path)), nil
}
