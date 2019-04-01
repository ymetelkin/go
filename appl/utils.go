package appl

import "github.com/ymetelkin/go/json"

type UniqueStrings struct {
	keys   map[string]bool
	values []string
}

func (us *UniqueStrings) Add(s string) {
	if us.values == nil {
		us.keys = make(map[string]bool)
		us.keys[s] = true
		us.values = []string{s}
	} else {
		_, ok := us.keys[s]
		if !ok {
			us.keys[s] = true
			us.values = append(us.values, s)
		}
	}

}

func (us *UniqueStrings) IsEmpty() bool {
	return us.values == nil
}

/*
func (us *UniqueStrings) Size() int {
	if us.values == nil {
		return 0
	}

	return len(us.values)
}
*/

func (us *UniqueStrings) Values() []string {
	return us.values
}

func (us *UniqueStrings) ToJsonProperty(field string) *json.Property {
	if us.values == nil {
		return nil
	}

	ja := json.Array{}

	for _, s := range us.values {
		ja.AddString(s)
	}

	return json.NewArrayProperty(field, &ja)
}

func codeNamesToJsonArray(hash map[string]string) (*json.Array, bool) {
	if hash == nil || len(hash) == 0 {
		return nil, false
	}

	ja := json.Array{}

	for code, name := range hash {
		jo := json.Object{}
		jo.AddString("code", code)
		jo.AddString("name", name)
		ja.AddObject(&jo)
	}

	return &ja, true
}
