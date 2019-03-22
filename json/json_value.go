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

type JsonValue struct {
	Value interface{}
	Type  JsonType
}

func NewJsonValue(value interface{}) (*JsonValue, error) {
	if value == nil {
		return &JsonValue{Value: nil, Type: NULL}, nil
	}

	switch value.(type) {
	case JsonValue:
		jv, ok := value.(JsonValue)
		if ok {
			return &jv, nil
		}
	case string:
		return &JsonValue{Value: value, Type: STRING}, nil
	case JsonObject:
		return &JsonValue{Value: value, Type: OBJECT}, nil
	case JsonArray:
		return &JsonValue{Value: value, Type: ARRAY}, nil
	case bool:
		return &JsonValue{Value: value, Type: BOOLEAN}, nil
	case ParameterizedString:
		return &JsonValue{Value: value, Type: PARAMETERIZED}, nil
	case int, int8, int16, int32, int64:
		return &JsonValue{Value: value, Type: NUMBER}, nil
	case float32, float64:
		return &JsonValue{Value: value, Type: NUMBER}, nil
	case uint, uint8, uint16, uint32, uint64:
		return &JsonValue{Value: value, Type: NUMBER}, nil
	}

	err := fmt.Sprintf("Unsupported value type: %T", value)
	return nil, errors.New(err)
}

func (jv JsonValue) GetInt() (int, error) {
	if jv.Type == NUMBER {
		i, ok := jv.Value.(int)
		if ok {
			return i, nil
		} else {
			f, ok := jv.Value.(float64)
			if ok {
				return int(f), nil
			} else {
				u, ok := jv.Value.(uint)
				if ok {
					return int(u), nil
				} else {
					err := fmt.Sprintf("Unsupported integer type: %T", jv.Value)
					return 0, errors.New(err)
				}
			}
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

func (jv JsonValue) GetFloat() (float64, error) {
	if jv.Type == NUMBER {
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

func (jv JsonValue) GetString() (string, error) {
	if jv.Type == STRING {
		s, ok := jv.Value.(string)
		if ok {
			return s, nil
		} else {
			return "", errors.New("Cannot read string value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return "", errors.New(err)
}

func (jv JsonValue) GetBoolean() (bool, error) {
	if jv.Type == BOOLEAN {
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

func (jv *JsonValue) GetObject() (*JsonObject, error) {
	if jv.Type == OBJECT {
		jo, ok := jv.Value.(JsonObject)
		if ok {
			return &jo, nil
		} else {
			return nil, errors.New("Cannot read JsonObject value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return nil, errors.New(err)
}

func (jv *JsonValue) GetArray() (*JsonArray, error) {
	if jv.Type == ARRAY {
		ja, ok := jv.Value.(JsonArray)
		if ok {
			return &ja, nil
		} else {
			return nil, errors.New("Cannot read JsonArray value")
		}
	}

	err := fmt.Sprintf("Unsupported value type: %d", jv.Type)
	return nil, errors.New(err)
}

func (jv JsonValue) GetParameterizedString() (ParameterizedString, error) {
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

func (jv *JsonValue) ToString() string {
	return jv.toString(true, 0)
}

func (jv *JsonValue) toString(pretty bool, level int) string {
	if jv.Value == nil {
		return "null"
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
	case OBJECT:
		jo, ok := jv.Value.(JsonObject)
		if ok {
			return jo.toString(pretty, level)
		}
	case NUMBER, BOOLEAN:
		return fmt.Sprintf("%v", jv.Value)
	case ARRAY:
		ja, ok := jv.Value.(JsonArray)
		if ok {
			return ja.toString(pretty, level)
		}
	}

	return "null"
}
