package http

import "fmt"

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PATCH  Method = "PATCH"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

func StringToMethod(methodStr string) (Method, error) {
	switch methodStr {
	case "GET", "POST", "PATCH", "PUT", "DELETE":
		return Method(methodStr), nil
	default:
		return "", fmt.Errorf("unsupported HTTP method: %s", methodStr)
	}
}
