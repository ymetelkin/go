package json

import (
	"fmt"
	"testing"
)

func TestNewJsonValue(t *testing.T) {
	jv, err := newJsonValue("test")
	fmt.Printf("%T\t%s %d\n", jv.Value, jv.Value, jv.Type)
	if err != nil {
		t.Error(err.Error())
	}
	if jv.Type != STRING {
		t.Errorf("Expecting type %d, actual is %d", STRING, jv.Type)
	}

	i, err := jv.GetInt()
	if err == nil {
		t.Error("String cannot be converted to int")
	} else {
		fmt.Printf("%s; %s is not %d\n", err.Error(), jv.Value, i)
	}

	jv, err = newJsonValue(123)
	fmt.Printf("%T\t%d %d\n", jv.Value, jv.Value, jv.Type)
	if err != nil {
		t.Error(err.Error())
	}
	if jv.Type != NUMBER {
		t.Errorf("Expecting type %d, actual is %d", NUMBER, jv.Type)
	}

	i, err = jv.GetInt()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %d is %d\n", jv.Value, i)
	}

	jv, err = newJsonValue(1.23)
	fmt.Printf("%T\t%f %d\n", jv.Value, jv.Value, jv.Type)
	if err != nil {
		t.Error(err.Error())
	}
	if jv.Type != NUMBER {
		t.Errorf("Expecting type %d, actual is %d", NUMBER, jv.Type)
	}

	i, err = jv.GetInt()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %d\n", jv.Value, i)
	}

	jv, err = newJsonValue("1.23")
	i, err = jv.GetInt()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %s is %d\n", jv.Value, i)
	}

	f, err := jv.GetFloat()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %s is %f\n", jv.Value, f)
	}

	s, err := jv.GetString()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %s is %s\n", jv.Value, s)
	}

	jv, err = newJsonValue(true)
	b, err := jv.GetBoolean()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %t is %t\n", jv.Value, b)
	}

	jv, err = newJsonValue("true")
	b, err = jv.GetBoolean()
	if err != nil {
		t.Error(err.Error())
	} else {
		fmt.Printf("Value of %s is %t\n", jv.Value, b)
	}
}
