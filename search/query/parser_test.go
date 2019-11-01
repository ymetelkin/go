package query

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	var (
		p     jsParser
		tests = []string{
			"status:active",
			`persons.name:"John Smith"`,
			"_exists_:title",
			"qu?ck bro*",
			"quikc~ brwn~1 foks~",
			`"fox quick"~5`,
			"quick^2 fox",
			`"john smith"^2   (foo bar)^4`,
			`body.text:"fox quick"~5`,
			"date:[2012-01-01 TO 2012-12-31]",
			"count:[1 TO 5]",
			"tag:{alpha TO omega}",
			"count:[10 TO *]",
			"date:{* TO 2012-01-01}",
			"count:[1 TO 5}",
			"age:>10",
			"age:>=10",
			"age:<10",
			"age:<=10",
			"quick OR brown",
			"title:(quick OR brown)",
			`book.\*:(quick OR brown)`,
			"((quick AND fox) OR (brown AND fox) OR fox) AND NOT news",
			"quick AND fox OR brown AND fox OR fox AND NOT news",
			`status:active AND persons.name:"John Smith" OR book.\*:(quick OR brown) AND NOT age:<=10`,
			`status:active and persons.name:"John Smith" AND age:<=10`,
		}
	)

	for i, test := range tests {
		qs, err := p.Parse(test)
		var out string
		if err != nil {
			t.Error(err.Error())
			out = err.Error()
		} else {
			out = qs.String()
		}
		fmt.Printf("%d\t%s\t%s\n", i+1, test, out)
	}

}
