package apql

import (
	"errors"
	"fmt"
	"strings"
)

const (
	tokenText     int = 0
	tokenOperator int = 1
	tokenGroup    int = 2
	tokenBool     int = 3
)

const (
	tokenNull        rune = 0
	tokenBL          rune = 7
	tokenBS          rune = 8
	tokenHT          rune = 9
	tokenLF          rune = 10
	tokenVT          rune = 11
	tokenFF          rune = 12
	tokenCR          rune = 13
	tokenSpace       rune = 32
	tokenQuote       rune = 34
	tokenDollar      rune = 36
	tokenOpenParen   rune = 40
	tokenCloseParen  rune = 41
	tokenComma       rune = 44
	tokenMinus       rune = 45
	tokenPeriod      rune = 46
	token0           rune = 48
	token9           rune = 57
	tokenColon       rune = 58
	tokenLT          rune = 60
	tokenEQ          rune = 61
	tokenGT          rune = 62
	tokenQuestion    rune = 63
	tokenEUpper      rune = 69
	tokenLeftSquare  rune = 91
	tokenBackslash   rune = 92
	tokenRightSquare rune = 93
	tokenA           rune = 97
	tokenB           rune = 98
	tokenE           rune = 101
	tokenF           rune = 102
	tokenL           rune = 108
	tokenN           rune = 110
	tokenR           rune = 114
	tokenS           rune = 115
	tokenT           rune = 116
	tokenU           rune = 117
	tokenV           rune = 118
	tokenLeftCurly   rune = 123
	tokenRightCurly  rune = 125
)

type token struct {
	Type   int
	Text   string
	Tokens []token
}

func tokenize(s string) ([]token, error) {
	s = strings.Trim(s, " ")
	if s == "" {
		return nil, errors.New("Missing string input")
	}

	runes := []rune(s)
	size := len(runes)

	tokens, _, err := tokenizeRunes(runes, size, 0)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func tokenizeRunes(runes []rune, size int, i int) ([]token, int, error) {
	var (
		qt, sp bool
		sb     strings.Builder
	)

	tokens := make([]token, 0)

	for i < size {
		r := runes[i]

		if qt {
			if r == tokenQuote {
				qt = false
				sb.WriteRune(r)
				tokens, sb = addToken(tokens, sb)
				i++
				continue
			}
		} else {
			switch r {
			case tokenQuote:
				qt = true
				sp = false
				tokens, sb = addToken(tokens, sb)
			case tokenSpace:
				if !sp {
					sp = true
					tokens, sb = addToken(tokens, sb)
				}
				i++
				continue
			case tokenEQ:
				tokens, sb = addToken(tokens, sb)
				tok := token{
					Text: "=",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case tokenGT:
				tokens, sb = addToken(tokens, sb)
				tok := token{
					Text: ">",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case tokenLT:
				tokens, sb = addToken(tokens, sb)
				tok := token{
					Text: "<",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case tokenOpenParen:
				sp = false
				tokens, sb = addToken(tokens, sb)
				toks, idx, err := tokenizeRunes(runes, size, i+1)
				if err != nil {
					return nil, idx, err
				}
				if runes[idx] != tokenCloseParen {
					return nil, idx, fmt.Errorf("Expected ')', found '%v'", runes[idx])
				}
				if toks != nil && len(toks) > 0 {
					tok := token{
						Type:   tokenGroup,
						Tokens: toks,
					}
					tokens = append(tokens, tok)
				}
				i = idx + 1
				continue
			case tokenCloseParen:
				tokens, _ = addToken(tokens, sb)
				return tokens, i, nil
			default:
				if sp {
					sp = false
					tokens, sb = addToken(tokens, sb)
				}
			}
		}

		sb.WriteRune(r)
		i++
	}

	tokens, _ = addToken(tokens, sb)

	return tokens, i, nil
}

func addToken(tokens []token, sb strings.Builder) ([]token, strings.Builder) {
	if sb.Len() > 0 {
		text := sb.String()

		var tp int

		switch text {
		case "OR", "AND", "NOT":
			tp = tokenBool
		case "=", "HAS":
			tp = tokenOperator
		default:
			tp = tokenText
		}

		tok := token{
			Text: text,
			Type: tp,
		}
		tokens = append(tokens, tok)
		sb = strings.Builder{}
	}

	return tokens, sb
}
