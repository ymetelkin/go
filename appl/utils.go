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

func (us *UniqueStrings) ToJsonProperty(field string) *json.JsonProperty {
	if us.values == nil {
		return nil
	}

	ja := json.JsonArray{}

	for _, s := range us.values {
		ja.AddString(s)
	}

	return &json.JsonProperty{Field: field, Value: &json.JsonArrayValue{Value: ja}}
}

func codeNamesToJsonArray(hash map[string]string) (*json.JsonArray, bool) {
	if hash == nil || len(hash) == 0 {
		return nil, false
	}

	ja := json.JsonArray{}

	for code, name := range hash {
		jo := json.JsonObject{}
		jo.AddString("code", code)
		jo.AddString("name", name)
		ja.AddObject(&jo)
	}

	return &ja, true
}
