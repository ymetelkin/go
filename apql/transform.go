package apql

import (
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
)

//Transform transforms APQL to Qury DSL
type Transform struct {
	Fields map[string]field
}

//New constructs Transform
func New() Transform {
	return Transform{
		Fields: getFields(),
	}
}

//Query transforms APQL to Qury DSL
func (tr *Transform) Query(apql string) (json.Object, error) {
	toks, err := tokenize(apql)
	if err != nil {
		return json.Object{}, err
	}
	set, err := booleanize(toks)
	if err != nil {
		return json.Object{}, err
	}

	query := tr.setToJSON(set)
	return query, nil
}

func (tr *Transform) queryToJSON(q query) json.Object {
	f := strings.ToLower(q.Field)
	if strings.HasPrefix(f, "@") {
		f = strings.TrimLeft(f, "@")
	}

	es, ok := tr.Fields[f]
	if ok {
		if es.Replace.IsEmpty() {
			var (
				v      string
				m, txt bool
				qs     json.Array
			)

			if q.Value == "" {
				v = strings.Join(q.Values, " ")
				m = true
			} else {
				v = q.Value
			}

			for _, e := range es.Fields {
				if e.Transform != nil {
					for _, t := range e.Transform {
						if t == "lowercase" {
							v = strings.ToLower(v)
						}
					}
				}

				switch e.Type {
				case "text":
					txt = true
					qs.AddObject(textQueryJSON(e.Name, v, false))
				case "keyword":
					if m {
						qs.AddObject(textsQueryJSON(e.Name, strings.Split(v, " "), true))
					} else {
						qs.AddObject(textQueryJSON(e.Name, v, true))
					}
				case "integer":
					if m {
						ints := make([]int, 0)
						for _, v = range q.Values {
							i, err := strconv.Atoi(v)
							if err == nil {
								ints = append(ints, i)
							}
						}
						if len(ints) > 0 {
							qs.AddObject(intsQueryJSON(e.Name, ints))
						}
					} else {
						i, err := strconv.Atoi(q.Value)
						if err == nil {
							qs.AddObject(intQueryJSON(e.Name, i))
						}
					}
				case "boolean":
					b, err := strconv.ParseBool(v)
					if err == nil {
						qs.AddObject(boolQueryJSON(e.Name, b))
					}
				}
			}

			size := qs.Length()

			if es.And.Field == "" {
				if size == 0 {
					return json.Object{}
				} else if size == 1 {
					test, _ := qs.GetObject(0)
					return test
				} else {
					bl := json.Object{}
					bl.AddArray("should", qs)
					jo := json.Object{}
					jo.AddObject("bool", bl)
					return jo
				}
			} else {
				if size == 1 {
					test, err := qs.GetObject(0)
					if err == nil {
						bl := json.Object{}
						fr := textQueryJSON(es.And.Field, es.And.Value, true)

						if txt {
							bl.AddObject("must", test)
							bl.AddObject("filter", fr)
						} else {
							ja := json.Array{}
							ja.AddObject(test)
							ja.AddObject(fr)
							bl.AddArray("filter", ja)
						}
						qr := json.Object{}
						qr.AddObject("bool", bl)

						nest := json.Object{}
						nest.AddString("path", es.And.Path)
						nest.AddObject("query", qr)

						jo := json.Object{}
						jo.AddObject("nested", nest)
						return jo
					}
				}
				return json.Object{}
			}
		}

		return es.Replace
	}

	if q.Value == "" {
		return textsQueryJSON(f, q.Values, false)
	} else {
		return textQueryJSON(f, q.Value, false)
	}
}

func (tr *Transform) setToJSON(st set) json.Object {
	query := json.Object{}

	qs := json.Array{}

	if st.Queries != nil && len(st.Queries) > 0 {
		for _, q := range st.Queries {
			jo := tr.queryToJSON(q)
			if !jo.IsEmpty() {
				qs.AddObject(jo)
			}
		}
	}

	if st.Sets != nil && len(st.Sets) > 0 {
		for _, s := range st.Sets {
			jo := tr.setToJSON(s)
			if !jo.IsEmpty() {
				qs.AddObject(jo)
			}
		}
	}

	size := qs.Length()

	if size == 0 {
		return json.Object{}
	} else if size == 1 {
		test, _ := qs.GetObject(0)
		return test
	} else {
		bl := json.Object{}
		if st.And {
			bl.AddArray("must", qs)
		} else {
			bl.AddArray("should", qs)
		}
		query.AddObject("bool", bl)
	}

	if st.Not {
		query = not(query)
	}

	return query
}
