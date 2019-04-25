package json

//Property represents JSON property
type Property struct {
	Field string
	Value value
}

//IsEmtpy checks for field presense
func (jp *Property) IsEmtpy() bool {
	return jp.Field == ""
}

//NewStringProperty constructs new string value property
func NewStringProperty(name string, value string) Property {
	return Property{Field: name, Value: newString(value)}
}

//NewIntProperty constructs new int value property
func NewIntProperty(name string, value int) Property {
	return Property{Field: name, Value: newInt(value)}
}

//NewFloatProperty constructs new float value property
func NewFloatProperty(name string, value float64) Property {
	return Property{Field: name, Value: newFloat(value)}
}

//NewBoolProperty constructs new bool value property
func NewBoolProperty(name string, value bool) Property {
	return Property{Field: name, Value: newBool(value)}
}

//NewObjectProperty constructs new JSON object value property
func NewObjectProperty(name string, value Object) Property {
	return Property{Field: name, Value: newObject(value)}
}

//NewArrayProperty constructs new JSON array value property
func NewArrayProperty(name string, value Array) Property {
	return Property{Field: name, Value: newArray(value)}
}
