package appl

import (
	"fmt"
	"strings"

	"github.com/ymetelkin/go/json"
	"github.com/ymetelkin/go/xml"
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
	ParentID string
}

type companies struct {
	Keys      map[string]int
	Companies []company
}

func (cs *companies) Parse(nd xml.Node) {
	if nd.Nodes == nil {
		return
	}

	system := nd.GetAttribute("System")

	for _, n := range nd.Nodes {
		code, name := getOccurrenceCodeName(n)
		if code != "" && name != "" {
			var comp company

			key := fmt.Sprintf("%s_%s", code, name)

			if cs.Keys == nil {
				cs.Keys = make(map[string]int)
				cs.Companies = []company{}
			}

			i, ok := cs.Keys[key]
			if ok {
				comp = cs.Companies[i]
			} else {
				comp = company{Name: name, Code: code, Creator: system}
				comp.Rels.AddString("direct")
				cs.Companies = append(cs.Companies, comp)
				i = len(cs.Companies) - 1
				cs.Keys[key] = i
			}

			var (
				tickers   []ticker
				exchanges map[string]string
			)

			if n.Nodes != nil {
				for _, p := range n.Nodes {
					if p.Attributes != nil {
						var id, pid, n, v string
						for _, a := range p.Attributes {
							switch a.Name {
							case "Id":
								id = a.Value
							case "Name":
								n = a.Value
							case "Value":
								v = a.Value
							case "ParentID":
								pid = a.Value
							}
						}

						if n != "" && v != "" {
							key := strings.ToLower(n)
							if key == "apindustry" && id != "" {
								comp.Industries.AddKeyValue("code", id, "name", v)
							} else if key == "instrument" {
								instrument := strings.ToUpper(v)
								tokens := strings.Split(instrument, ":")
								if len(tokens) == 2 {
									symbol := json.Object{}
									symbol.AddString("ticker", tokens[1])
									symbol.AddString("exchange", tokens[0])
									symbol.AddString("instrument", instrument)
									comp.Symbols.AddObject(instrument, symbol)
								}
							} else if key == "primaryticker" || key == "ticker" {
								t := ticker{Value: strings.ToUpper(v), ParentID: pid}
								if tickers == nil {
									tickers = []ticker{t}
								} else {
									tickers = append(tickers, t)
								}
							} else if key == "exchange" {
								if exchanges == nil {
									exchanges = make(map[string]string)
								}
								exchanges[id] = strings.ToUpper(v)
							}
						}
					}
				}
			}

			if tickers != nil && exchanges != nil {
				def := ""
				for _, ticker := range tickers {
					var exchange string

					ex, ok := exchanges[ticker.ParentID]
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
						comp.Symbols.AddObject(instrument, symbol)
					}
				}
			}

			if comp.Creator == "" || strings.EqualFold(system, "Editorial") {
				comp.Creator = system
			}

			cs.Companies[i] = comp
		}
	}
}

func (cs *companies) ToJSONProperty() json.Property {
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
				company.AddProperty(comp.Rels.ToJSONProperty("rels"))
			}
			if !comp.Industries.IsEmpty() {
				company.AddProperty(comp.Industries.ToJSONProperty("industries"))
			}
			if !comp.Symbols.IsEmpty() {
				company.AddProperty(comp.Symbols.ToJSONProperty("symbols"))
			}

			ja.AddObject(company)
		}

		return json.NewArrayProperty("companies", ja)
	}

	return json.Property{}
}
