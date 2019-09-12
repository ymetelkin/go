package appl

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
)

type texts struct {
	Texts []json.Property
}

func (txts *texts) ParseTextComponent(pc pubcomponent) error {
	nd := pc.Node.Node("DataContent")
	nd = nd.Node("nitf")
	if nd.Nodes == nil {
		return errors.New("[nitf] block is missing")
	}

	nd = getBodyContent(nd)
	if nd.Nodes == nil {
		return errors.New("[body.content] block is missing")
	}

	var sb strings.Builder

	for _, b := range nd.Nodes {
		if b.Name == "block" {
			if b.Nodes != nil {
				var (
					txt strings.Builder
					add bool
				)
				for _, p := range b.Nodes {
					s, b := p.InlineString()
					if b {
						add = true
					}
					txt.WriteString(s)
				}

				if add {
					sb.WriteString(txt.String())
				}
			}

			if b.Text != "" {
				sb.WriteString(b.Text)
			}
		}
	}

	nitf := sb.String()
	if nitf == "" {
		return errors.New("[block] blocks are missing or empty")
	}

	jo := json.Object{}
	jo.AddString("nitf", nitf)

	nd = pc.Node.Node("Characteristics")
	nd = nd.Node("Words")
	i, err := strconv.Atoi(nd.Text)
	if err == nil && i > 0 {
		jo.AddInt("words", i)
	}

	jp := json.NewObjectProperty(strings.ToLower(pc.Role), jo)

	if txts.Texts == nil {
		txts.Texts = []json.Property{jp}
	} else {
		txts.Texts = append(txts.Texts, jp)
	}

	return nil
}

func (txts *texts) AddProperties(jo *json.Object) {
	if txts.Texts != nil {
		for _, jp := range txts.Texts {
			jo.AddProperty(jp)
		}
	}
}

func getBodyContent(nd xml.Node) xml.Node {
	if nd.Nodes != nil {
		for _, n := range nd.Nodes {
			if n.Name == "body.content" {
				return n
			}

			test := getBodyContent(n)
			if test.Nodes != nil {
				return test
			}
		}
	}
	return xml.Node{}
}
