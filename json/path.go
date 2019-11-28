package json

import (
	"strings"
)

//PathString returns string value from path
func (jo *Object) PathString(path string) (s string, ok bool) {
	v,k:= walk(*jo, strings.Split(path, "."))
	if k {
		s, ok = v.GetString()
	}
	return 
}

//PathStrings returns string values from path
func (jo *Object) PathStrings(path string) (ss []string, ok bool) {
	v, k := walk(*jo, strings.Split(path, "."))
	if k {
		ja, k := v.GetArray()
		if k {
			ss, ok = ja.GetStrings()
		}
	}
	return
}

//PathObject returns JSON object value from path
func (jo *Object) PathObject(path string) (o Object, ok bool) {
	v, k := walk(*jo, strings.Split(path, "."))
	if k {
		o, ok = v.GetObject()
	}
	return
}

//PathObjects returns JSON object values from path
func (jo *Object) PathObjects(path string) (jos []Object, ok bool) {
	v, k := walk(*jo, strings.Split(path, "."))
	if k {
		ja, k := v.GetArray()
		if k {
			jos, ok = ja.GetObjects()
		}
	}
	return
}

func walk(jo Object, toks []string) (v value, ok bool) {
	if toks[0] == "" {
		return
	}

	for i, tok := range toks {
		if i == 0 {
			v, ok = jo.getValue(tok)
			if !ok {
				return
			}
			continue
		}

		switch v.Type {
		case jsonObject:
			o, k := v.GetObject()
			if !k {
				return
			}
			v, ok = o.getValue(tok)
			if !ok {
				return
			}
		case jsonArray:
			a, k := v.GetArray()
			if !k {
				return
			}
			jos, k := a.GetObjects()
			if !k {
				return
			}

			var ja Array
			for _, o := range jos {
				jv, k := walk(o, toks[i:])
				if !k {
					return
				}
				if jv.Type == jsonArray {
					aa, k := jv.GetArray()
					if !k {
						return
					}
					for _, jvv := range aa.Values {
						ja.addValue(jvv)
					}
				} else {
					ja.addValue(jv)
				}
			}
			v = newArray(ja)
			ok = true
		}
	}

	return
}