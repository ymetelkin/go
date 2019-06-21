package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/ymetelkin/go/appl"
	"github.com/ymetelkin/go/xml"
)

func main() {
	dirname := "C:/tmp/appl/xml"

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	data := []string{}

	for i, file := range files {
		if i == 1000 {
			break
		}
		path := fmt.Sprintf("%s/%s", dirname, file.Name())
		fmt.Printf("%d\t%s\n", i+1, path)
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println(err.Error())
		}
		data = append(data, string(bytes))
	}

	size := len(data)

	start1 := time.Now()
	for _, file := range data {
		appl.XMLToJSON(file)
	}
	elapsed1 := time.Since(start1)

	var ids [1000]string

	start2 := time.Now()
	for _, file := range data {
		xml.New(file)
	}
	elapsed2 := time.Since(start2)

	for i, id := range ids {
		if i == size {
			break
		}
		fmt.Printf("%d\t%s\n", i+1, id)
	}

	fmt.Printf("Took %v to parse and %v to parse and read %d documents", elapsed1, elapsed2, size)
}
