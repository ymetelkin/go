package query

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type jsParser struct {
	idx     int
	last    int
	buffer  []byte
	aliases map[string][]string
}

type qsToken struct {
	Type  int //0 query, 1 OR, 2 AND, 3 NOT
	Query qsString
}

func newParser(aliases map[string][]string) jsParser {
	return jsParser{
		aliases: aliases,
	}
}

func (p *jsParser) Parse(s string) (qsQuery, error) {
	p.idx = 0
	p.buffer = []byte(s)
	p.last = len(p.buffer) - 1
	return p.parse()
}

func (p *jsParser) parse() (qs qsQuery, err error) {
	var (
		sb           strings.Builder
		ok           = true
		tokens       []qsToken
		and, not, or bool
		sign         string
	)

	for {
		c, e := p.read()
		if e != nil {
			if e == io.EOF {
				if sb.Len() > 0 {
					q := qsField{
						Term: qsSingleTerm{
							Value: sb.String(),
						},
					}
					tokens = append(tokens, qsToken{Query: q})
				}
				break
			} else {
				err = e
				return
			}
		}

		if isWS(c) {
			if sb.Len() == 0 {
				continue
			}

			v := sb.String()
			switch v {
			case "AND":
				and = true
				tokens = append(tokens, qsToken{Type: 2})
			case "OR":
				or = true
				tokens = append(tokens, qsToken{Type: 1})
			case "NOT":
				not = true
				tokens = append(tokens, qsToken{Type: 3})
			default:
				q := qsField{
					Term: qsSingleTerm{
						Value: v,
					},
				}
				tokens = append(tokens, qsToken{Query: q})
			}
			sb.Reset()
			continue
		}

		switch c {
		case '(':
			q, e := p.parse()
			if e != nil {
				err = e
				return
			}
			tokens = append(tokens, qsToken{Query: q})
		case ')':
			if sb.Len() > 0 {
				q := qsField{
					Term: qsSingleTerm{
						Value: sb.String(),
					},
				}
				tokens = append(tokens, qsToken{Query: q})
			}
			goto TOKENS
		case '"':
			p.idx--
			term, e := p.parseTerm()
			if e != nil {
				err = e
				return
			}
			q := qsField{
				Term: term,
			}
			tokens = append(tokens, qsToken{Query: q})
			c, e := p.read()
			if e != nil {
				if e == io.EOF {
					break
				} else {
					err = e
					return
				}
			}
			if !isWS(c) {
				err = fmt.Errorf(`Expected space or EOF, found "%c"`, c)
				return
			}
		case ':':
			if !ok {
				err = fmt.Errorf("Invalid field characters: [ %s ]", sb.String())
				return
			}
			field := sb.String()
			if strings.HasSuffix(field, ".") {
				err = fmt.Errorf("Invalid field characters: [ %s ]", sb.String())
				return
			}

			if sign != "" {
				field = field[1:len(field)]
			}

			term, e := p.parseTerm()
			if e != nil {
				err = e
				return
			}

			var (
				q qsString
				k bool
			)

			if p.aliases != nil {
				fs, ok := p.aliases[field]
				if ok {
					if len(fs) > 1 {
						qq := qsQuery{
							Sign:     sign,
							Operator: "OR",
						}
						for _, f := range fs {
							q = qsField{
								Field: f,
								Term:  term,
							}
							qq.Queries = append(qq.Queries, q)
						}
						q = qq
						k = true
					} else {
						field = fs[0]
					}
				}
			}
			if !k {
				q = qsField{
					Sign:  sign,
					Field: field,
					Term:  term,
				}
			}
			tokens = append(tokens, qsToken{Query: q})
			sb.Reset()
			sign = ""
		default:
			sb.WriteByte(c)
			if sb.Len() == 1 {
				if !isAlpha(c) && c != '_' {
					if c == '+' || c == '-' {
						sign = string([]byte{c})
					} else {
						ok = false
					}
				}
			} else {
				if !isAlpha(c) && c != '.' && c != '_' && c != '-' && !isDigit(c) {
					if c == '\\' {
						if p.buffer[p.idx-2] == '\\' {
							ok = false
						}
					} else if c == '*' {
						if p.buffer[p.idx-2] != '\\' {
							ok = false
						}
					} else {
						ok = false
					}
				}
			}
		}
	}

TOKENS:
	last := len(tokens) - 1
	if last == -1 {
		return
	}

	if not {
		var i int
		for i <= last {
			token := tokens[i]

			if token.Type == 3 { //NOT
				if last == 0 {
					err = errors.New("NOT is not valid query")
					break
				}
				if i == last {
					err = errors.New("Missing NOT value")
					break
				}
				i++
				tok := tokens[i]
				if tok.Type > 0 {
					err = fmt.Errorf(`"NOT %s" is not expected`, tok.String())
					break
				}
				tokens[i-1] = qsToken{
					Query: qsQuery{
						Operator: "NOT",
						Queries:  []qsString{tok.Query},
					},
				}
				tokens[i] = qsToken{
					Type: -1,
				}
			}
			i++
		}
	}

	if and {
		var (
			i       int
			first   = -1
			queries []qsString
		)

		for i <= last {
			token := tokens[i]

			if token.Type == -1 {
				i++
				continue
			}

			if token.Type == 2 { //AND
				if i == 0 || first == -1 {
					err = errors.New("Query cannot start with AND")
					break
				}
				if i == last {
					err = errors.New("Query cannot end with AND")
					break
				}

				var tok = qsToken{
					Type: -1,
				}
				i++
				for i <= last {
					tok = tokens[i]
					if tok.Type >= 0 {
						break
					}
					i++
				}
				if tok.Type == -1 {
					err = errors.New("Query cannot end with AND")
					break
				}
				if tok.Type > 0 {
					err = fmt.Errorf(`"AND %s" is not expected`, tok.String())
					break
				}

				if len(queries) == 0 {
					queries = append(queries, tokens[first].Query)
				}
				queries = append(queries, tok.Query)

				tokens[i-1] = qsToken{
					Type: -1,
				}
				tokens[i] = qsToken{
					Type: -1,
				}
			} else if len(queries) > 0 {
				tokens[first] = qsToken{
					Query: qsQuery{
						Operator: "AND",
						Queries:  queries,
					},
				}
				queries = []qsString{}
				first = -1
			} else if token.Type == 0 {
				first = i
			} else {
				first = -1
			}
			i++
		}

		if len(queries) > 0 {
			tokens[first] = qsToken{
				Query: qsQuery{
					Operator: "AND",
					Queries:  queries,
				},
			}
		}
	}

	if or {
		var (
			i       int
			first   = -1
			queries []qsString
		)

		for i <= last {
			token := tokens[i]

			if token.Type == -1 {
				i++
				continue
			}

			if token.Type == 1 { //OR
				if i == 0 || first == -1 {
					err = errors.New("Query cannot start with OR")
					break
				}
				if i == last {
					err = errors.New("Query cannot end with OR")
					break
				}

				var tok = qsToken{
					Type: -1,
				}
				i++
				for i <= last {
					tok = tokens[i]
					if tok.Type >= 0 {
						break
					}
					i++
				}
				if tok.Type == -1 {
					err = errors.New("Query cannot end with OR")
					break
				}
				if tok.Type > 0 {
					err = fmt.Errorf(`"OR %s" is not expected`, tok.String())
					break
				}

				if len(queries) == 0 {
					queries = append(queries, tokens[first].Query)
				}
				queries = append(queries, tok.Query)

				tokens[i-1] = qsToken{
					Type: -1,
				}
				tokens[i] = qsToken{
					Type: -1,
				}
			} else if len(queries) > 0 {
				tokens[first] = qsToken{
					Query: qsQuery{
						Operator: "OR",
						Queries:  queries,
					},
				}
				queries = []qsString{}
				first = -1
			} else if token.Type == 0 {
				first = i
			} else {
				first = -1
			}
			i++
		}

		if len(queries) > 0 {
			tokens[first] = qsToken{
				Query: qsQuery{
					Operator: "OR",
					Queries:  queries,
				},
			}
		}
	}

	if and || or || not {
		var (
			i       int
			queries []qsString
		)

		for i <= last {
			token := tokens[i]
			if token.Type == 0 { //Query
				queries = append(queries, token.Query)
			}
			i++
		}

		if len(queries) == 1 {
			qs, _ = tokens[0].Query.(qsQuery)
		} else {
			qs.Queries = queries
		}
	} else {
		for _, token := range tokens {
			qs.Queries = append(qs.Queries, token.Query)
		}
	}

	return
}

func (p *jsParser) read() (byte, error) {
	if p.idx > p.last {
		return 0, io.EOF
	}

	b := p.buffer[p.idx]
	p.idx++
	return b, nil
}

func (p *jsParser) parseTerm() (term qsString, err error) {
	var (
		sb strings.Builder
		pt qsPhraseTerm
	)

	for {
		c, e := p.read()
		if e != nil {
			if e == io.EOF {
				if pt.Value == "" {
					term = qsSingleTerm{
						Value: sb.String(),
					}
				} else {
					term = pt
				}

			} else {
				err = e
			}
			return
		}

		if isWS(c) {
			if sb.Len() > 0 {
				term = qsSingleTerm{
					Value: sb.String(),
				}
			} else if pt.Value != "" {
				term = pt
			} else {
				err = errors.New("Expected field value, found whitespace")
			}
			return
		}

		switch c {
		case '"':
			s, e := p.parsePhrase()
			if e != nil {
				err = e
			} else {
				pt = qsPhraseTerm{
					Value: s,
				}
			}
		case '(':
			return p.parseGroup()
		case '[', '{':
			return p.parseRange(c)
		case '>', '<':
			return p.parseSignRange(c)
		case '~':
			n, e := p.parseNumber(false)
			if e != nil {
				err = e
			} else if sb.Len() > 0 {
				var fuzzy int
				if n == "" {
					fuzzy = 2
				} else {
					fuzzy, _ = strconv.Atoi(n)
				}

				term = qsSingleTerm{
					Value: sb.String(),
					Fuzzy: fuzzy,
				}
			} else if pt.Value != "" {
				pt.Proximity, _ = strconv.Atoi(n)
				term = pt
			}
			return
		case '^':
			n, e := p.parseNumber(true)
			if e != nil {
				err = e
			} else if n == "" {
				err = errors.New("Expected boost value")
			} else {
				boost, _ := strconv.ParseFloat(n, 32)
				if boost == 0 {
					err = errors.New("Expected boost value")
				} else if sb.Len() > 0 {
					term = qsSingleTerm{
						Value: sb.String(),
						Boost: float32(boost),
					}
				} else if pt.Value != "" {
					pt.Boost = float32(boost)
					term = pt
				}
			}
			return
		default:
			sb.WriteByte(c)
		}
	}
}

func (p *jsParser) parseRange(left byte) (term qsRangeTerm, err error) {
	var (
		sb strings.Builder
		to bool
	)

	for {
		c, e := p.read()
		if e != nil {
			if e == io.EOF {
				err = errors.New(`Expected "]" or "}", found EOF`)
			} else {
				err = e
			}
			return
		}

		if isWS(c) {
			if sb.Len() > 0 {
				if term.Left == "" {
					term.Left = sb.String()
					sb.Reset()
				} else {
					term.Right = sb.String()
				}
			} else {
				continue
			}
		} else {
			switch c {
			case 'T':
				if term.Left != "" && !to {
					if p.idx < p.last-2 && p.buffer[p.idx] == 'O' && isWS(p.buffer[p.idx+1]) {
						to = true
						p.idx++
						p.idx++
						continue
					} else {
						err = fmt.Errorf(`Expected " TO ", found " T%c%c"`, p.buffer[p.idx], p.buffer[p.idx+1])
						return
					}
				} else {
					sb.WriteByte(c)
				}
			case '"':
				s, e := p.parsePhrase()
				if e != nil {
					err = e
					return
				}
				if term.Left == "" {
					term.Left = s
				} else {
					term.Right = s
				}
			case ']', '}':
				if to {
					term.Right = sb.String()
					term.IncludeLeft = left == '['
					term.IncludeRight = c == ']'
					return
				}

			default:
				sb.WriteByte(c)
			}
		}
	}
}

func (p *jsParser) parseSignRange(sign byte) (term qsRangeTerm, err error) {
	var (
		sb    strings.Builder
		value string
		inc   bool
	)

	for {
		c, e := p.read()
		if e != nil {
			if e == io.EOF {
				if sb.Len() == 0 {
					err = errors.New(`Expected range value, found EOF`)
				} else {
					value = sb.String()
					break
				}
			} else {
				err = e
			}
			return
		}

		if isWS(c) {
			if sb.Len() > 0 {
				value = sb.String()
				break
			} else {
				continue
			}
		}

		switch c {
		case '=':
			if sb.Len() > 0 {
				err = fmt.Errorf(`"=" is not expected at [ %d ] position`, p.idx-1)
				return
			}
			inc = true
		case '"':
			s, e := p.parsePhrase()
			if e != nil {
				err = e
				return
			}
			value = s
		default:
			sb.WriteByte(c)
		}
	}

	if sign == '>' {
		term.Left = value
		term.IncludeLeft = inc
	} else {
		term.Right = value
		term.IncludeRight = inc
	}

	return
}

func (p *jsParser) parsePhrase() (string, error) {
	var sb strings.Builder

	for {
		c, err := p.read()
		if err != nil {
			if err == io.EOF {
				return "", errors.New(`Expected ", found EOF`)
			}
			return "", err
		}

		if c == '"' && p.buffer[p.idx-2] != '\\' {
			return sb.String(), nil
		}
		sb.WriteByte(c)
	}
}

func (p *jsParser) parseGroup() (term qsGroupTerm, err error) {
	var sb strings.Builder

	for {
		c, e := p.read()
		if e != nil {
			if e == io.EOF {
				err = errors.New("Expected ), found EOF")
			} else {
				err = e
			}
			return
		}

		if c == ')' {
			if sb.Len() > 0 {
				term.Values = append(term.Values, sb.String())
			}
			return
		}

		if isWS(c) {
			if sb.Len() > 0 {
				term.Values = append(term.Values, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteByte(c)
		}
	}
}

func (p *jsParser) parseNumber(float bool) (string, error) {
	var (
		sb        strings.Builder
		zero, dot bool
	)

	for {
		c, err := p.read()
		if err != nil {
			if err == io.EOF {
				return "", nil
			}
			return "", err
		}

		if isWS(c) {
			return sb.String(), nil
		}

		if isDigit(c) {
			z := c == '0' && sb.Len() == 0
			if z && !float {
				return "", errors.New(`Expected integer, found leading "0"`)
			}
			if z && zero {
				return "", errors.New(`"00" is not expected`)
			}
			zero = z
			sb.WriteByte(c)
		} else if c == '.' {
			if !float {
				return "", errors.New(`Expected integer, found decimal "."`)
			}
			if dot {
				return "", errors.New(`Multiple "." are not expected`)
			}
			sb.WriteByte(c)
		} else {
			return "", fmt.Errorf(`"%c" at [ %d ] in not a digit`, c, p.idx-1)
		}
	}
}

func (tok *qsToken) String() string {
	switch tok.Type {
	case 1:
		return "OR"
	case 2:
		return "AND"
	case 3:
		return "NOT"
	}
	return tok.Query.String()
}

func isWS(b byte) bool {
	return b == ' ' || b == '\n' || b == '\r' || b == '\t'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
