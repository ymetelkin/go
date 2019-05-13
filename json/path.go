package json

import (
	"fmt"
	"strings"
)

//PathString returns string value from path
func (jo *Object) PathString(path string) (string, error) {
	var s string

	v, e := walk(*jo, strings.Split(path, "."))
	if e == nil {
		s, e = v.GetString()
	}
	return s, e
}

//PathStrings returns string values from path
func (jo *Object) PathStrings(path string) ([]string, error) {
	var ss []string

	v, e := walk(*jo, strings.Split(path, "."))
	if e == nil {
		a, e := v.GetArray()
		if e == nil {
			ss, e = a.GetStrings()
		}
	}
	return ss, e
}

//PathObject returns JSON object value from path
func (jo *Object) PathObject(path string) (Object, error) {
	var o Object

	v, e := walk(*jo, strings.Split(path, "."))
	if e == nil {
		o, e = v.GetObject()
	}
	return o, e
}

//PathObjects returns JSON object values from path
func (jo *Object) PathObjects(path string) ([]Object, error) {
	var os []Object

	v, e := walk(*jo, strings.Split(path, "."))
	if e == nil {
		a, e := v.GetArray()
		if e == nil {
			os, e = a.GetObjects()
		}
	}
	return os, e
}

func walk(jo Object, toks []string) (value, error) {
	if toks[0] == "" {
		return walkError(toks)
	}

	var (
		v value
		e error
	)

	for i, tok := range toks {
		if i == 0 {
			v, e = jo.getValue(tok)
			if e != nil {
				return walkError(toks)
			}
			continue
		}

		switch v.Type {
		case jsonObject:
			o, e := v.GetObject()
			if e != nil {
				return walkError(toks)
			}

			v, e = o.getValue(tok)
			if e != nil {
				return walkError(toks)
			}
		case jsonArray:
			a, e := v.GetArray()
			if e != nil {
				return walkError(toks)
			}

			objs, e := a.GetObjects()
			if e != nil {
				return walkError(toks)
			}

			ja := Array{}

			for _, obj := range objs {
				jv, e := walk(obj, toks[i:])
				if e != nil {
					return walkError(toks)
				}

				if jv.Type == jsonArray {
					aa, e := jv.GetArray()
					if e != nil {
						return walkError(toks)
					}
					for _, jvv := range aa.Values {
						ja.addValue(jvv)
					}
				} else {
					ja.addValue(jv)
				}
			}

			return newArray(ja), nil

		default:
			return walkError(toks)
		}
	}

	return v, nil
}

func walkError(toks []string) (value, error) {
	return value{}, fmt.Errorf("Invalid path: [%s]", strings.Join(toks, "."))
}
