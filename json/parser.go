package json

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	runeNull        rune = 0
	runeBL          rune = 7
	runeBS          rune = 8
	runeHT          rune = 9
	runeLF          rune = 10
	runeVT          rune = 11
	runeFF          rune = 12
	runeCR          rune = 13
	runeSpace       rune = 32
	runeQuote       rune = 34
	runeDollar      rune = 36
	runeComma       rune = 44
	runeMinus       rune = 45
	runePeriod      rune = 46
	rune0           rune = 48
	rune9           rune = 57
	runeColon       rune = 58
	runeQuestion    rune = 63
	runeEUpper      rune = 69
	runeLeftSquare  rune = 91
	runeBackslash   rune = 92
	runeRightSquare rune = 93
	runeA           rune = 97
	runeB           rune = 98
	runeE           rune = 101
	runeF           rune = 102
	runeL           rune = 108
	runeN           rune = 110
	runeR           rune = 114
	runeS           rune = 115
	runeT           rune = 116
	runeU           rune = 117
	runeV           rune = 118
	runeLeftCurly   rune = 123
	runeRightCurly  rune = 125
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
	runes := []rune(s)
	size := len(runes)
	r, i := skipWhitespace(runes, size, 0)

	if size < 2 {
		return value{}, errors.New("Invalid string input")
	}

<<<<<<< HEAD
	if r == tokenLeftCurly {
		jo, _, err := parseObject(runes, size, i, parameterize)
=======
	if runes[0] == runeLeftCurly {
		jo, _, err := parseObject(runes, size, 0, parameterize)
>>>>>>> 5f47947789048c5e033d95409fb25ea7dbbfa033

		if err != nil {
			return value{}, err
		}
		return newObject(jo), nil
<<<<<<< HEAD
	} else if runes[0] == tokenLeftSquare {
		ja, _, err := parseArray(runes, size, i, false)
=======
	} else if runes[0] == runeLeftSquare {
		ja, _, err := parseArray(runes, size, 0, false)
>>>>>>> 5f47947789048c5e033d95409fb25ea7dbbfa033

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

	if r == runeQuote {
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
	} else if r == runeRightCurly {
		return r, index, nil
	} else {
		return r, index, fmt.Errorf("Expected '\"', found '%c'", r)
	}
}

func parsePropertyName(runes []rune, size int, index int) (string, int, error) {
	index++

	r := runeNull
	start := index

	for index < size {
		r = runes[index]

		if r == runeQuote {
			end := index

			index++
			r, index = skipWhitespace(runes, size, index)
			if r == runeColon {
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

	if r == runeQuote {
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

			if r == runeBackslash {
				index++
				if index < size {
					test := runes[index]
					if test == runeR {
						sb.WriteRune(runeCR)
					} else if test == runeN {
						sb.WriteRune(runeLF)
					} else if test == runeT {
						sb.WriteRune(runeHT)
					} else if test == runeB {
						sb.WriteRune(runeBS)
					} else if test == runeF {
						sb.WriteRune(runeFF)
					} else if test == runeA {
						sb.WriteRune(runeBL)
					} else if test == runeV {
						sb.WriteRune(runeVT)
					} else if test == runeQuote {
						sb.WriteRune(runeQuote)
					}
				}
			} else if r == runeQuote {
				return newString(sb.String()), index, ParameterizedString{}, nil
			} else {
				sb.WriteRune(r)
			}

			index++
		}
		return value{}, index, ParameterizedString{}, nil
	}

	if r == runeLeftCurly {
		jo, index, err := parseObject(runes, size, index, parameterize)
		if err != nil {
			return value{}, index, ParameterizedString{}, err
		}
		return newObject(jo), index, ParameterizedString{}, nil
	}

	if r == runeLeftSquare {
		ja, index, err := parseArray(runes, size, index, parameterize)
		if err != nil {
			return value{}, index, ParameterizedString{}, err
		}
		return newArray(ja), index, ParameterizedString{}, nil
	}

	if r == runeT {
		index++
		if index < size && runes[index] == runeR {
			index++
			if index < size && runes[index] == runeU {
				index++
				if index < size && runes[index] == runeE {
					jv := newBool(true)
					jv.Text = "true"
					return jv, index, ParameterizedString{}, nil
				}
			}
		}
	} else if r == runeF {
		index++
		if index < size && runes[index] == runeA {
			index++
			if index < size && runes[index] == runeL {
				index++
				if index < size && runes[index] == runeS {
					index++
					if index < size && runes[index] == runeE {
						jv := newBool(false)
						jv.Text = "false"
						return jv, index, ParameterizedString{}, nil
					}
				}
			}
		}
	} else if r == runeN {
		index++
		if index < size && runes[index] == runeU {
			index++
			if index < size && runes[index] == runeL {
				index++
				if index < size && runes[index] == runeL {
					return newNull(), index, ParameterizedString{}, nil
				}
			}
		}
	} else if r > runeComma && r < runeColon {
		start := index
		floating := false

		index++
		r = runes[index]
		for r > runeMinus && r < runeColon {
			if r == runePeriod {
				floating = true
			}
			index++
			r = runes[index]
		}

		if r == runeE || r == runeEUpper { //Scientific
			floating = true
			index++
			r = runes[index] // skip sign + -
			for index < size {
				index++
				r = runes[index]
				if r < rune0 || r > rune9 { //skip digits
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

	for r == runeComma {
		r, index, err = addProperty(&jo, runes, size, index, parameterize)
		if err != nil {
			return Object{}, index, err
		}
	}

	if r != runeRightCurly {
		return Object{}, index, fmt.Errorf("Expected '}', found '%c'", r)
	}

	return jo, index, nil
}

func addValue(ja *Array, runes []rune, size int, index int, parameterize bool) (rune, int, error) {
	index++
	r, index := skipWhitespace(runes, size, index)

	if r == runeRightSquare {
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

	for r == runeComma {
		r, index, err = addValue(&ja, runes, size, index, parameterize)
		if err != nil {
			return Array{}, index, err
		}
	}

	if r != runeRightSquare {
		return Array{}, index, fmt.Errorf("Expected ']', found '%c'", r)
	}

	return ja, index, nil
}

func skipWhitespace(runes []rune, size int, index int) (rune, int) {
	for index < size {
		r := runes[index]

		if r == runeNull || r == runeSpace || r == runeLF || r == runeCR || r == runeHT || r == runeBS || r == runeFF || r == runeVT {
			index++
		} else {
			return r, index
		}
	}

	return runeNull, index
}
