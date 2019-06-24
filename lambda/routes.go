package lambda

import (
	"fmt"
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
	Method         string
	Path           string
	PathParameters map[string]string
	Body           string
	handler        Handler
}

//Execute executes handler
func (req *Request) Execute() (events.APIGatewayProxyResponse, error) {
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
	if rt.routes == nil {
		return Request{}, false
	}

	var (
		path string
	)

	if proxy {
		if req.PathParameters != nil {
			path, _ = req.PathParameters["proxy"]
			if path != "" {
				toks := strings.Split(path, "?")
				path = strings.TrimSuffix(toks[0], "/")
			}
		}
	} else {
		path = req.Path
	}

	r, ok := rt.routes.Match(path, req.HTTPMethod)
	if ok {
		r.Body = req.Body
	}

	return r, ok
}

func (tree *routeNode) Add(pattern string, method string, handler Handler) error {
	node := tree

	toks := strings.Split(pattern, "/")
	for _, tok := range toks {
		if tok != "" {
			var (
				tmp routeNode
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
					tmp = *nd
					ok = true
					break
				}
			}

			if !ok {
				if len(node.Nodes) > 0 {
					nodes := append(node.Nodes, &tmp)
					sort.Slice(nodes, func(i, j int) bool {
						return nodes[i].Type < nodes[j].Type
					})
					node.Nodes = nodes
				} else {
					node.Nodes = []*routeNode{&tmp}
				}
			}

			node = &tmp
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
		handler:        h,
		Method:         method,
		Path:           "/" + strings.Join(toks, "/"),
		PathParameters: params,
	}, ok
}
