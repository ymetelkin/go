package lambda

import (
	"fmt"
	"sort"
	"strings"
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

//RouteMatch contains handler and route parameters
type RouteMatch struct {
	Params  map[string]string
	Handler Handler
}

func newRouteTree() *routeNode {
	return &routeNode{
		Handlers: make(map[string]Handler),
	}
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

func (tree *routeNode) Match(path string, method string) (RouteMatch, bool) {
	var (
		node   = tree
		params = make(map[string]string)
	)

	for _, tok := range strings.Split(path, "/") {
		if tok != "" {
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
				return RouteMatch{}, false
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
		return RouteMatch{}, false
	}

	h, ok := node.Handlers[method]

	return RouteMatch{
		Handler: h,
		Params:  params,
	}, ok
}
