package json

import (
	"fmt"
	"testing"
)

func TestStringValue(test *testing.T) {
	jv := newString("test")
	if jv.Type != jsonString {
		test.Errorf("Expecting type %d, actual is %d", jsonString, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	i, err := jv.GetInt()
	if err == nil {
		test.Error("String cannot be converted to int")
	} else {
		fmt.Printf("%s; %s is not %d\n", err.Error(), jv.Value, i)
	}
}

func TestIntValue(test *testing.T) {
	jv := newInt(123)
	i, err := jv.GetInt()
	if jv.Type != jsonInt {
		test.Errorf("Expecting type %d, actual is %d", jsonInt, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	i, err = jv.GetInt()
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %d is %d\n", jv.Value, i)
	}
}

func TestFloatValue(test *testing.T) {
	jv := newFloat(1.23)
	if jv.Type != jsonFloat {
		test.Errorf("Expecting type %d, actual is %d", jsonFloat, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	f, err := jv.GetFloat()
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %f\n", jv.Value, f)
	}

	i, err := jv.GetInt()
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %d\n", jv.Value, i)
	}

	s, err := jv.GetString()
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %s\n", jv.Value, s)
	}
}

func TestBooleanValue(test *testing.T) {
	jv := newBool(true)
	if jv.Type != jsonBool {
		test.Errorf("Expecting type %d, actual is %d", jsonBool, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	b, err := jv.GetBool()
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %t is %t\n", jv.Value, b)
	}
}
