package router

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_GetHostNetworks(t *testing.T) {
	router := NewRouter()
	router.Handle("/api/test", ApiTestHandler)
	router.Handle("/api/:id", IdHandler)
	router.Handle("/api/:id/test", IdTestHandler)

	res, _ := router.CallHandler("/api/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the /api/test!"), res)

	res, _ = router.CallHandler("/api/abc", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the /api/:id with id: abc!"), res)

	res, _ = router.CallHandler("/api/abc/test", nargs.Args{}, nil)
	assert.Equal(t, []byte("Hello, this is the /api/:id/test with id: abc!"), res)
}

func ApiTestHandler(params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	return []byte("Hello, this is the /api/test!"), nil
}

func IdHandler(params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	if id, ok := params["id"]; ok {
		message := fmt.Sprintf("Hello, this is the /api/:id with id: %s!", id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}

func IdTestHandler(params map[string]string, args nargs.Args, body []byte) ([]byte, error) {
	if id, ok := params["id"]; ok {
		message := fmt.Sprintf("Hello, this is the /api/:id/test with id: %s!", id)
		return []byte(message), nil
	}
	return nil, fmt.Errorf("missing id parameter")
}
