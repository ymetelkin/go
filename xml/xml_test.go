package xml

import (
	"fmt"
	"io/ioutil"
	"testing"
	"time"
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
			Body: string(data),
		})
	}
}

func TestXml(t *testing.T) {
	s := `	
  <?xml version="1.0" encoding="utf-8" standalone = "no" ?>
  <!--comments must be ignored-->
	<Publication Version="5.3.0" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <Identification>
    <ItemId>dd587867190648f3a8de18f25339fca8</ItemId>
    <RecordId>dd587867190648f3a8de18f25339fca8</RecordId>
    <CompositeId>00000000000000000000000000000000</CompositeId>
    <CompositionType>StandardPrintPhoto</CompositionType>
    <MediaType>Photo</MediaType>
    <Priority>5</Priority>
    <EditorialPriority>d</EditorialPriority>
    <DefaultLanguage>en-us</DefaultLanguage>
    <RecordSequenceNumber>0</RecordSequenceNumber>
    <FriendlyKey>19120542327253</FriendlyKey>
    <Title>Vikings, <a href="#">Sage Rosenfels</a> agree to <a href="#">2-year</a> contract</Title>
    <HeadLine><a href="#">Vikings</a>, Sage Rosenfels agree to 2-year contract</HeadLine>
  </Identification>
  <script>
   <![CDATA[
      <message> Welcome to TutorialsPoint </message>
   ]] >
</script >
</Publication>`

	//s = `<Identification><HeadLine>Vikings, Sage Rosenfels agree to 2-year contract</HeadLine></Identification>`

	start := time.Now()
	nd, err := ParseString(s)
	ts := time.Since(start)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", nd.String())
	fmt.Printf("Duration: %v\n", ts)
}

func BenchmarkXML(b *testing.B) {
	var i int
	for {
		if i == 10 {
			break
		}
		i++

		for _, f := range tests {
			//fmt.Println(f.Name)
			_, err := ParseString(f.Body)
			if err != nil {
				b.Error(err.Error())
			}
		}
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

func benchmarkXML(b *testing.B, idx int) {
	f := tests[idx]
	//fmt.Println(f.Name)
	_, err := ParseString(f.Body)
	if err != nil {
		b.Error(err.Error())
	}
}
