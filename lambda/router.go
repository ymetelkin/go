package lambda

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

//Handler is a function used by lambda
type Handler func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

//ProxyRouter manages routes mapping ang matching
type ProxyRouter struct {
	routes map[string]map[string]Handler
}

//Add adds a handler mapped to HTTP method and path
func (pr *ProxyRouter) Add(method string, path string, handler Handler) {
	if pr.routes == nil {
		pr.routes = make(map[string]map[string]Handler)
	}

	hs, ok := pr.routes[method]
	if !ok {
		hs = make(map[string]Handler)
	}

	hs[path] = handler
	pr.routes[method] = hs
}

//GetHandler maps handler from router map
func (pr *ProxyRouter) GetHandler(req events.APIGatewayProxyRequest) (Handler, bool) {
	var (
		hs   map[string]Handler
		h    Handler
		ok   bool
		path string
	)

	if pr.routes == nil {
		return h, ok
	}

	hs, ok = pr.routes[req.HTTPMethod]
	if !ok {
		return h, ok
	}

	if req.PathParameters != nil {
		path, _ = req.PathParameters["proxy"]
		if path != "" {
			toks := strings.Split(path, "?")
			path = strings.TrimSuffix(toks[0], "/")
		}
	}

	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	h, ok = hs[path]
	return h, ok
}
