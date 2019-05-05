package apql

import (
	"errors"
	"fmt"
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
}

type bools struct {
	Queries []query
	Ands    []int
	Nots    []int
	Bools   map[int]bools
}

func booleanize(tokens []token) (bools, error) {
	if tokens == nil {
		return bools{}, errors.New("No input tokens found")
	}

	size := len(tokens)
	if size == 0 {
		return bools{}, errors.New("No input tokens found")
	}

	var (
		txt, op string
	)

	qs := make([]query, 0)
	ands := make([]int, 0)
	nots := make([]int, 0)
	bls := make(map[int]bools)

	for _, tok := range tokens {
		switch tok.Type {
		case tokenText:
			if txt != "" {
				if op != "" {
					q := query{
						Field:    txt,
						Operator: op,
						Value:    tok.Text,
					}
					qs = append(qs, q)
					txt = ""
					op = ""
				} else {
					q := query{
						Field:    "_all",
						Operator: "=",
						Value:    txt,
					}
					qs = append(qs, q)
					txt = tok.Text
				}
			} else {
				txt = tok.Text
			}
		case tokenOperator:
			if op == ">" || op == "<" {
				if tok.Text == "=" {
					op += "="
				} else {
					return bools{}, fmt.Errorf("Expected '=', found '%s'", tok.Text)
				}
			} else if op != "" {
				return bools{}, fmt.Errorf("Expected '%s', found '%s%s'", op, op, tok.Text)
			} else {
				op = tok.Text
			}
		case tokenGroup:
			if txt != "" {
				if op == "HAS" || op == "HASANY" || op == "HASALL" {
					vals, err := getValues(tok.Tokens)
					if err != nil {
						return bools{}, err
					}
					q := query{
						Field:    txt,
						Operator: op,
						Values:   vals,
					}
					qs = append(qs, q)
					txt = ""
					op = ""
				} else if op != "" {
					return bools{}, fmt.Errorf("Expected 'HAS', found '%s'", op)
				} else {
					bs, err := booleanize(tok.Tokens)
					if err != nil {
						return bools{}, err
					}

					q := query{
						Field:    "_all",
						Operator: "=",
						Value:    txt,
					}
					qs = append(qs, q)
					bls[len(qs)] = bs
					txt = ""
				}
			} else {
				bs, err := booleanize(tok.Tokens)
				if err != nil {
					return bools{}, err
				}
				bls[len(qs)] = bs
			}
		case tokenBool:
			if tok.Text == "AND" {
				ands = append(ands, len(qs))
			} else if tok.Text == "NOT" {
				nots = append(nots, len(qs))
			}
		}
	}

	return bools{
		Queries: qs,
		Ands:    ands,
		Nots:    nots,
		Bools:   bls,
	}, nil
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
			vals = append(vals, tok.Text)
		} else {
			return nil, fmt.Errorf("Expected text, found '%s'", tok.Text)
		}
	}

	return vals, nil
}
