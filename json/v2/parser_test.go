package v2

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	v1 "github.com/ymetelkin/go/json/v1"
)

var result *Object
var resultV1 v1.Object
var resultGo map[string]interface{}

func benchmarkParsing(b *testing.B, data []byte) {
	var jo *Object
	for x := 0; x < b.N; x++ {
		jo, _ = ParseObject(data)
	}
	result = jo
}

func benchmarkParsingV1(b *testing.B, data []byte) {
	var jo v1.Object
	for x := 0; x < b.N; x++ {
		jo, _ = v1.ParseObject(data)
	}
	resultV1 = jo
}

func benchmarkParsingGo(b *testing.B, data []byte) {
	var jo map[string]interface{}
	for x := 0; x < b.N; x++ {
		jo = make(map[string]interface{})
		json.Unmarshal(data, &jo)
	}
	resultGo = jo
}

func BenchmarkHits(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/hits.json")
	b.ResetTimer()
	benchmarkParsing(b, data)
}

func BenchmarkHitsV1(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/hits.json")
	b.ResetTimer()
	benchmarkParsingV1(b, data)
}

func BenchmarkHitsGo(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/hits.json")
	b.ResetTimer()
	benchmarkParsingGo(b, data)
}

func BenchmarkMessi(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/messi.json")
	b.ResetTimer()
	benchmarkParsing(b, data)
}

func BenchmarkMessiV1(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/messi.json")
	b.ResetTimer()
	benchmarkParsingV1(b, data)
}

func BenchmarkMessiGo(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/messi.json")
	b.ResetTimer()
	benchmarkParsingGo(b, data)
}

func BenchmarkSearch(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/search.json")
	b.ResetTimer()
	benchmarkParsing(b, data)
}

func BenchmarkSearchV1(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/search.json")
	b.ResetTimer()
	benchmarkParsingV1(b, data)
}

func BenchmarkSearchGo(b *testing.B) {
	data, _ := ioutil.ReadFile("test_data/search.json")
	b.ResetTimer()
	benchmarkParsingGo(b, data)
}
