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

	s, ok := jv.GetString()
	if !ok {
		test.Error("Failed to get string")
	} else {
		fmt.Printf("Value of %v is %s\n", jv.Value, s)
	}
}

func TestIntValue(test *testing.T) {
	jv := newInt(123)
	if jv.Type != jsonInt {
		test.Errorf("Expecting type %d, actual is %d", jsonInt, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	i, ok := jv.GetInt()
	if !ok {
		test.Error("Failed to get int")
	} else {
		fmt.Printf("Value of %v is %d\n", jv.Value, i)
	}
}

func TestFloatValue(test *testing.T) {
	jv := newFloat(1.23)
	if jv.Type != jsonFloat {
		test.Errorf("Expecting type %d, actual is %d", jsonFloat, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	f, ok := jv.GetFloat()
	if !ok {
		test.Error("Failed to get float")
	} else {
		fmt.Printf("Value of %v is %f\n", jv.Value, f)
	}

	i, ok := jv.GetInt()
	if !ok {
		test.Error("Failed to get int")
	} else {
		fmt.Printf("Value of %v is %d\n", jv.Value, i)
	}

	s, ok := jv.GetString()
	if !ok {
		test.Error("Failed to get string")
	} else {
		fmt.Printf("Value of %v is %s\n", jv.Value, s)
	}
}

func TestBooleanValue(test *testing.T) {
	jv := newBool(true)
	if jv.Type != jsonBool {
		test.Errorf("Expecting type %d, actual is %d", jsonBool, jv.Type)
	} else {
		fmt.Printf("%T\t%v %d\n", jv.Value, jv.Value, jv.Type)
	}

	b, ok := jv.GetBool()
	if !ok {
		test.Error("Failed to get bool")
	} else {
		fmt.Printf("Value of %v is %t\n", jv.Value, b)
	}
}
