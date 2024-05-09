package router

import (
	"fmt"
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/lib-module-go/nmodule"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type Request struct {
	Path        string
	Pattern     string
	PathParams  map[string]string
	QueryParams url.Values
	Headers     http.Header
	Body        []byte
}

// HandlerFunc defines the type for request handlers
type HandlerFunc func(*nmodule.Module, *Request) ([]byte, error)

// Router is a simple router that maps URL patterns to handlers
type Router struct {
	routes          map[string]map[nhttp.Method]HandlerFunc
	orderedPatterns []string
	needsReorder    bool
}

// NewRouter creates a new Router instance
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]map[nhttp.Method]HandlerFunc),
	}
}

func (router *Router) OrderPatterns() []string {
	if router.needsReorder {
		return router.orderedPatterns
	}
	var patternsWithWildcard []string
	var staticPatterns []string
	var dynamicPatterns []string
	dynamicPatternCount := make(map[string]int) // Initialize the map

	for pattern := range router.routes {
		if strings.Contains(pattern, "/*") {
			patternsWithWildcard = append(patternsWithWildcard, pattern)
		} else if containsDynamicSegments(pattern) {
			dynamicPatterns = append(dynamicPatterns, pattern)
			// Count the number of dynamic segments in the pattern
			dynamicPatternCount[pattern] = countDynamicSegments(pattern)
		} else {
			staticPatterns = append(staticPatterns, pattern)
		}
	}

	// Sort dynamic patterns based on the number of dynamic segments
	sort.Slice(dynamicPatterns, func(i, j int) bool {
		return dynamicPatternCount[dynamicPatterns[i]] > dynamicPatternCount[dynamicPatterns[j]]
	})

	router.orderedPatterns = append(staticPatterns, dynamicPatterns...)
	router.orderedPatterns = append(router.orderedPatterns, patternsWithWildcard...)

	return router.orderedPatterns
}

// Helper function to count the number of dynamic segments in a pattern
func countDynamicSegments(pattern string) int {
	count := 0
	for _, segment := range strings.Split(pattern, "/") {
		if strings.HasPrefix(segment, ":") {
			count++
		}
	}
	return count
}

func containsDynamicSegments(pattern string) bool {
	return strings.Contains(pattern, ":")
}

// Handle registers a handler for a specific pattern and HTTP method
func (router *Router) Handle(method nhttp.Method, pattern string, handler HandlerFunc) {
	if _, exists := router.routes[pattern]; !exists {
		router.routes[pattern] = make(map[nhttp.Method]HandlerFunc)
	}
	router.routes[pattern][method] = handler
}

func (router *Router) CallHandler(module *nmodule.Module, method nhttp.Method, urlString string, headers http.Header, body []byte) ([]byte, error) {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	orderedPatterns := router.OrderPatterns()
	for _, pattern := range orderedPatterns {
		log.Errorf("pattern %v %v", pattern, parsedURL.Path)
		if params, ok := match(pattern, parsedURL.Path); ok {
			log.Errorf("pattern match %v", urlString)
			if handlers, exists := router.routes[pattern]; exists {
				if handler, exists := handlers[method]; exists {
					return handler(module, &Request{
						Path:        parsedURL.Path,
						Pattern:     pattern,
						PathParams:  params,
						QueryParams: parsedURL.Query(),
						Headers:     headers,
						Body:        body,
					})
				}
			}
		}
	}
	return nil, fmt.Errorf("handler not found for %s: %s", method, urlString)
}

// match checks if the given pattern matches the request path and extracts parameters
func match(pattern, path string) (map[string]string, bool) {
	patternSegments := strings.Split(pattern, "/")
	pathSegments := strings.Split(path, "/")

	if len(patternSegments) != len(pathSegments) && !strings.HasSuffix(pattern, "*") {
		return nil, false
	}

	params := make(map[string]string)
	for i, segment := range patternSegments {
		if segment == "*" {
			return params, true
		}
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
