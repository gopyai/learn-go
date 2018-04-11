// xml
package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type (
	yed struct {
		tree []treeStruct

		elName    string
		processed bool
	}
	treeStruct struct {
		elName    string
		processed bool
	}
)

func main() {
	y := new(yed)
	y.parseXml()
}

func isProcessed(el string) bool {
	switch el {

	case "graph":
	case "graphml":

	case "Arrows":
	case "BorderInsets":
	case "BorderStyle":
	case "Fill":
	case "GenericEdge":
	case "GenericGroupNode":
	case "GenericNode":
	case "Geometry":
	case "Insets":
	case "LabelModel":
	case "LineStyle":
	case "ModelParameter":
	case "Path":
	case "Point":
	case "ProxyAutoBoundsNode":
	case "Realizers":
	case "Resources":
	case "SmartNodeLabelModel":
	case "SmartNodeLabelModelParameter":
	case "State":
	case "StyleProperties":

	default:
		return true

	}
	return false
}

func (y *yed) evStartElement(el string, attr map[string]string) {
	fullName := ""
	for _, s := range y.tree {
		if len(fullName) > 0 {
			fullName += " + "
		}
		fullName += s.elName
	}
	fullName += " => "
	fmt.Println(fullName, attr)
}

func (y *yed) evData(s string) {
	fmt.Println("Data:", s)
}

func (y *yed) evEndElement(s string) {
}

func (y *yed) parseXml() {
	f, e := os.Open("yed2.graphml")
	isErr(e)
	defer f.Close()
	dec := xml.NewDecoder(f)

	y.tree = make([]treeStruct, 0)

	for {
		t, _ := dec.Token()
		if t == nil {
			break
		}
		switch el := t.(type) {
		case xml.ProcInst:
		case xml.CharData:

			// Processing
			if !y.processed {
				break
			}
			if d := strings.TrimSpace(string(el)); len(d) > 0 {
				y.evData(d)
			}

		case xml.StartElement:

			// Tree
			y.elName = el.Name.Local
			y.processed = isProcessed(el.Name.Local)
			y.tree = append(y.tree, treeStruct{el.Name.Local, y.processed})

			// Processing
			if y.processed {
				attr := make(map[string]string)
				for _, at := range el.Attr {
					attr[at.Name.Local] = at.Value
				}
				y.evStartElement(el.Name.Local, attr)
			}

		case xml.EndElement:

			// Processing
			if y.processed {
				y.evEndElement(el.Name.Local)
			}

			// Tree
			if y.tree[len(y.tree)-1].elName != el.Name.Local {
				panic("Error end of element")
			}
			y.tree = y.tree[:len(y.tree)-1]

			if len(y.tree) > 0 {
				last := y.tree[len(y.tree)-1]
				y.elName = last.elName
				y.processed = last.processed
			} else {
				y.elName = ""
				y.processed = false
			}

		case xml.Comment:
		default:
			panic("TODO")
		}
	}
}

func isErr(e error) {
	if e != nil {
		panic(e)
	}
}
