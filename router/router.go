package router

import (
	"fmt"
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/lib-module-go/shared"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"strings"
)

// HandlerFunc defines the type for request handlers
type HandlerFunc func(*shared.Module, map[string]string, nargs.Args, []byte) ([]byte, error)

// Router is a simple router that maps URL patterns to handlers
type Router struct {
	routes map[string]map[http.Method]HandlerFunc
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[http.Method]HandlerFunc),
	}
}

// Handle registers a handler for a specific pattern and HTTP method
func (router *Router) Handle(method http.Method, pattern string, handler HandlerFunc) {
	if _, exists := router.routes[pattern]; !exists {
		router.routes[pattern] = make(map[http.Method]HandlerFunc)
	}
	router.routes[pattern][method] = handler
}

func (router *Router) CallHandler(module *shared.Module, method http.Method, path string, args nargs.Args, body []byte) ([]byte, error) {
	for pattern, handlers := range router.routes {
		if params, ok := match(pattern, path); ok {
			if handler, exists := handlers[method]; exists {
				return handler(module, params, args, body)
			}
		}
	}
	return nil, fmt.Errorf("handler not found for path: %s and method: %s", path, method)
}

// match checks if the given pattern matches the request path and extracts parameters
func match(pattern, path string) (map[string]string, bool) {
	patternSegments := strings.Split(pattern, "/")
	pathSegments := strings.Split(path, "/")

	if len(patternSegments) != len(pathSegments) {
		return nil, false
	}

	params := make(map[string]string)
	for i, segment := range patternSegments {
		if strings.HasPrefix(segment, ":") {
			// Capture dynamic parts and store in params map
			paramName := strings.TrimPrefix(segment, ":")
			params[paramName] = pathSegments[i]
		} else if segment != pathSegments[i] {
			return nil, false
		}
	}

	return params, true
}
