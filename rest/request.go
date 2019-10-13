package rest

import "strings"

//Request is REST request
type Request struct {
	Path                  string
	HTTPMethod            string
	Headers               map[string]string
	QueryStringParameters map[string]string
	PathParameters        map[string]string
	Body                  string
}

//Match finds a matching handler
func (r *Request) Match(router Router) (Execute, bool) {
	var node = router.root

	toks := strings.Split(r.Path, "/")
	for _, tok := range toks {
		if tok == "" {
			for _, p := range node.Parameters {
				if p.Type == 2 {
					if r.PathParameters == nil {
						r.PathParameters = make(map[string]string)
					}
					r.PathParameters[p.Parameter] = p.Value
					node = p
					continue
				}
			}
			continue
		} else {
			if node.Nodes != nil {
				nd, ok := node.Nodes[tok]
				if ok {
					node = nd
					continue
				}
			}

			var match bool
			for _, p := range node.Parameters {
				if r.PathParameters == nil {
					r.PathParameters = make(map[string]string)
				}
				r.PathParameters[p.Parameter] = tok
				node = p
				match = true
				continue
			}
			if match {
				continue
			}
		}

		return nil, false
	}

	h, ok := node.Handlers[r.HTTPMethod]
	if ok {
		return h, true
	}

	return nil, false
}
