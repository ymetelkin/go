package rest

import (
	"fmt"
	"strings"
)

//Execute handles request
type Execute func(req Request) Response

//Router is collection of routes
type Router struct {
	root *routeNode
}

type routeNode struct {
	Type       int
	Value      string
	Parameter  string
	Handlers   map[string]Execute
	Nodes      map[string]*routeNode
	Parameters []*routeNode
}

//Add method adds handler to collection of routes
func (r *Router) Add(path string, method string, handler Execute) {
	if r.root == nil {
		r.root = newNode("")
	}
	r.root.Add(path, method, handler)
}

//Execute handles request
func (r *Router) Execute(req Request) Response {
	exe, ok := req.Match(*r)
	if ok {
		return exe(req)
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"

	return Response{
		StatusCode: 404,
		Headers:    headers,
		Body:       fmt.Sprintf(`{"error":"%s: %s not supported"}`, req.HTTPMethod, req.Path),
	}
}

func (r *routeNode) Add(path string, method string, handler Execute) {
	node := r

	toks := strings.Split(path, "/")
	for _, tok := range toks {
		if tok == "" {
			continue
		}

		if tok[0] == '{' {
			tks := strings.Split(tok[1:len(tok)-1], "?")
			nd := routeNode{
				Type:      1,
				Parameter: tks[0],
			}
			if len(tks) == 2 {
				nd.Type = 2
				nd.Value = tks[1]
			}
			var found bool
			for _, p := range node.Parameters {
				if p.Parameter == nd.Parameter {
					node = p
					found = true
					break
				}
			}
			if !found {
				nd.Handlers = make(map[string]Execute)
				nd.Nodes = make(map[string]*routeNode)
				node.Parameters = append(node.Parameters, &nd)
				node = &nd
			}
		} else {
			nd, ok := node.Nodes[tok]
			if !ok {
				nd = newNode(tok)
				node.Nodes[tok] = nd
			}
			node = nd
		}
	}

	node.Handlers[method] = handler
}

func newNode(value string) *routeNode {
	return &routeNode{
		Value:    value,
		Handlers: make(map[string]Execute),
		Nodes:    make(map[string]*routeNode),
	}
}
