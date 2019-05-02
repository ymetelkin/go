package json

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	tokenComma       rune = 44
	tokenMinus       rune = 45
	tokenPeriod      rune = 46
	token0           rune = 48
	token9           rune = 57
	tokenColon       rune = 58
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

//ParseJSONObject parses string to JSON object
func ParseJSONObject(s string) (Object, error) {
	return parseJSONObject(s, false)
}

//ParseJSONArray parses string to JSON array
func ParseJSONArray(s string) (Array, error) {
	jv, err := parseJSONValue(s, false)
	if err == nil {
		ja, err := jv.GetArray()
		if err == nil {
			return ja, nil
		}
	}

	return Array{}, err
}

func parseJSONValue(s string, parameterize bool) (value, error) {
	s = strings.Trim(s, " ")
	if s == "" {
		return value{}, errors.New("Missing string input")
	}

	runes := []rune(s)
	size := len(runes)

	if size < 2 {
		return value{}, errors.New("Invalid string input")
	}

	if runes[0] == tokenLeftCurly {
		jo, _, err := parseObject(runes, size, 0, parameterize)

		if err != nil {
			return value{}, err
		}
		return newObject(jo), nil
	} else if runes[0] == tokenLeftSquare {
		ja, _, err := parseArray(runes, size, 0, false)

		if err != nil {
			return value{}, err
		}
		return newArray(ja), nil
	} else {
		return value{}, errors.New("Invalid string input")
	}
}

func parseJSONObject(s string, parameterize bool) (Object, error) {
	jv, err := parseJSONValue(s, parameterize)
	if err == nil {
		jo, err := jv.GetObject()
		if err == nil {
			return jo, nil
		}
	}

	return Object{}, err
}

func addProperty(jo *Object, runes []rune, size int, index int, parameterize bool) (rune, int, error) {
	index++
	r, index := skipWhitespace(runes, size, index)

	if r == tokenQuote {
		var (
			name  string
			err   error
			pname ParameterizedString
		)

		if parameterize {
			pname, index, err = parsePropertyNameWithParameters(runes, size, index)
			name = pname.Value
		} else {
			name, index, err = parsePropertyName(runes, size, index)
		}
		if err != nil {
			return r, index, err
		}

		index++ //skip colon
		value, index, pvalue, err := parseValue(runes, size, index, parameterize)
		if err != nil {
			return r, index, err
		}

		if pvalue.Value != "" {
			if pvalue.IsParameterized {
				value = newParameterizedString(pvalue)
			} else {
				value = newString(pvalue.Value)
			}
		}

		if !value.IsEmpty() {
			if pname.IsParameterized || pvalue.IsParameterized {
				jo.AddWithParameters(pname, value)
			} else {
				jo.addValue(name, value)
			}
		}
		index++
		r, index = skipWhitespace(runes, size, index)
		return r, index, nil
	} else if r == tokenRightCurly {
		return r, index, nil
	} else {
		return r, index, fmt.Errorf("Expected '\"', found '%c'", r)
	}
}

func parsePropertyName(runes []rune, size int, index int) (string, int, error) {
	index++

	r := tokenNull
	start := index

	for index < size {
		r = runes[index]

		if r == tokenQuote {
			end := index

			index++
			r, index = skipWhitespace(runes, size, index)
			if r == tokenColon {
				return string(runes[start:end]), index, nil
			}
			return "", index, fmt.Errorf("Expected ':', found '%c'", r)
		}

		index++
	}

	return "", index, fmt.Errorf("Expected '\"', found '%c'", r)
}

func parseValue(runes []rune, size int, index int, parameterize bool) (value, int, ParameterizedString, error) {
	r, index := skipWhitespace(runes, size, index)

	if r == tokenQuote {
		index++

		if parameterize {
			ps, index, err := parseTextValueWithParameters(runes, size, index)
			if err != nil {
				return value{}, index, ps, err
			}
			return value{}, index, ps, nil
		}

		var sb strings.Builder

		for index < size {
			r := runes[index]

			if r == tokenBackslash {
				index++
				if index < size {
					test := runes[index]
					if test == tokenR {
						sb.WriteRune(tokenCR)
					} else if test == tokenN {
						sb.WriteRune(tokenLF)
					} else if test == tokenT {
						sb.WriteRune(tokenHT)
					} else if test == tokenB {
						sb.WriteRune(tokenBS)
					} else if test == tokenF {
						sb.WriteRune(tokenFF)
					} else if test == tokenA {
						sb.WriteRune(tokenBL)
					} else if test == tokenV {
						sb.WriteRune(tokenVT)
					} else if test == tokenQuote {
						sb.WriteRune(tokenQuote)
					}
				}
			} else if r == tokenQuote {
				return newString(sb.String()), index, ParameterizedString{}, nil
			} else {
				sb.WriteRune(r)
			}

			index++
		}
		return value{}, index, ParameterizedString{}, nil
	}

	if r == tokenLeftCurly {
		jo, index, err := parseObject(runes, size, index, parameterize)
		if err != nil {
			return value{}, index, ParameterizedString{}, err
		}
		return newObject(jo), index, ParameterizedString{}, nil
	}

	if r == tokenLeftSquare {
		ja, index, err := parseArray(runes, size, index, parameterize)
		if err != nil {
			return value{}, index, ParameterizedString{}, err
		}
		return newArray(ja), index, ParameterizedString{}, nil
	}

	if r == tokenT {
		index++
		if index < size && runes[index] == tokenR {
			index++
			if index < size && runes[index] == tokenU {
				index++
				if index < size && runes[index] == tokenE {
					jv := newBool(true)
					jv.Text = "true"
					return jv, index, ParameterizedString{}, nil
				}
			}
		}
	} else if r == tokenF {
		index++
		if index < size && runes[index] == tokenA {
			index++
			if index < size && runes[index] == tokenL {
				index++
				if index < size && runes[index] == tokenS {
					index++
					if index < size && runes[index] == tokenE {
						jv := newBool(false)
						jv.Text = "false"
						return jv, index, ParameterizedString{}, nil
					}
				}
			}
		}
	} else if r == tokenN {
		index++
		if index < size && runes[index] == tokenU {
			index++
			if index < size && runes[index] == tokenL {
				index++
				if index < size && runes[index] == tokenL {
					return newNull(), index, ParameterizedString{}, nil
				}
			}
		}
	} else if r > tokenComma && r < tokenColon {
		start := index
		floating := false

		index++
		r = runes[index]
		for r > tokenMinus && r < tokenColon {
			if r == tokenPeriod {
				floating = true
			}
			index++
			r = runes[index]
		}

		if r == tokenE || r == tokenEUpper { //Scientific
			floating = true
			index++
			r = runes[index] // skip sign + -
			for index < size {
				index++
				r = runes[index]
				if r < token0 || r > token9 { //skip digits
					break
				}
			}
		}

		s := string(runes[start:index])

		index--

		if floating {
			f, err := strconv.ParseFloat(s, 64)
			if err == nil {
				jv := newFloat(f)
				jv.Text = s
				return jv, index, ParameterizedString{}, nil
			}
		} else {
			i, err := strconv.Atoi(s)
			if err == nil {
				jv := newInt(i)
				jv.Text = s
				return jv, index, ParameterizedString{}, nil
			}
		}

		return value{}, index, ParameterizedString{}, fmt.Errorf("Expected number, found '%s'", s)
	}

	return value{}, index, ParameterizedString{}, fmt.Errorf("Unexpected character: '%c'", r)
}

func parseObject(runes []rune, size int, index int, parameterize bool) (Object, int, error) {
	jo := Object{}

	r, index, err := addProperty(&jo, runes, size, index, parameterize)
	if err != nil {
		return Object{}, index, err
	}

	for r == tokenComma {
		r, index, err = addProperty(&jo, runes, size, index, parameterize)
		if err != nil {
			return Object{}, index, err
		}
	}

	if r != tokenRightCurly {
		return Object{}, index, fmt.Errorf("Expected '}', found '%c'", r)
	}

	return jo, index, nil
}

func addValue(ja *Array, runes []rune, size int, index int, parameterize bool) (rune, int, error) {
	index++
	r, index := skipWhitespace(runes, size, index)

	if r == tokenRightSquare {
		return r, index, nil
	}

	value, index, ps, err := parseValue(runes, size, index, parameterize)
	if err != nil {
		return r, index, err
	}

	if ps.Value != "" {
		if ps.IsParameterized {
			value = newParameterizedString(ps)
		} else {
			value = newString(ps.Value)
		}
	}

	if !value.IsEmpty() {
		ja.addValue(value)
	}
	index++
	r, index = skipWhitespace(runes, size, index)
	return r, index, nil
}

func parseArray(runes []rune, size int, index int, parameterize bool) (Array, int, error) {
	ja := Array{}

	r, index, err := addValue(&ja, runes, size, index, parameterize)
	if err != nil {
		return Array{}, index, err
	}

	for r == tokenComma {
		r, index, err = addValue(&ja, runes, size, index, parameterize)
		if err != nil {
			return Array{}, index, err
		}
	}

	if r != tokenRightSquare {
		return Array{}, index, fmt.Errorf("Expected ']', found '%c'", r)
	}

	return ja, index, nil
}

func skipWhitespace(runes []rune, size int, index int) (rune, int) {
	for index < size {
		r := runes[index]

		if r == tokenNull || r == tokenSpace || r == tokenLF || r == tokenCR || r == tokenHT || r == tokenBS || r == tokenFF || r == tokenVT {
			index++
		} else {
			return r, index
		}
	}

	return tokenNull, index
}
