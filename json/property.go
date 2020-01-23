package json

//Property represents JSON property
type Property struct {
	Field string
	Value Value
}

//StringField creates new string property
func StringField(field string, value string) Property {
	return Property{
		Field: field,
		Value: NewString(value),
	}
}

//IntField creates new int property
func IntField(field string, value int) Property {
	return Property{
		Field: field,
		Value: NewInt(value),
	}
}

//FloatField creates new float property
func FloatField(field string, value float64) Property {
	return Property{
		Field: field,
		Value: NewFloat(value),
	}
}

//BoolField creates new bool property
func BoolField(field string, value bool) Property {
	return Property{
		Field: field,
		Value: NewBool(value),
	}
}

//ObjectField creates new Object property
func ObjectField(field string, value Object) Property {
	return Property{
		Field: field,
		Value: NewObject(value),
	}
}

//ArrayField creates new Array property
func ArrayField(field string, value Array) Property {
	return Property{
		Field: field,
		Value: NewArray(value),
	}
}
