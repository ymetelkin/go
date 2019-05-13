package apql

import (
	"errors"
	"fmt"
	"strings"
)

type set struct {
	And     bool
	Not     bool
	Queries []query
	Sets    []set
}

type query struct {
	Field    string
	Operator string
	Value    string
	Values   []string
	Phrase   bool
}

type bools struct {
	Queries []query
	Ands    []int
	Nots    []int
	Bools   map[int]bools
}

func booleanize(tokens []token) (set, error) {
	if tokens == nil || len(tokens) == 0 {
		return set{}, errors.New("No input tokens found")
	}

	var (
		txt, op string
		idx     int
	)

	queries := make(map[int]query)
	sets := make(map[int]set)
	ors := make([]int, 0)
	ands := make([]int, 0)
	nots := make([]int, 0)

	for _, tok := range tokens {
		switch tok.Type {
		case tokenText:
			if txt != "" {
				if op != "" {
					q := query{
						Field:    txt,
						Operator: op,
						Value:    tok.Text,
						Phrase:   tok.Phrase,
					}
					queries[idx] = q
					idx++
					txt = ""
					op = ""
				} else {
					q := query{
						Field:    "_all",
						Operator: "=",
						Value:    txt,
						Phrase:   tok.Phrase,
					}
					queries[idx] = q
					idx++
					txt = tok.Text
				}
			} else {
				txt = tok.Text
			}
		case tokenOperator:
			if op == "!" || op == ">" || op == "<" {
				if tok.Text == "=" {
					op += "="
				} else {
					return set{}, fmt.Errorf("Expected '=', found '%s'", tok.Text)
				}
			} else if op != "" {
				return set{}, fmt.Errorf("Expected '%s', found '%s%s'", op, op, tok.Text)
			} else {
				op = tok.Text
			}
		case tokenGroup:
			if txt != "" {
				if op == "HAS" || op == "HASANY" || op == "HASALL" {
					vals, err := getValues(tok.Tokens)
					if err != nil {
						return set{}, err
					}
					q := query{
						Field:    txt,
						Operator: op,
						Values:   vals,
					}
					queries[idx] = q
					idx++
					txt = ""
					op = ""
				} else if op != "" {
					return set{}, fmt.Errorf("Expected 'HAS', found '%s'", op)
				} else {
					st, err := booleanize(tok.Tokens)
					if err != nil {
						return set{}, err
					}

					q := query{
						Field:    "_all",
						Operator: "=",
						Value:    txt,
					}
					queries[idx] = q
					idx++
					sets[idx] = st
					idx++
					txt = ""
				}
			} else {
				st, err := booleanize(tok.Tokens)
				if err != nil {
					return set{}, err
				}
				sets[idx] = st
				idx++
			}
		case tokenBool:
			if op != "" {
				return set{}, fmt.Errorf("Expected '%s', found '%s %s'", tok.Text, op, tok.Text)
			}
			if txt != "" {
				q := query{
					Field:    "_all",
					Operator: "=",
					Value:    txt,
				}
				queries[idx] = q
				idx++
				txt = ""
			}

			size := len(queries) + len(sets)
			asz := len(ands)
			osz := len(ors)
			nsz := len(nots)

			if tok.Text == "AND" {
				if asz > 0 && ands[asz-1] == size {
					return set{}, errors.New("Expected 'AND', found 'AND AND")
				}
				if osz > 0 && ors[osz-1] == size {
					return set{}, errors.New("Expected 'AND', found 'OR AND")
				}
				if nsz > 0 && nots[nsz-1] == size {
					return set{}, errors.New("Expected 'AND', found 'NOT AND")
				}
				ands = append(ands, size)
			} else if tok.Text == "NOT" {
				if nsz > 0 && nots[nsz-1] == size {
					return set{}, errors.New("Expected 'NOT', found 'NOT NOT")
				}
				nots = append(nots, size)
			} else if tok.Text == "OR" {
				if asz > 0 && ands[asz-1] == size {
					return set{}, errors.New("Expected 'OR', found 'AND OR")
				}
				if osz > 0 && ors[osz-1] == size {
					return set{}, errors.New("Expected 'OR', found 'OR OR")
				}
				if nsz > 0 && nots[nsz-1] == size {
					return set{}, errors.New("Expected 'OR', found 'NOT OR")
				}
				ors = append(ors, size)
			}
		}
	}

	ret := set{}

	if len(nots) > 0 {
		for i := range nots {
			st := set{Not: true}
			if q, ok := queries[i]; ok {
				st.Queries = []query{q}
				delete(queries, i)
			} else if s, ok := sets[i]; ok {
				st.Sets = []set{s}
				delete(sets, i)
			} else {
				return set{}, errors.New("Expected expression, found 'NOT'")
			}

			if ret.Sets == nil {
				ret.Sets = []set{st}
			} else {
				ret.Sets = append(ret.Sets, st)
			}
		}
	}

	if len(ands) > 0 {
		var (
			last int
			st   set
		)

		for _, i := range ands {
			prev := last > 0 && last == i-1

			if !prev {
				if st.And {
					if ret.Sets == nil {
						ret.Sets = []set{st}
					} else {
						ret.Sets = append(ret.Sets, st)
					}
				}
				st = set{And: true}
			}

			if !prev {
				if q, ok := queries[i-1]; ok {
					if st.Queries == nil {
						st.Queries = []query{q}
					} else {
						st.Queries = append(st.Queries, q)
					}
					delete(queries, i-1)
				} else if s, ok := sets[i-1]; ok {
					if st.Sets == nil {
						st.Sets = []set{s}
					} else {
						st.Sets = append(st.Sets, s)
					}
					delete(sets, i-1)
				} else {
					return set{}, errors.New("Expected expression, found 'AND'")
				}
			}

			if q, ok := queries[i]; ok {
				st.Queries = append(st.Queries, q)
				delete(queries, i)
				last = i
			} else if s, ok := sets[i]; ok {
				st.Sets = append(st.Sets, s)
				delete(sets, i)
				last = i
			} else {
				return set{}, errors.New("Expected expression, found 'AND'")
			}
		}

		if st.And {
			if ret.Sets == nil {
				ret.Sets = []set{st}
			} else {
				ret.Sets = append(ret.Sets, st)
			}
		}
	}

	if len(queries) > 0 {
		ret.Queries = make([]query, 0)
		for _, q := range queries {
			ret.Queries = append(ret.Queries, q)
		}
	}

	if len(sets) > 0 {
		ret.Sets = make([]set, 0)
		for _, s := range sets {
			ret.Sets = append(ret.Sets, s)
		}
	}

	if (ret.Queries == nil || len(ret.Queries) == 0) && ret.Sets != nil && len(ret.Sets) == 1 {
		return ret.Sets[0], nil
	}

	return ret, nil
}

func getValues(tokens []token) ([]string, error) {
	if tokens == nil {
		return nil, errors.New("No input tokens found")
	}

	size := len(tokens)
	if size == 0 {
		return nil, errors.New("No input tokens found")
	}

	vals := make([]string, 0)

	for _, tok := range tokens {
		if tok.Type == tokenText {
			txt := tok.Text
			if tok.Phrase {
				txt = fmt.Sprintf("\"%s\"", txt)
			}
			vals = append(vals, txt)
		} else {
			return nil, fmt.Errorf("Expected text, found '%s'", tok.Text)
		}
	}

	return vals, nil
}

func (q *query) string() string {
	var sb strings.Builder
	if !strings.HasPrefix(q.Field, "@") {
		sb.WriteString("@")
	}
	sb.WriteString(q.Field)
	sb.WriteString(q.Operator)

	if q.Value == "" {
		sb.WriteString("(")
		for i, v := range q.Values {
			if i > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(v)
		}
		sb.WriteString(")")
	} else {
		sb.WriteString(q.Value)
	}

	return sb.String()
}

func (st *set) string() string {
	var (
		sb strings.Builder
		sp bool
	)

	if st.Not {
		sb.WriteString("NOT (")

		if st.Queries != nil && len(st.Queries) > 0 {
			sb.WriteString(st.Queries[0].string())
		} else if st.Sets != nil && len(st.Sets) > 0 {
			sb.WriteString(st.Sets[0].string())
		}

		sb.WriteString(")")
	} else {
		b := " OR "
		if st.And {
			b = " AND "
		}

		if st.Queries != nil && len(st.Queries) > 0 {
			for _, q := range st.Queries {
				if sp {
					sb.WriteString(b)
				} else {
					sp = true
				}

				sb.WriteString(q.string())
			}
		}

		if st.Sets != nil && len(st.Sets) > 0 {
			for _, s := range st.Sets {
				if sp {
					sb.WriteString(b)
				} else {
					sp = true
				}

				sb.WriteString("(")
				sb.WriteString(s.string())
				sb.WriteString(")")
			}
		}
	}

	return sb.String()
}
