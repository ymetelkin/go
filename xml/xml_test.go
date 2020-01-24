package xml

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type fs struct {
	Name string
	Body []byte
}

var tests []fs

func init() {
	files, err := ioutil.ReadDir("tests")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		bs, err := ioutil.ReadFile("tests/" + file.Name())
		if err != nil {
			panic(err)
		}

		tests = append(tests, fs{
			Name: file.Name(),
			Body: bs,
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
func Benchmark13(b *testing.B) { benchmarkXML(b, 12) }
func Benchmark14(b *testing.B) { benchmarkXML(b, 13) }
func Benchmark15(b *testing.B) { benchmarkXML(b, 14) }
func Benchmark16(b *testing.B) { benchmarkXML(b, 15) }
func Benchmark17(b *testing.B) { benchmarkXML(b, 16) }
func Benchmark18(b *testing.B) { benchmarkXML(b, 17) }
func Benchmark19(b *testing.B) { benchmarkXML(b, 18) }
func Benchmark20(b *testing.B) { benchmarkXML(b, 19) }
func Benchmark21(b *testing.B) { benchmarkXML(b, 20) }
func Benchmark22(b *testing.B) { benchmarkXML(b, 21) }

func benchmarkXML(b *testing.B, idx int) {
	f := tests[idx]
	//fmt.Println(f.Name)
	nd, err := Parse(f.Body)
	if err != nil {
		fmt.Println(f.Name)
		fmt.Println(nd.String())
		b.Error(err.Error())
	}
}
