package query

import (
	"fmt"
	"testing"
)

func TestQueryString(t *testing.T) {
	var qs qsString

	qs = qsField{
		Field: "name",
		Term: qsSingleTerm{
			Value: "YM",
		},
	}
	s := qs.String()
	fmt.Println(s)
	if s != "name:YM" {
		t.Error("Failed single term")
	}

	qs = qsField{
		Field: "name",
		Term: qsSingleTerm{
			Value: "YM",
			Boost: 4,
		},
		Sign: "+",
	}
	s = qs.String()
	fmt.Println(s)
	if s != "+name:YM^4" {
		t.Error("Failed single term with sign and boost")
	}

	qs = qsField{
		Field: "name",
		Term: qsSingleTerm{
			Value: "YM",
			Fuzzy: 2,
		},
		Sign: "+",
	}
	s = qs.String()
	if s != "+name:YM~" {
		t.Error("Failed single term with sign and fuzzy")
	} else {
		fmt.Println(s)
	}

	qs = qsField{
		Field: "name",
		Term: qsSingleTerm{
			Value: "YM",
			Fuzzy: 1,
		},
		Sign: "+",
	}
	s = qs.String()
	if s != "+name:YM~1" {
		t.Error("Failed single term with sign and fuzzy")
	} else {
		fmt.Println(s)
	}

	qs = qsField{
		Field: "name",
		Term: qsPhraseTerm{
			Value:     "Yuri Metelkin",
			Proximity: 4,
		},
		Sign: "+",
	}
	s = qs.String()
	if s != `+name:"Yuri Metelkin"~4` {
		t.Error("Failed phrase term with sign and proximity")
	} else {
		fmt.Println(s)
	}

	qs = qsField{
		Field: "age",
		Term: qsRangeTerm{
			Left:         "25",
			Right:        "40",
			IncludeRight: true,
		},
		Sign: "-",
	}
	s = qs.String()
	fmt.Println(s)
	if s != "-age:{25 TO 40]" {
		t.Error("Failed range term with sign")
	}

	qs = qsQuery{
		Queries: []qsString{
			qsField{
				Field: "name",
				Term: qsPhraseTerm{
					Value:     "Yuri Metelkin",
					Proximity: 4,
				},
				Sign: "+",
			},
			qsField{
				Field: "age",
				Term: qsRangeTerm{
					Left:         "25",
					Right:        "40",
					IncludeRight: true,
				},
				Sign: "-",
			},
		},
	}
	s = qs.String()
	fmt.Println(s)
	if s != `+name:"Yuri Metelkin"~4 -age:{25 TO 40]` {
		t.Error("Failed boolean phrase and range terms")
	}

	qs = qsQuery{
		Operator: "AND",
		Queries: []qsString{
			qsField{
				Field: "name",
				Term: qsPhraseTerm{
					Value:     "Yuri Metelkin",
					Proximity: 4,
				},
			},
			qsField{
				Field: "age",
				Term: qsRangeTerm{
					Left:         "25",
					Right:        "40",
					IncludeRight: true,
				},
			},
		},
	}
	s = qs.String()
	fmt.Println(s)
	if s != `name:"Yuri Metelkin"~4 AND age:{25 TO 40]` {
		t.Error("Failed AND boolean phrase and range terms")
	}

	qs = qsQuery{
		Operator: "NOT",
		Queries: []qsString{
			qsField{
				Field: "name",
				Term: qsPhraseTerm{
					Value:     "Yuri Metelkin",
					Proximity: 4,
				},
			},
			qsField{
				Field: "age",
				Term: qsRangeTerm{
					Left:         "25",
					Right:        "40",
					IncludeRight: true,
				},
			},
		},
	}
	s = qs.String()
	fmt.Println(s)
	if s != `NOT name:"Yuri Metelkin"~4 AND NOT age:{25 TO 40]` {
		t.Error("Failed NOT boolean phrase and range terms")
	}

	qs = qsQuery{
		Operator: "AND",
		Queries: []qsString{
			qsQuery{
				Operator: "NOT",
				Queries: []qsString{
					qsField{
						Field: "name",
						Term: qsPhraseTerm{
							Value:     "Yuri Metelkin",
							Proximity: 4,
						},
					},
					qsField{
						Field: "age",
						Term: qsRangeTerm{
							Left:         "25",
							Right:        "40",
							IncludeRight: true,
						},
					},
				},
			},
			qsQuery{
				Operator: "OR",
				Queries: []qsString{
					qsField{
						Field: "name",
						Term: qsSingleTerm{
							Value: "YM",
							Boost: 4,
						},
					},
					qsField{
						Field: "age",
						Term: qsRangeTerm{
							Left:        "40",
							IncludeLeft: true,
						},
					},
				},
			},
		},
	}
	s = qs.String()
	fmt.Println(s)
	if s != `(NOT name:"Yuri Metelkin"~4 AND NOT age:{25 TO 40]) AND (name:YM^4 OR age:[40 TO *])` {
		t.Error("Failed group query")
	}
}
