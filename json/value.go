package json

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	OBJECT int = iota
	ARRAY
	STRING
	INT
	FLOAT
	BOOL
	PARAMETERIZED
	NULL
)

type value struct {
	Value interface{}
	Type  int
	Text  string
}

func newInt(i int) *value {
	return &value{Value: i, Type: INT}
}

func newFloat(f float64) *value {
	return &value{Value: f, Type: FLOAT}
}

func newBool(b bool) *value {
	return &value{Value: b, Type: BOOL}
}

func newString(s string) *value {
	return &value{Value: s, Type: STRING}
}

func newObject(o *Object) *value {
	return &value{Value: o, Type: OBJECT}
}

func newArray(a *Array) *value {
	return &value{Value: a, Type: ARRAY}
}

func newNull() *value {
	return &value{Value: nil, Type: NULL}
}

func newParameterizedString(ps ParameterizedString) *value {
	return &value{Value: ps, Type: PARAMETERIZED}
}

func (jv *value) GetInt() (int, error) {
	if jv.Type == INT {
		i, ok := jv.Value.(int)
		if ok {
			return i, nil
		} else {
			u, ok := jv.Value.(uint)
			if ok {
				return int(u), nil
			}
		}
	} else if jv.Type == FLOAT {
		f, ok := jv.Value.(float64)
		if ok {
			return int(f), nil
		}
	} else if jv.Type == STRING {
		s, ok := jv.Value.(string)
		if ok {
			if strings.Contains(s, ".") {
				f, err := strconv.ParseFloat(s, 64)
				return int(f), err
			} else {
				i, err := strconv.ParseInt(s, 0, 64)
				return int(i), err
			}
		} else {
			return 0, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return 0, errors.New(err)
}

func (jv *value) GetFloat() (float64, error) {
	if jv.Type == FLOAT || jv.Type == INT {
		f, ok := jv.Value.(float64)
		if ok {
			return f, nil
		} else {
			err := fmt.Sprintf("Unsupported integer type: %T", jv.Value)
			return 0, errors.New(err)
		}
	} else if jv.Type == STRING {
		s, ok := jv.Value.(string)
		if ok {
			f, err := strconv.ParseFloat(s, 64)
			return f, err
		} else {
			return 0, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return 0, errors.New(err)
}

func (jv *value) GetString() (string, error) {
	if jv.Type == STRING {
		s, ok := jv.Value.(string)
		if ok {
			return s, nil
		} else {
			return "", errors.New("Cannot read string value")
		}
	} else {
		return jv.ToString(true, 0), nil
	}
}

func (jv *value) GetBool() (bool, error) {
	if jv.Type == BOOL {
		b, ok := jv.Value.(bool)
		if ok {
			return b, nil
		} else {
			return false, errors.New("Cannot read string value")
		}
	} else if jv.Type == STRING {
		s, ok := jv.Value.(string)
		if ok {
			b, err := strconv.ParseBool(s)
			return b, err
		} else {
			return false, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return false, errors.New(err)
}

func (jv *value) GetObject() (*Object, error) {
	if jv.Type == OBJECT {
		jo, ok := jv.Value.(*Object)
		if ok {
			return jo, nil
		} else {
			return nil, errors.New("Cannot read Object value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return nil, errors.New(err)
}

func (jv *value) GetArray() (*Array, error) {
	if jv.Type == ARRAY {
		ja, ok := jv.Value.(*Array)
		if ok {
			return ja, nil
		} else {
			return nil, errors.New("Cannot read Array value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return nil, errors.New(err)
}

func (jv *value) GetParameterizedString() (ParameterizedString, error) {
	if jv.Type == PARAMETERIZED {
		ps, ok := jv.Value.(ParameterizedString)
		if ok {
			return ps, nil
		} else {
			return ParameterizedString{}, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return ParameterizedString{}, errors.New(err)
}

func (jv *value) ToString(pretty bool, level int) string {
	if jv.Value == nil {
		return "null"
	}

	if jv.Text != "" {
		return jv.Text
	}

	switch jv.Type {
	case STRING:
		s, ok := jv.Value.(string)
		if ok {
			runes := []rune(s)
			var sb strings.Builder
			sb.WriteRune(TOKEN_QUOTE)

			for _, r := range runes {
				switch r {
				case TOKEN_QUOTE:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_QUOTE)
				case TOKEN_BACKSLASH:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_BACKSLASH)
				case TOKEN_CR:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_R)
				case TOKEN_LF:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_N)
				case TOKEN_HT:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_T)
				case TOKEN_BS:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_B)
				case TOKEN_FF:
					sb.WriteRune(TOKEN_BACKSLASH)
					sb.WriteRune(TOKEN_F)
				case TOKEN_VT:
					sb.WriteRune(TOKEN_SPACE)
				default:
					sb.WriteRune(r)
				}
			}

			sb.WriteRune(TOKEN_QUOTE)
			return sb.String()
		}
	case INT:
		i, ok := jv.Value.(int)
		if ok {
			return strconv.Itoa(i)
		}
	case FLOAT:
		f, ok := jv.Value.(float64)
		if ok {
			return strconv.FormatFloat(f, 'f', -1, 64)
		}
	case BOOL:
		b, ok := jv.Value.(bool)
		if ok {
			return strconv.FormatBool(b)
		}
	case OBJECT:
		jo, ok := jv.Value.(*Object)
		if ok {
			return jo.toString(pretty, level)
		}
	case ARRAY:
		ja, ok := jv.Value.(*Array)
		if ok {
			return ja.toString(pretty, level)
		}
	}

	return "null"
}
