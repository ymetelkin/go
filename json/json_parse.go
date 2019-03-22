package json

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	TOKEN_NULL         rune = 0
	TOKEN_BL           rune = 7
	TOKEN_BS           rune = 8
	TOKEN_HT           rune = 9
	TOKEN_LF           rune = 10
	TOKEN_VT           rune = 11
	TOKEN_FF           rune = 12
	TOKEN_CR           rune = 13
	TOKEN_SPACE        rune = 32
	TOKEN_QUOTE        rune = 34
	TOKEN_DOLLAR       rune = 36
	TOKEN_COMMA        rune = 44
	TOKEN_MINUS        rune = 45
	TOKEN_PERIOD       rune = 46
	TOKEN_0            rune = 48
	TOKEN_9            rune = 57
	TOKEN_COLON        rune = 58
	TOKEN_QUESTION     rune = 63
	TOKEN_E_UPPER      rune = 69
	TOKEN_LEFT_SQUARE  rune = 91
	TOKEN_BACKSLASH    rune = 92
	TOKEN_RIGHT_SQUARE rune = 93
	TOKEN_A            rune = 97
	TOKEN_B            rune = 98
	TOKEN_E            rune = 101
	TOKEN_F            rune = 102
	TOKEN_L            rune = 108
	TOKEN_N            rune = 110
	TOKEN_R            rune = 114
	TOKEN_S            rune = 115
	TOKEN_T            rune = 116
	TOKEN_U            rune = 117
	TOKEN_V            rune = 118
	TOKEN_LEFT_CURLY   rune = 123
	TOKEN_RIGHT_CURLY  rune = 125
)

func Parse(s string, parameterize bool) (*JsonValue, error) {
	s = strings.Trim(s, " ")
	if s == "" {
		return nil, errors.New("Missing string input")
	}

	runes := []rune(s)
	size := len(runes)

	if size < 2 {
		return nil, errors.New("Invalid string input")
	}

	if runes[0] == TOKEN_LEFT_CURLY {
		jo, _, err := parseObject(runes, size, 0, parameterize)

		if err != nil {
			return nil, err
		}
		return &JsonValue{Value: *jo, Type: OBJECT}, nil
	} else if runes[0] == TOKEN_LEFT_SQUARE {
		ja, _, err := parseArray(runes, size, 0, false)

		if err != nil {
			return nil, err
		}
		return &JsonValue{Value: *ja, Type: ARRAY}, nil
	} else {
		return nil, errors.New("Invalid string input")
	}
}

func ParseJsonObject(s string) (*JsonObject, error) {
	return parseJsonObject(s, false)
}

func parseJsonObject(s string, parameterize bool) (*JsonObject, error) {
	jv, err := Parse(s, parameterize)
	if err != nil {
		return nil, err
	} else {
		jo, err := jv.GetObject()
		if err != nil {
			return nil, err
		} else {
			return jo, nil
		}
	}
}

func ParseJsonArray(s string) (*JsonArray, error) {
	jv, err := Parse(s, false)
	if err != nil {
		return nil, err
	} else {
		ja, err := jv.GetArray()
		if err != nil {
			return nil, err
		} else {
			return ja, nil
		}
	}
}

func addProperty(jo *JsonObject, runes []rune, size int, index int, parameterize bool) (rune, int, error) {
	index++
	r, index := skipWhitespace(runes, size, index)

	if r == TOKEN_QUOTE {
		var name string
		var err error
		var pname ParameterizedString

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
				value = &JsonValue{Value: pvalue, Type: PARAMETERIZED}
			} else {
				value = &JsonValue{Value: pvalue.Value, Type: STRING}
			}
		}

		if value != nil {
			if pname.IsParameterized || pvalue.IsParameterized {
				jo.AddWithParameters(pname, value)
			} else {
				jo.Add(name, *value)
			}
		}
		index++
		r, index = skipWhitespace(runes, size, index)
		return r, index, nil
	} else if r == TOKEN_RIGHT_CURLY {
		return r, index, nil
	} else {
		err := fmt.Sprintf("Expected '\"', found '%c'", r)
		return r, index, errors.New(err)
	}
}

func parsePropertyName(runes []rune, size int, index int) (string, int, error) {
	index++

	r := TOKEN_NULL
	start := index

	for index < size {
		r = runes[index]

		if r == TOKEN_QUOTE {
			end := index

			index++
			r, index = skipWhitespace(runes, size, index)
			if r == TOKEN_COLON {
				return string(runes[start:end]), index, nil
			} else {
				err := fmt.Sprintf("Expected ':', found '%c'", r)
				return "", index, errors.New(err)
			}
		}

		index++
	}

	err := fmt.Sprintf("Expected '\"', found '%c'", r)
	return "", index, errors.New(err)
}

func parseValue(runes []rune, size int, index int, parameterize bool) (*JsonValue, int, ParameterizedString, error) {
	r, index := skipWhitespace(runes, size, index)

	if r == TOKEN_QUOTE {
		index++

		if parameterize {
			ps, index, err := parseTextValueWithParameters(runes, size, index)
			if err != nil {
				return nil, index, ps, err
			}
			return nil, index, ps, nil
		} else {
			var sb strings.Builder

			for index < size {
				r := runes[index]

				if r == TOKEN_BACKSLASH {
					index++
					if index < size {
						test := runes[index]
						if test == TOKEN_R {
							sb.WriteRune(TOKEN_CR)
						} else if test == TOKEN_N {
							sb.WriteRune(TOKEN_LF)
						} else if test == TOKEN_T {
							sb.WriteRune(TOKEN_HT)
						} else if test == TOKEN_B {
							sb.WriteRune(TOKEN_BS)
						} else if test == TOKEN_F {
							sb.WriteRune(TOKEN_FF)
						} else if test == TOKEN_A {
							sb.WriteRune(TOKEN_BL)
						} else if test == TOKEN_V {
							sb.WriteRune(TOKEN_VT)
						}
					}
				} else if r == TOKEN_QUOTE {
					return &JsonValue{Value: sb.String(), Type: STRING}, index, ParameterizedString{}, nil
				} else {
					sb.WriteRune(r)
				}

				index++
			}
			return nil, index, ParameterizedString{}, nil
		}
	}

	if r == TOKEN_LEFT_CURLY {
		jo, index, err := parseObject(runes, size, index, parameterize)
		if err != nil {
			return nil, index, ParameterizedString{}, err
		}
		return &JsonValue{Value: *jo, Type: OBJECT}, index, ParameterizedString{}, nil
	}

	if r == TOKEN_LEFT_SQUARE {
		ja, index, err := parseArray(runes, size, index, parameterize)
		if err != nil {
			return nil, index, ParameterizedString{}, err
		}
		return &JsonValue{Value: *ja, Type: ARRAY}, index, ParameterizedString{}, nil
	}

	if r == TOKEN_T {
		index++
		if index < size && runes[index] == TOKEN_R {
			index++
			if index < size && runes[index] == TOKEN_U {
				index++
				if index < size && runes[index] == TOKEN_E {
					return &JsonValue{Value: true, Type: BOOLEAN}, index, ParameterizedString{}, nil
				}
			}
		}
	} else if r == TOKEN_F {
		index++
		if index < size && runes[index] == TOKEN_A {
			index++
			if index < size && runes[index] == TOKEN_L {
				index++
				if index < size && runes[index] == TOKEN_S {
					index++
					if index < size && runes[index] == TOKEN_E {
						return &JsonValue{Value: false, Type: BOOLEAN}, index, ParameterizedString{}, nil
					}
				}
			}
		}
	} else if r == TOKEN_N {
		index++
		if index < size && runes[index] == TOKEN_U {
			index++
			if index < size && runes[index] == TOKEN_L {
				index++
				if index < size && runes[index] == TOKEN_L {
					return &JsonValue{Value: nil, Type: NULL}, index, ParameterizedString{}, nil
				}
			}
		}
	} else if r > TOKEN_COMMA && r < TOKEN_COLON {
		start := index
		floating := false

		index++
		r = runes[index]
		for r > TOKEN_MINUS && r < TOKEN_COLON {
			if r == TOKEN_PERIOD {
				floating = true
			}
			index++
			r = runes[index]
		}

		if r == TOKEN_E || r == TOKEN_E_UPPER { //Scientific
			floating = true
			index++
			r = runes[index] // skip sign + -
			for index < size {
				index++
				r = runes[index]
				if r < TOKEN_0 || r > TOKEN_9 { //skip digits
					break
				}
			}
		}

		s := string(runes[start:index])

		index--

		if floating {
			f, err := strconv.ParseFloat(s, 64)
			if err == nil {
				return &JsonValue{Value: f, Type: NUMBER}, index, ParameterizedString{}, nil
			}
		} else {
			i, err := strconv.ParseInt(s, 0, 64)
			if err == nil {
				return &JsonValue{Value: i, Type: NUMBER}, index, ParameterizedString{}, nil
			}
		}

		err := fmt.Sprintf("Expected number, found '%s'", s)
		return nil, index, ParameterizedString{}, errors.New(err)
	}

	err := fmt.Sprintf("Unexpected character: '%c'", r)
	return nil, index, ParameterizedString{}, errors.New(err)
}

func parseObject(runes []rune, size int, index int, parameterize bool) (*JsonObject, int, error) {
	jo := JsonObject{}

	r, index, err := addProperty(&jo, runes, size, index, parameterize)
	if err != nil {
		return nil, index, err
	}

	for r == TOKEN_COMMA {
		r, index, err = addProperty(&jo, runes, size, index, parameterize)
		if err != nil {
			return nil, index, err
		}
	}

	if r != TOKEN_RIGHT_CURLY {
		err := fmt.Sprintf("Expected '}', found '%c'", r)
		return nil, index, errors.New(err)
	}

	return &jo, index, nil
}

func addValue(ja *JsonArray, runes []rune, size int, index int, parameterize bool) (rune, int, error) {
	index++
	r, index := skipWhitespace(runes, size, index)

	if r == TOKEN_RIGHT_SQUARE {
		return r, index, nil
	} else {
		value, index, ps, err := parseValue(runes, size, index, parameterize)
		if err != nil {
			return r, index, err
		}

		if ps.Value != "" {
			if ps.IsParameterized {
				value = &JsonValue{Value: ps, Type: PARAMETERIZED}
			} else {
				value = &JsonValue{Value: ps.Value, Type: STRING}
			}
		}

		if value != nil {
			ja.Add(*value)
		}
		index++
		r, index = skipWhitespace(runes, size, index)
		return r, index, nil
	}
}

func parseArray(runes []rune, size int, index int, parameterize bool) (*JsonArray, int, error) {
	ja := JsonArray{}

	r, index, err := addValue(&ja, runes, size, index, parameterize)
	if err != nil {
		return nil, index, err
	}

	for r == TOKEN_COMMA {
		r, index, err = addValue(&ja, runes, size, index, parameterize)
		if err != nil {
			return nil, index, err
		}
	}

	if r != TOKEN_RIGHT_SQUARE {
		err := fmt.Sprintf("Expected ']', found '%c'", r)
		return nil, index, errors.New(err)
	}

	return &ja, index, nil
}

func skipWhitespace(runes []rune, size int, index int) (rune, int) {
	for index < size {
		r := runes[index]

		if r == TOKEN_NULL || r == TOKEN_SPACE || r == TOKEN_LF || r == TOKEN_CR || r == TOKEN_HT || r == TOKEN_BS || r == TOKEN_FF || r == TOKEN_VT {
			index++
		} else {
			return r, index
		}
	}

	return TOKEN_NULL, index
}
