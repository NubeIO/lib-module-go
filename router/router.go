package router

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
	"strings"
)

// HandlerFunc defines the type for request handlers
type HandlerFunc func(map[string]string, nargs.Args, []byte) ([]byte, error)

// Router is a simple router that maps URL patterns to handlers
type Router struct {
	routes map[string]HandlerFunc
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

// Handle registers a handler for a specific pattern
func (router *Router) Handle(pattern string, handler HandlerFunc) {
	router.routes[pattern] = handler
}

// CallHandler calls the registered handler for the given pattern
func (router *Router) CallHandler(path string, args nargs.Args, body []byte) ([]byte, error) {
	for pattern, handler := range router.routes {
		if params, ok := match(pattern, path); ok {
			return handler(params, args, body)
		}
	}
	return nil, fmt.Errorf("handler not found for path: %s", path)
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
