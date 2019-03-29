package json

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type JsonType int

const (
	OBJECT JsonType = iota
	ARRAY
	STRING
	NUMBER
	BOOLEAN
	PARAMETERIZED
	NULL
)

type JsonValue interface {
	Get() (interface{}, JsonType)
	ToString() string
}

type JsonIntValue struct {
	Value int
}

func (jv *JsonIntValue) Get() (interface{}, JsonType) {
	return jv.Value, NUMBER
}

func (jv *JsonIntValue) ToString() string {
	return strconv.Itoa(jv.Value)
}

type JsonFloatValue struct {
	Value  float64
	Format byte
}

func (jv *JsonFloatValue) Get() (interface{}, JsonType) {
	return jv.Value, NUMBER
}

func (jv *JsonFloatValue) ToString() string {
	var format byte
	if jv.Format == 0 {
		format = 'f'
	} else {
		format = jv.Format
	}
	return strconv.FormatFloat(jv.Value, format, -1, 64)
}

type JsonBooleanValue struct {
	Value bool
}

func (jv *JsonBooleanValue) Get() (interface{}, JsonType) {
	return jv.Value, BOOLEAN
}

func (jv *JsonBooleanValue) ToString() string {
	return strconv.FormatBool(jv.Value)
}

type JsonStringValue struct {
	Value string
}

func (jv *JsonStringValue) Get() (interface{}, JsonType) {
	return jv.Value, STRING
}

func (jv *JsonStringValue) ToString() string {
	runes := []rune(jv.Value)
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

type JsonObjectValue struct {
	Value JsonObject
}

func (jv *JsonObjectValue) Get() (interface{}, JsonType) {
	return jv.Value, OBJECT
}

func (jv *JsonObjectValue) ToString() string {
	return jv.Value.toString(true, 0)
}

type JsonArrayValue struct {
	Value JsonArray
}

func (jv *JsonArrayValue) Get() (interface{}, JsonType) {
	return jv.Value, ARRAY
}

type JsonNullValue struct {
}

func (jv *JsonNullValue) Get() (interface{}, JsonType) {
	return nil, NULL
}

func (jv *JsonNullValue) ToString() string {
	return "null"
}

func (jv *JsonArrayValue) ToString() string {
	return jv.Value.toString(true, 0)
}

func getInt(jv JsonValue) (int, error) {
	v, t := jv.Get()
	if t == NUMBER {
		i, ok := v.(int)
		if ok {
			return i, nil
		} else {
			f, ok := v.(float64)
			if ok {
				return int(f), nil
			} else {
				u, ok := v.(uint)
				if ok {
					return int(u), nil
				} else {
					err := fmt.Sprintf("Unsupported integer type: %T", v)
					return 0, errors.New(err)
				}
			}
		}
	} else if t == STRING {
		s, ok := v.(string)
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

	err := fmt.Sprintf("Unsupported value type: %d", t)
	return 0, errors.New(err)
}

func getFloat(jv JsonValue) (float64, error) {
	v, t := jv.Get()
	if t == NUMBER {
		f, ok := v.(float64)
		if ok {
			return f, nil
		} else {
			err := fmt.Sprintf("Unsupported integer type: %T", v)
			return 0, errors.New(err)
		}
	} else if t == STRING {
		s, ok := v.(string)
		if ok {
			f, err := strconv.ParseFloat(s, 64)
			return f, err
		} else {
			return 0, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", t)
	return 0, errors.New(err)
}

func getString(jv JsonValue) (string, error) {
	v, t := jv.Get()
	if t == STRING {
		s, ok := v.(string)
		if ok {
			return s, nil
		} else {
			return "", errors.New("Cannot read string value")
		}
	} else {
		return jv.ToString(), nil
	}
}

func getBoolean(jv JsonValue) (bool, error) {
	v, t := jv.Get()
	if t == BOOLEAN {
		b, ok := v.(bool)
		if ok {
			return b, nil
		} else {
			return false, errors.New("Cannot read string value")
		}
	} else if t == STRING {
		s, ok := v.(string)
		if ok {
			b, err := strconv.ParseBool(s)
			return b, err
		} else {
			return false, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", t)
	return false, errors.New(err)
}

func getObject(jv JsonValue) (*JsonObject, error) {
	v, t := jv.Get()
	if t == OBJECT {
		jo, ok := v.(JsonObject)
		if ok {
			return &jo, nil
		} else {
			return nil, errors.New("Cannot read JsonObject value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", t)
	return nil, errors.New(err)
}

func getArray(jv JsonValue) (*JsonArray, error) {
	v, t := jv.Get()
	if t == ARRAY {
		ja, ok := v.(JsonArray)
		if ok {
			return &ja, nil
		} else {
			return nil, errors.New("Cannot read JsonArray value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", t)
	return nil, errors.New(err)
}

func getParameterizedString(jv JsonValue) (ParameterizedString, error) {
	v, t := jv.Get()
	if t == PARAMETERIZED {
		ps, ok := v.(ParameterizedString)
		if ok {
			return ps, nil
		} else {
			return ParameterizedString{}, errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", t)
	return ParameterizedString{}, errors.New(err)
}

func toString(jv JsonValue, pretty bool, level int) string {
	v, t := jv.Get()
	if v == nil {
		return "null"
	}

	switch t {
	case OBJECT:
		jo, ok := v.(JsonObject)
		if ok {
			return jo.toString(pretty, level)
		}
	case ARRAY:
		ja, ok := v.(JsonArray)
		if ok {
			return ja.toString(pretty, level)
		}
	default:
		return jv.ToString()

	}

	return "null"
}
