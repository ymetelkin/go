package v2

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestObjectParse(t *testing.T) {
	s := `{ "text": "abc", "number": 3.14, "flag": true, "array": [ 1, 2, 3 ], "object": { "a": "b" }}`
	p := newParser([]byte(s))
	err := p.SkipWS()
	if err != nil {
		t.Error(err.Error())
	}
	if p.Byte != '{' {
		t.Error("Failed to parse {")
	}
	v, err := p.ParseObject(false)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(v.String())
}

func TestObjectPointers(t *testing.T) {
	jo := New(Field("name", String("YM")))
	fmt.Println(jo.String())
	jo.Add("person", jo)
	fmt.Println(jo.String())
}

func TestGraph(t *testing.T) {
	data, _ := ioutil.ReadFile("test_data/graph.json")
	jo, _ := ParseObject(data)
	ja, _ := jo.GetObjects("vertices")
	vertices := make(map[int]graphPerson)
	for i, v := range ja {
		name, _ := v.GetString("term")
		vertices[i] = graphPerson{
			Name: name,
		}
	}

	ja, _ = jo.GetObjects("connections")
	for _, o := range ja {
		source, _ := o.GetInt("source")
		target, _ := o.GetInt("target")
		weight, _ := o.GetFloat("weight")
		count, _ := o.GetInt("doc_count")
		v, _ := vertices[source]
		c, _ := vertices[target]
		v.Connections = append(v.Connections, graphConnection{
			Name:   c.Name,
			Weight: weight,
			Count:  count,
		})
		vertices[source] = v
	}

	for _, p := range vertices {
		if len(p.Connections) == 0 {
			continue
		}

		fmt.Println(p.Name)
		for _, c := range p.Connections {
			fmt.Printf("\t%s\n", c.Name)
		}
		fmt.Println()
	}

}

type graphPerson struct {
	Name        string
	Connections []graphConnection
}

type graphConnection struct {
	Name   string
	Weight float64
	Count  int
}
