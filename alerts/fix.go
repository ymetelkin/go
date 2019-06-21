package main

import (
	"strings"
)

const (
	runeSpace      rune = 32
	runeQuote      rune = 34
	runeLeftParen  rune = 40
	runeRightParen rune = 41
	runeColon      rune = 58
	runeA          rune = 65
	runeD          rune = 68
	runeN          rune = 78
	runeO          rune = 79
	runeR          rune = 82
	runeT          rune = 84
)

func fixQuery(q string) (string, bool) {
	var (
		i, min int
		qt, sp bool
		s      string
		sb     strings.Builder
	)

	s = strings.ReplaceAll(q, ",,)", ")")
	s = strings.ReplaceAll(s, ",)", ")")
	s = strings.ReplaceAll(s, "_all:", "")
	s = strings.ReplaceAll(s, "\"\"\"", "\"")
	s = strings.ReplaceAll(s, "\"\"", "\"")
	s = strings.ReplaceAll(s, ", NOT ", " AND NOT ")

	runes := []rune(s)
	size := len(runes)
	min = size - 4

	for i < size {
		r := runes[i]

		if qt {
			if r == runeQuote {
				qt = false
			}
		} else {
			switch r {
			case runeQuote:
				qt = true
				if sp {
					sb.WriteString(" OR ")
					sp = false
				}
			case runeSpace:
				sp = true
				i++
				continue
			case runeColon:
				if i < min && runes[i+1] == runeLeftParen {
					sb.WriteRune(runeColon)
					sb.WriteRune(runeLeftParen)
					s, idx := fixField(runes, i, size)
					i = idx
					sb.WriteString(s)
					i++
					continue
				}
			default:
				if sp {
					op, idx := checkOperator(runes, i, size)
					if op == "" {
						sb.WriteString(" OR ")
						sp = false
					} else {
						sb.WriteRune(runeSpace)
						sb.WriteString(op)
						sb.WriteRune(runeSpace)
						i = idx + 1
						sp = false
						continue
					}
				}
			}
		}

		sb.WriteRune(r)
		i++
	}

	qs := sb.String()
	qs = strings.ReplaceAll(qs, " OR +", " +")
	qs = strings.ReplaceAll(qs, " OR -", " -")
	//qs = strings.ReplaceAll(qs, " OR OR OR ", " OR ")
	//qs = strings.ReplaceAll(qs, " OR AND OR ", " AND ")
	qs = strings.ReplaceAll(qs, "NOT OR ", "NOT ")
	//qs = strings.ReplaceAll(qs, " OR NOT OR ", " AND NOT ")
	qs = strings.ReplaceAll(qs, ", OR ", " OR ")
	//qs = strings.ReplaceAll(qs, " OR  OR ", " OR ")
	qs = strings.ReplaceAll(qs, ", AND ", " AND ")

	return qs, qs != q
}

func fixField(runes []rune, i int, size int) (string, int) {
	var (
		qt, sp bool
		sb     strings.Builder
	)

	i++
	i++

	for i < size {
		r := runes[i]

		if qt {
			if r == runeQuote {
				qt = false
			}
		} else {
			switch r {
			case runeQuote:
				qt = true
				if sp {
					sb.WriteString(" OR ")
					sp = false
				}
			case runeSpace:
				sp = true
				i++
				continue
			case runeRightParen:
				sb.WriteRune(r)
				return sb.String(), i
			default:
				if sp {
					op, idx := checkOperator(runes, i, size)
					if op == "" {
						sb.WriteString(" OR ")
						sp = false
					} else {
						sb.WriteRune(runeSpace)
						sb.WriteString(op)
						sb.WriteRune(runeSpace)
						i = idx + 1
						sp = false
						continue
					}
				}
			}
		}

		sb.WriteRune(r)
		i++
	}

	return sb.String(), i
}

func checkOperator(runes []rune, i int, size int) (string, int) {
	var op string

	for i < size {
		r := runes[i]

		switch r {
		case runeA:
			if op == "" {
				op = "A"
			} else {
				return "", i
			}
		case runeD:
			if op == "AN" {
				op = "AND"
			} else {
				return "", i
			}
		case runeN:
			if op == "A" {
				op = "AN"
			} else if op == "" {
				op = "N"
			} else {
				return "", i
			}
		case runeO:
			if op == "" {
				op = "O"
			} else if op == "N" {
				op = "NO"
			} else if op == "T" {
				op = "TO"
			} else {
				return "", i
			}
		case runeR:
			if op == "O" {
				op = "OR"
			} else {
				return "", i
			}
		case runeT:
			if op == "NO" {
				op = "NOT"
			} else if op == "" {
				op = "T"
			} else {
				return "", i
			}
		case runeSpace:
			if op == "AND" || op == "OR" || op == "NOT" || op == "TO" {
				return op, i
			}
			return "", i
		default:
			return "", i
		}
		i++
	}

	return "", i
}
