package json

import (
	"fmt"
	"strings"
)

//Array is JSON array structure
type Array struct {
	Values []value
}

func (ja *Array) addValue(jv value) int {
	if ja.Values == nil {
		ja.Values = []value{jv}
	} else {
		ja.Values = append(ja.Values, jv)
	}

	return len(ja.Values) - 1
}

func (ja *Array) addValues(jvs []value) int {
	if jvs != nil {
		if ja.Values == nil {
			ja.Values = jvs
		} else {
			ja.Values = append(ja.Values, jvs...)
		}
	}

	if ja.Values == nil {
		return 0
	}

	return len(ja.Values) - 1
}

//AddString adds string to JSON array
func (ja *Array) AddString(value string) int {
	return ja.addValue(newString(value))
}

//AddInt adds int to JSON array
func (ja *Array) AddInt(value int) int {
	return ja.addValue(newInt(value))
}

//AddFloat adds float to JSON array
func (ja *Array) AddFloat(value float64) int {
	return ja.addValue(newFloat(value))
}

//AddBool adds bool to JSON array
func (ja *Array) AddBool(value bool) int {
	return ja.addValue(newBool(value))
}

//AddObject adds JSON object to JSON array
func (ja *Array) AddObject(value Object) int {
	return ja.addValue(newObject(value))
}

//AddArray adds JSON array to JSON array
func (ja *Array) AddArray(value Array) int {
	return ja.addValue(newArray(value))
}

//AddStrings adds strings to JSON array
func (ja *Array) AddStrings(values []string) int {
	return ja.addValues(newStrings(values))
}

//AddInts adds ints to JSON array
func (ja *Array) AddInts(values []int) int {
	return ja.addValues(newInts(values))
}

//AddFloats adds floats to JSON array
func (ja *Array) AddFloats(values []float64) int {
	return ja.addValues(newFloats(values))
}

//AddBools adds bools to JSON array
func (ja *Array) AddBools(values []bool) int {
	return ja.addValues(newBools(values))
}

//AddObjects adds JSON objects to JSON array
func (ja *Array) AddObjects(values []Object) int {
	return ja.addValues(newObjects(values))
}

//AddArrays adds JSON arrays to JSON array
func (ja *Array) AddArrays(values []Array) int {
	return ja.addValues(newArrays(values))
}

func (ja *Array) setValue(index int, value value) error {
	if ja.Values == nil || len(ja.Values) <= index {
		return fmt.Errorf("Position [%d] does not exist", index)
	}
	ja.Values[index] = value
	return nil
}

//SetInt adds int JSON array at specified position
func (ja *Array) SetInt(index int, value int) error {
	return ja.setValue(index, newInt(value))
}

//SetFloat sets float in JSON array at specified position
func (ja *Array) SetFloat(index int, value float64) error {
	return ja.setValue(index, newFloat(value))
}

//SetBool sets bool in JSON array at specified position
func (ja *Array) SetBool(index int, value bool) error {
	return ja.setValue(index, newBool(value))
}

//SetString sets string in JSON array at specified position
func (ja *Array) SetString(index int, value string) error {
	return ja.setValue(index, newString(value))
}

//SetObject sets JSON object in JSON array at specified position
func (ja *Array) SetObject(index int, value Object) error {
	return ja.setValue(index, newObject(value))
}

//SetArray sets JSON array in JSON array at specified position
func (ja *Array) SetArray(index int, value Array) error {
	return ja.setValue(index, newArray(value))
}

//Remove removes element from JSON array at specified position
func (ja *Array) Remove(index int) error {
	if ja.Values == nil || len(ja.Values) <= index {
		return fmt.Errorf("Position [%d] does not exist", index)
	}

	values := ja.Values
	ja.Values = append(values[:index], values[index+1:]...)

	return nil
}

func (ja *Array) getValue(index int) (*value, error) {
	if ja.Values == nil || len(ja.Values) <= index {
		return nil, fmt.Errorf("Position [%d] does not exist", index)
	}

	return &ja.Values[index], nil
}

//GetString gets string from JSON array at specified position
func (ja *Array) GetString(index int) (string, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return "", err
	}

	s, err := jv.GetString()
	if err != nil {
		return "", err
	}

	return s, nil
}

//GetInt gets int from JSON array at specified position
func (ja *Array) GetInt(index int) (int, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return 0, err
	}

	i, err := jv.GetInt()
	if err != nil {
		return 0, err
	}

	return i, nil
}

//GetFloat gets bool from JSON array at specified position
func (ja *Array) GetFloat(index int) (float64, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return 0, err
	}

	f, err := jv.GetFloat()
	if err != nil {
		return 0, err
	}

	return f, nil
}

//GetBool gets bool from JSON array at specified position
func (ja *Array) GetBool(index int) (bool, error) {
	jv, err := ja.getValue(index)
	if err != nil {
		return false, err
	}

	b, err := jv.GetBool()
	if err != nil {
		return false, err
	}

	return b, nil
}

//GetObject gets JSON object from JSON array at specified position
func (ja *Array) GetObject(index int) (Object, error) {
	jv, err := ja.getValue(index)
	if err == nil {
		jo, err := jv.GetObject()
		if err == nil {
			return jo, nil
		}
	}

	return Object{}, err
}

//GetArray gets JSON array from JSON array at specified position
func (ja *Array) GetArray(index int) (Array, error) {
	jv, err := ja.getValue(index)
	if err == nil {
		a, err := jv.GetArray()
		if err == nil {
			return a, nil
		}
	}

	return Array{}, err
}

//GetStrings converts all values to strings
func (ja *Array) GetStrings() ([]string, error) {
	size := ja.Length()
	if size == 0 {
		return []string{}, nil
	}

	values := make([]string, size)
	for i, v := range ja.Values {
		value, err := v.GetString()
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return values, nil
}

//GetInts converts all values to ints
func (ja *Array) GetInts() ([]int, error) {
	size := ja.Length()
	if size == 0 {
		return []int{}, nil
	}

	values := make([]int, size)
	for i, v := range ja.Values {
		value, err := v.GetInt()
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return values, nil
}

//GetFloats converts all values to floats
func (ja *Array) GetFloats() ([]float64, error) {
	size := ja.Length()
	if size == 0 {
		return []float64{}, nil
	}

	values := make([]float64, size)
	for i, v := range ja.Values {
		value, err := v.GetFloat()
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return values, nil
}

//GetBools converts all values to bools
func (ja *Array) GetBools() ([]bool, error) {
	size := ja.Length()
	if size == 0 {
		return []bool{}, nil
	}

	values := make([]bool, size)
	for i, v := range ja.Values {
		value, err := v.GetBool()
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return values, nil
}

//GetObjects converts all values to JSON objects
func (ja *Array) GetObjects() ([]Object, error) {
	size := ja.Length()
	if size == 0 {
		return []Object{}, nil
	}

	values := make([]Object, size)
	for i, v := range ja.Values {
		value, err := v.GetObject()
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return values, nil
}

//GetArrays converts all values to JSON arrays
func (ja *Array) GetArrays() ([]Array, error) {
	size := ja.Length()
	if size == 0 {
		return []Array{}, nil
	}

	values := make([]Array, size)
	for i, v := range ja.Values {
		value, err := v.GetArray()
		if err != nil {
			return nil, err
		}
		values[i] = value
	}

	return values, nil
}

//IsEmpty checks if JSON array has any elements
func (ja *Array) IsEmpty() bool {
	if ja.Values == nil || len(ja.Values) == 0 {
		return true
	}

	return false
}

//Length returns number of elements in JSON array
func (ja *Array) Length() int {
	if ja.Values == nil {
		return 0
	}

	return len(ja.Values)
}

//String transforms JSON array to pretty string
func (ja *Array) String() string {
	return ja.toString(true, 0)
}

//InlineString transforms JSON array into inline string
func (ja *Array) InlineString() string {
	return ja.toString(false, 0)
}

func (ja *Array) toString(pretty bool, level int) string {
	if ja.Values == nil || len(ja.Values) == 0 {
		return "[]"
	}

	var sb strings.Builder

	sb.WriteRune(runeLeftSquare)
	if pretty {
		sb.WriteRune(runeCR)
		sb.WriteRune(runeLF)
	}

	next := level + 1

	for index, jv := range ja.Values {
		if index > 0 {
			sb.WriteRune(runeComma)

			if pretty {
				sb.WriteRune(runeCR)
				sb.WriteRune(runeLF)
			}
		}

		if pretty {
			i := 0
			for i <= level {
				sb.WriteRune(runeSpace)
				sb.WriteRune(runeSpace)
				i++
			}
		}

		s := jv.ToString(pretty, next)
		sb.WriteString(s)
	}

	if pretty {
		sb.WriteRune(runeCR)
		sb.WriteRune(runeLF)
		i := 0
		for i < level {
			sb.WriteRune(runeSpace)
			sb.WriteRune(runeSpace)
			i++
		}
	}
	sb.WriteRune(runeRightSquare)

	return sb.String()
}
