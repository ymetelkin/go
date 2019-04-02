package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
)

type company struct {
	Name       string
	Code       string
	Creator    string
	Rels       uniqueArray
	Symbols    uniqueArray
	Industries uniqueArray
}

type ticker struct {
	Value    string
	ParentId string
}

type companies struct {
	Keys      map[string]int
	Companies []company
}

func (cs *companies) Parse(c Classification) {
	for _, o := range c.Occurrence {
		if o.Id != "" && o.Value != "" {
			var comp company

			key := fmt.Sprintf("%s_%s", o.Id, o.Value)

			if cs.Keys == nil {
				cs.Keys = make(map[string]int)
				cs.Companies = []company{}
			}

			i, ok := cs.Keys[key]
			if ok {
				comp = cs.Companies[i]
			} else {
				comp = company{Name: o.Value, Code: o.Id, Creator: c.System}
				comp.Rels.AddString("direct")
				cs.Companies = append(cs.Companies, comp)
				i = len(cs.Companies) - 1
				cs.Keys[key] = i
			}

			var (
				tickers   []ticker
				exchanges map[string]string
			)

			for _, prop := range o.Property {
				if prop.Name != "" && prop.Value != "" {
					name := strings.ToLower(prop.Name)
					if name == "apindustry" && prop.Id != "" {
						comp.Industries.AddKeyValue("code", prop.Id, "name", prop.Value)
					} else if name == "instrument" {
						instrument := strings.ToUpper(prop.Value)
						tokens := strings.Split(instrument, ":")
						if len(tokens) == 2 {
							symbol := json.Object{}
							symbol.AddString("ticker", tokens[1])
							symbol.AddString("exchange", tokens[0])
							symbol.AddString("instrument", instrument)
							comp.Symbols.AddObject(instrument, &symbol)
						}
					} else if name == "primaryticker" || name == "ticker" {
						t := ticker{Value: strings.ToUpper(prop.Value), ParentId: prop.ParentId}
						if tickers == nil {
							tickers = []ticker{t}
						} else {
							tickers = append(tickers, t)
						}
					} else if name == "exchange" {
						if exchanges == nil {
							exchanges = make(map[string]string)
						}
						exchanges[prop.Id] = strings.ToUpper(prop.Value)
					}
				}
			}

			if tickers != nil && exchanges != nil {
				def := ""
				for _, ticker := range tickers {
					var exchange string

					ex, ok := exchanges[ticker.ParentId]
					if ok {
						exchange = ex
					} else {
						if def == "" {
							for _, v := range exchanges {
								def = v
								break
							}
						}
						exchange = def
					}

					if exchange != "" {
						instrument := fmt.Sprintf("%s:%s", exchange, ticker.Value)
						symbol := json.Object{}
						symbol.AddString("ticker", ticker.Value)
						symbol.AddString("exchange", exchange)
						symbol.AddString("instrument", instrument)
						comp.Symbols.AddObject(instrument, &symbol)
					}
				}
			}

			if comp.Creator == "" || strings.EqualFold(c.System, "Editorial") {
				comp.Creator = c.System
			}

			cs.Companies[i] = comp
		}
	}
}

func (cs *companies) ToJsonProperty() *json.Property {
	if cs.Keys != nil {
		ja := json.Array{}
		for _, item := range cs.Companies {
			comp := item
			company := json.Object{}
			company.AddString("name", comp.Name)
			company.AddString("scheme", "http://cv.ap.org/id/")
			company.AddString("code", comp.Code)
			if comp.Creator != "" {
				company.AddString("creator", comp.Creator)
			}
			if !comp.Rels.IsEmpty() {
				company.AddProperty(comp.Rels.ToJsonProperty("rels"))
			}
			if !comp.Industries.IsEmpty() {
				company.AddProperty(comp.Industries.ToJsonProperty("industries"))
			}
			if !comp.Symbols.IsEmpty() {
				company.AddProperty(comp.Symbols.ToJsonProperty("symbols"))
			}

			ja.AddObject(&company)
		}

		return json.NewArrayProperty("companies", &ja)
	}

	return nil
}
