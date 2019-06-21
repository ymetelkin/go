package lambda

import (
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

//Handler is a function used by lambda
type Handler interface {
	Execute() (events.APIGatewayProxyResponse, error)
}

//ProxyRouter manages routes mapping ang matching
type ProxyRouter struct {
	routes *routeNode
}

//Add adds a handler mapped to HTTP method and path
func (pr *ProxyRouter) Add(pattern string, method string, handler Handler) error {
	if pr.routes == nil {
		pr.routes = newRouteTree()
	}

	return pr.routes.Add(pattern, method, handler)
}

//GetHandler maps handler from router map
func (pr *ProxyRouter) GetHandler(req events.APIGatewayProxyRequest) (RouteMatch, bool) {
	if pr.routes == nil {
		return RouteMatch{}, false
	}

	var (
		path string
	)

	if req.PathParameters != nil {
		path, _ = req.PathParameters["proxy"]
		if path != "" {
			toks := strings.Split(path, "?")
			path = strings.TrimSuffix(toks[0], "/")
		}
	}

	return pr.routes.Match(path, req.HTTPMethod)
}
