package query

import (
	"fmt"
	"strings"
)

type qsString interface {
	String() string
}

type qsQuery struct {
	Operator string
	Queries  []qsString
	Sign     string
}

type qsField struct {
	Field string
	Term  qsString
	Sign  string
}

type qsSingleTerm struct {
	Value string
	Fuzzy int
	Boost float32
}

type qsGroupTerm struct {
	Values []string
}

type qsPhraseTerm struct {
	Value     string
	Proximity int
	Boost     float32
}

type qsRangeTerm struct {
	Left         string
	Right        string
	IncludeLeft  bool
	IncludeRight bool
}

func (q qsQuery) String() string {
	if len(q.Queries) == 0 {
		return ""
	}

	var (
		op  = q.Operator != ""
		not = q.Operator == "NOT"
	)

	if len(q.Queries) == 1 {
		s := q.Queries[0].String()
		if not {
			return "NOT " + s
		}
		return s
	}

	var sb strings.Builder

	for i, s := range q.Queries {
		if i == 0 {
			if not {
				sb.WriteString("NOT ")
			}
		} else {
			sb.WriteByte(' ')
			if not {
				sb.WriteString("AND NOT ")
			} else if op {
				sb.WriteString(q.Operator)
				sb.WriteByte(' ')
			}
		}

		if q.Sign != "" {
			sb.WriteString(q.Sign)
		}

		qs, ok := s.(qsQuery)
		if ok {
			if len(qs.Queries) > 1 {
				if qs.Sign != "" {
					sb.WriteString(qs.Sign)
					qs.Sign = ""
				}
				sb.WriteByte('(')
				sb.WriteString(qs.String())
				sb.WriteByte(')')
			} else {
				sb.WriteString(qs.String())
			}
		} else {
			sb.WriteString(s.String())
		}
	}

	return sb.String()
}

func (f qsField) String() string {
	var sb strings.Builder
	sb.WriteString(f.Sign)
	if f.Field != "" {
		sb.WriteString(f.Field)
		sb.WriteByte(':')
	}
	sb.WriteString(f.Term.String())
	return sb.String()
}

func (st qsSingleTerm) String() string {
	v := strings.TrimSpace(st.Value)
	if v == "" {
		return ""
	}

	if st.Boost > 0 {
		test := int(st.Boost)
		if st.Boost == float32(test) {
			return fmt.Sprintf("%s^%d", v, test)
		}
		return strings.TrimRight(fmt.Sprintf("%s^%f", v, st.Boost), "0")
	} else if st.Fuzzy > 0 {
		if st.Fuzzy == 2 {
			return v + "~"
		}
		return strings.TrimRight(fmt.Sprintf("%s~%d", v, st.Fuzzy), "0")
	}

	return v
}

func (gt qsGroupTerm) String() string {
	if len(gt.Values) == 0 {
		return ""
	}
	return fmt.Sprintf("(%s)", strings.Join(gt.Values, " "))
}

func (pt qsPhraseTerm) String() string {
	v := strings.TrimSpace(pt.Value)
	if v == "" {
		return ""
	}

	if pt.Boost > 0 {
		test := int(pt.Boost)
		if pt.Boost == float32(test) {
			return fmt.Sprintf(`"%s"^%d`, v, test)
		}
		return strings.TrimRight(fmt.Sprintf(`"%s"^%f`, v, pt.Boost), "0")
	} else if pt.Proximity > 0 {
		return fmt.Sprintf(`"%s"~%d`, v, pt.Proximity)
	}

	return fmt.Sprintf(`"%s"`, v)
}

func (rt qsRangeTerm) String() string {
	var (
		l  = strings.TrimSpace(rt.Left)
		r  = strings.TrimSpace(rt.Right)
		lo = "["
		ro = "]"
	)

	if r == "" {
		if l == "" {
			return ""
		}
		if !rt.IncludeLeft {
			lo = "{"
		}
		return fmt.Sprintf("%s%s TO *%s", lo, l, ro)
	}

	if !rt.IncludeRight {
		ro = "}"
	}

	if l == "" {
		return fmt.Sprintf("%s* TO %s%s", lo, r, ro)
	}
	if !rt.IncludeLeft {
		lo = "{"
	}
	return fmt.Sprintf("%s%s TO %s%s", lo, l, r, ro)
}
