package lambda

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

const (
	nodeLiteral  int = 0
	nodeCapture  int = 1
	nodeOptional int = 2
	nodeGreedy   int = 4
)

const (
	runeWildcard  rune = 42
	runeQuestion  rune = 63
	runeLeftCurly rune = 123
)

type routeNode struct {
	Type     int
	Text     string
	Value    string
	Handlers map[string]Handler
	Nodes    []*routeNode
}

//Router manages routes mapping ang matching
type Router struct {
	routes *routeNode
}

//Handler function to execute by lambda
type Handler func(req Request) (events.APIGatewayProxyResponse, error)

//Request contains handler and route parameters
type Request struct {
	Method       string
	Stage        string
	ResourcePath string
	Proxy        string
	Path         string
	Parameters   map[string]string
	Body         string
	handler      Handler
}

//Execute runs handler
func (req *Request) Execute() (events.APIGatewayProxyResponse, error) {
	log.Printf("%s: %s\n%s", req.Method, req.Path, req.Body)
	return req.handler(*req)
}

func newRouteTree() *routeNode {
	return &routeNode{
		Handlers: make(map[string]Handler),
	}
}

//Add adds a handler mapped to HTTP method and path
func (rt *Router) Add(pattern string, method string, handler Handler) error {
	if rt.routes == nil {
		rt.routes = newRouteTree()
	}

	return rt.routes.Add(pattern, method, handler)
}

//GetRequest maps lambda request from router map
func (rt *Router) GetRequest(req events.APIGatewayProxyRequest, proxy bool) (Request, bool) {
	var (
		path string
		tmp  = Request{
			Path:         req.Path,
			Stage:        req.RequestContext.Stage,
			ResourcePath: req.RequestContext.ResourcePath,
		}
	)

	if rt.routes == nil {
		return tmp, false
	}

	if tmp.ResourcePath == "/{proxy+}" && req.PathParameters != nil {
		path, _ = req.PathParameters["proxy"]
		tmp.Proxy = path
	} else {
		path = tmp.Path
		if tmp.Stage != "" {
			path = strings.Replace(path, req.RequestContext.Stage, "", 1)
		}
	}

	if path != "" {
		toks := strings.Split(path, "?")
		path = strings.TrimSuffix(toks[0], "/")
	}

	r, ok := rt.routes.Match(path, req.HTTPMethod)
	if ok {
		r.Body = req.Body
	} else {
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
		r.Path = path
		r.Stage = tmp.Stage
		r.Proxy = tmp.Proxy
		r.ResourcePath = tmp.ResourcePath
	}

	if req.QueryStringParameters != nil {
		if r.Parameters == nil {
			r.Parameters = req.QueryStringParameters
		} else {
			for k, v := range req.QueryStringParameters {
				_, ok := r.Parameters[k]
				if !ok {
					r.Parameters[k] = v
				}
			}
		}
	}

	return r, ok
}

//Execute executes lambda request
func (rt *Router) Execute(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	rq, ok := rt.GetRequest(req, true)
	if !ok {
		return Failure(404, fmt.Errorf("Invalid endpoint: %s; Stage: %s; Resource Path: %s; Proxy: %s", rq.Path, rq.Stage, rq.ResourcePath, rq.Proxy), false)
	}
	return rq.Execute()
}

func (tree *routeNode) Add(pattern string, method string, handler Handler) error {
	node := tree

	toks := strings.Split(pattern, "/")
	for _, tok := range toks {
		if tok != "" {
			var (
				tmp = &routeNode{}
				ok  bool
			)

			runes := []rune(tok)
			if runes[0] == runeLeftCurly {
				var (
					i    = 1
					size = len(runes)
				)

				for i < size {
					if runes[i] == runeQuestion {
						tmp.Text = string(runes[1:i])
						tmp.Value = string(runes[i+1 : size-1])
						tmp.Type = nodeOptional
						break
					} else if runes[i] == runeWildcard {
						tmp.Text = string(runes[1:i])
						tmp.Type = nodeGreedy
						break
					}
					i++
				}

				if tmp.Text == "" {
					tmp.Text = string(runes[1 : size-1])
					tmp.Type = nodeCapture
				}

			} else {
				tmp.Text = tok
			}

			for _, nd := range node.Nodes {
				if nd.Text == tmp.Text && nd.Type == tmp.Type {
					tmp = nd
					ok = true
					break
				}
			}

			if !ok {
				if len(node.Nodes) > 0 {
					nodes := append(node.Nodes, tmp)
					sort.Slice(nodes, func(i, j int) bool {
						return nodes[i].Type < nodes[j].Type
					})
					node.Nodes = nodes
				} else {
					node.Nodes = []*routeNode{tmp}
				}
			}

			node = tmp
		}
	}

	if node.Handlers == nil {
		node.Handlers = make(map[string]Handler)
	}

	_, ok := node.Handlers[method]
	if ok {
		return fmt.Errorf("%s %s already exists", method, pattern)
	}

	node.Handlers[method] = handler
	return nil
}

func (tree *routeNode) Match(path string, method string) (Request, bool) {
	var (
		node   = tree
		params = make(map[string]string)
		toks   []string
	)

	for _, tok := range strings.Split(path, "/") {
		if tok != "" {
			toks = append(toks, tok)

			var (
				found bool
			)

			for _, nd := range node.Nodes {
				var ok bool

				if nd.Text == tok && nd.Type == nodeLiteral {
					ok = true
				} else if nd.Type == nodeCapture {
					params[nd.Text] = tok
					ok = true
				} else if nd.Type == nodeOptional {
					params[nd.Text] = nd.Value
					ok = true
				}

				if ok {
					node = nd
					found = true
					break
				}
			}

			if !found {
				return Request{}, false
			}
		}
	}

	if node.Handlers == nil {
		for _, nd := range node.Nodes {
			if nd.Type == nodeOptional {
				params[nd.Text] = nd.Value
				node = nd
				break
			}
		}
	}

	if node.Handlers == nil {
		return Request{}, false
	}

	h, ok := node.Handlers[method]

	return Request{
		handler:    h,
		Method:     method,
		Path:       "/" + strings.Join(toks, "/"),
		Parameters: params,
	}, ok
}
