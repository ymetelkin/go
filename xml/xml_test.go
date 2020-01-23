package xml

import (
	"io/ioutil"
	"strings"
	"testing"
)

type fs struct {
	Name string
	Body string
}

var tests []fs

func init() {
	files, err := ioutil.ReadDir("tests")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		data, err := ioutil.ReadFile("tests/" + file.Name())
		if err != nil {
			panic(err)
		}

		tests = append(tests, fs{
			Name: file.Name(),
			Body: strings.ReplaceAll(string(data), "ï", ""),
		})
	}
}

func Benchmark1(b *testing.B)  { benchmarkXML(b, 0) }
func Benchmark2(b *testing.B)  { benchmarkXML(b, 1) }
func Benchmark3(b *testing.B)  { benchmarkXML(b, 2) }
func Benchmark4(b *testing.B)  { benchmarkXML(b, 3) }
func Benchmark5(b *testing.B)  { benchmarkXML(b, 4) }
func Benchmark6(b *testing.B)  { benchmarkXML(b, 5) }
func Benchmark7(b *testing.B)  { benchmarkXML(b, 6) }
func Benchmark8(b *testing.B)  { benchmarkXML(b, 7) }
func Benchmark9(b *testing.B)  { benchmarkXML(b, 8) }
func Benchmark10(b *testing.B) { benchmarkXML(b, 9) }
func Benchmark11(b *testing.B) { benchmarkXML(b, 10) }
func Benchmark12(b *testing.B) { benchmarkXML(b, 11) }

func benchmarkXML(b *testing.B, idx int) {
	f := tests[idx]
	//fmt.Println(f.Name)
	_, err := ParseString(f.Body)
	if err != nil {
		b.Error(err.Error())
	}
}
