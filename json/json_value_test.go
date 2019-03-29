package json

import (
	"fmt"
	"testing"
)

func TestJsonStringValue(test *testing.T) {
	jv := JsonStringValue{Value: "test"}
	v, t := jv.Get()
	fmt.Printf("%T\t%s %d\n", jv.Value, v, t)
	if t != STRING {
		test.Errorf("Expecting type %d, actual is %d", STRING, t)
	}

	i, err := getInt(&jv)
	if err == nil {
		test.Error("String cannot be converted to int")
	} else {
		fmt.Printf("%s; %s is not %d\n", err.Error(), jv.Value, i)
	}
}

func TestJsonIntValue(test *testing.T) {
	jv := JsonIntValue{Value: 123}
	v, t := jv.Get()
	fmt.Printf("%T\t%s %d\n", jv.Value, v, t)
	if t != NUMBER {
		test.Errorf("Expecting type %d, actual is %d", NUMBER, t)
	}

	i, err := getInt(&jv)
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %d is %d\n", jv.Value, i)
	}
}

func TestJsonFloatValue(test *testing.T) {
	jv := JsonFloatValue{Value: 1.23}
	v, t := jv.Get()
	fmt.Printf("%T\t%s %d\n", jv.Value, v, t)
	if t != NUMBER {
		test.Errorf("Expecting type %d, actual is %d", NUMBER, t)
	}

	f, err := getFloat(&jv)
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %f\n", jv.Value, f)
	}

	i, err := getInt(&jv)
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %d\n", jv.Value, i)
	}

	s, err := getString(&jv)
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %f is %s\n", jv.Value, s)
	}
}

func TestJsonBooleanValue(test *testing.T) {
	jv := JsonBooleanValue{Value: true}
	v, t := jv.Get()
	fmt.Printf("%T\t%s %d\n", jv.Value, v, t)
	if t != BOOLEAN {
		test.Errorf("Expecting type %d, actual is %d", BOOLEAN, t)
	}

	b, err := getBoolean(&jv)
	if err != nil {
		test.Error(err.Error())
	} else {
		fmt.Printf("Value of %t is %t\n", jv.Value, b)
	}
}
