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
	runeSpace      rune = 32
	runeNE         rune = 33
	runeQuote      rune = 34
	runeOpenParen  rune = 40
	runeCloseParen rune = 41
	runeLT         rune = 60
	runeEQ         rune = 61
	runeGT         rune = 62
)

type token struct {
	Type   int
	Text   string
	Phrase bool
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
			if r == runeQuote {
				qt = false
				tokens, sb = addToken(tokens, sb, true)
				i++
				continue
			}
		} else {
			switch r {
			case runeQuote:
				qt = true
				sp = false
				tokens, sb = addToken(tokens, sb, false)
				i++
				continue
			case runeSpace:
				if !sp {
					sp = true
					tokens, sb = addToken(tokens, sb, false)
				}
				i++
				continue
			case runeEQ:
				tokens, sb = addToken(tokens, sb, false)
				tok := token{
					Text: "=",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case runeNE:
				tokens, sb = addToken(tokens, sb, false)
				tok := token{
					Text: "!",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case runeGT:
				tokens, sb = addToken(tokens, sb, false)
				tok := token{
					Text: ">",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case runeLT:
				tokens, sb = addToken(tokens, sb, false)
				tok := token{
					Text: "<",
					Type: tokenOperator,
				}
				tokens = append(tokens, tok)
				i++
				continue
			case runeOpenParen:
				sp = false
				tokens, sb = addToken(tokens, sb, false)
				toks, idx, err := tokenizeRunes(runes, size, i+1)
				if err != nil {
					return nil, idx, err
				}
				if runes[idx] != runeCloseParen {
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
			case runeCloseParen:
				tokens, _ = addToken(tokens, sb, false)
				return tokens, i, nil
			default:
				if sp {
					sp = false
					tokens, sb = addToken(tokens, sb, false)
				}
			}
		}

		sb.WriteRune(r)
		i++
	}

	tokens, _ = addToken(tokens, sb, false)

	return tokens, i, nil
}

func addToken(tokens []token, sb strings.Builder, phrase bool) ([]token, strings.Builder) {
	if sb.Len() > 0 {
		text := sb.String()

		var tp int

		switch text {
		case "OR", "AND", "NOT":
			tp = tokenBool
		case "or", "and", "not":
			tp = tokenBool
			text = strings.ToUpper(text)
		case "=", "HAS":
			tp = tokenOperator
		default:
			tp = tokenText
		}

		tok := token{
			Text:   text,
			Type:   tp,
			Phrase: phrase,
		}
		tokens = append(tokens, tok)
		sb = strings.Builder{}
	}

	return tokens, sb
}
