package json

//Property property
type Property struct {
	Name  string
	Value Value
}

//Field Property constructor
func Field(name string, value Value) Property {
	return Property{
		Name:  name,
		Value: value,
	}
}
