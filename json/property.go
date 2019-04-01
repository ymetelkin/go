package json

type Property struct {
	Field string
	Value value
}

func NewStringProperty(name string, value string) *Property {
	return &Property{Field: name, Value: *newString(value)}
}

func NewIntProperty(name string, value int) *Property {
	return &Property{Field: name, Value: *newInt(value)}
}

func NewFloatProperty(name string, value float64) *Property {
	return &Property{Field: name, Value: *newFloat(value)}
}

func NewBoolProperty(name string, value bool) *Property {
	return &Property{Field: name, Value: *newBool(value)}
}

func NewObjectProperty(name string, value *Object) *Property {
	return &Property{Field: name, Value: *newObject(value)}
}

func NewArrayProperty(name string, value *Array) *Property {
	return &Property{Field: name, Value: *newArray(value)}
}
