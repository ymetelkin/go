package json

import (
	"strings"
)

//PathString returns string value from path
func (jo *Object) PathString(path string) (s string, ok bool) {
	v, k := walk(*jo, strings.Split(path, "."))
	if k {
		s, ok = v.String()
	}
	return
}

//PathStrings returns string values from path
func (jo *Object) PathStrings(path string) (ss []string, ok bool) {
	v, k := walk(*jo, strings.Split(path, "."))
	if k {
		ja, k := v.Array()
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
		o, ok = v.Object()
	}
	return
}

//PathObjects returns JSON object values from path
func (jo *Object) PathObjects(path string) (jos []Object, ok bool) {
	v, k := walk(*jo, strings.Split(path, "."))
	if k {
		ja, k := v.Array()
		if k {
			jos, ok = ja.GetObjects()
		}
	}
	return
}

func walk(jo Object, toks []string) (v Value, ok bool) {
	if toks[0] == "" {
		return
	}

	for i, tok := range toks {
		if i == 0 {
			v, ok = jo.Get(tok)
			if !ok {
				return
			}
			continue
		}

		switch v.Type {
		case TypeObject:
			o, k := v.Object()
			if !k {
				return
			}
			v, ok = o.Get(tok)
			if !ok {
				return
			}
		case TypeArray:
			a, k := v.Array()
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
				if jv.Type == TypeArray {
					aa, k := jv.Array()
					if !k {
						return
					}
					for _, jvv := range aa.Values {
						ja.Values = append(ja.Values, jvv)
					}
				} else {
					ja.Values = append(ja.Values, jv)
				}
			}
			v = NewArray(ja)
			ok = true
		}
	}

	return
}
