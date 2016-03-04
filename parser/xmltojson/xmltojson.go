package xmltojson

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"

	x2j "github.com/basgys/goxml2json"
)

type Node struct {
	Tag    string
	Value  string
	Parent *Node
}

func AddChild(parent, child *Node) {
	child.Parent = parent
}

func toJSON(data string) []Node {
	decoder := xml.NewDecoder(strings.NewReader(data))
	var nodes []Node
	for {
		element, err := decoder.Token()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		switch element.(type) {
		case xml.StartElement:
			node := Node{Tag: element.(xml.StartElement).Name.Local}
			nodes = append(nodes, node)
			break
		case xml.CharData:
			break
		case xml.EndElement:
			break
		}
	}
	return nodes
}

func XMLToJSON(data string) string {
	xmldata := strings.NewReader(data)
	json, err := x2j.Convert(xmldata)
	if err != nil {
		panic(err)
	}
	return json.String()
}

func MyXMLToJSON(xdata string) string {
	data := cleanXML(xdata)

	xmlReader := strings.NewReader(data)
	xmlDecoder := xml.NewDecoder(xmlReader)

	xmlNextDecoder := xml.NewDecoder(strings.NewReader(data))
	nextToken, _ := xmlNextDecoder.Token()

	var prevToken xml.Token

	startElement := xml.StartElement{}

	charData := xml.CharData{}

	endElement := xml.EndElement{}

	resultingJSON := ""

	nodesList := make([]Node, 0)

	for {

		token, err := xmlDecoder.Token()
		nextToken, _ = xmlNextDecoder.Token()

		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		switch element := token.(type) {
		case xml.ProcInst:
			resultingJSON = resultingJSON + fmt.Sprintf(`{`)
			break
		case xml.StartElement:
			if reflect.TypeOf(nextToken) != reflect.TypeOf(startElement) {
				resultingJSON = resultingJSON + fmt.Sprintf(`"%s":"%s"`, element.Name.Local, nextToken.(xml.CharData))

				node := Node{Tag: string(element.Name.Local), Value: string(nextToken.(xml.CharData))}
				if getIndexByTag(nodesList, node.Tag) == -1 {
					nodesList = append(nodesList, node)
				}

			} else if reflect.TypeOf(prevToken) == reflect.TypeOf(endElement) {
				resultingJSON = resultingJSON + fmt.Sprintf(`,{`)
			} else {
				resultingJSON = resultingJSON + fmt.Sprintf(`"%s":{`, element.Name.Local)
				child := Node{Tag: string(nextToken.(xml.StartElement).Name.Local)}
				node := Node{Tag: string(element.Name.Local)}
				AddChild(&node, &child)
				nodesList = append(nodesList, node)
				nodesList = append(nodesList, child)
			}
			break
		case xml.EndElement:
			if reflect.TypeOf(nextToken) == reflect.TypeOf(element) {
				resultingJSON = resultingJSON + fmt.Sprintf(`}`)
			} else if reflect.TypeOf(nextToken) == reflect.TypeOf(startElement) && reflect.TypeOf(prevToken) == reflect.TypeOf(charData) {
				resultingJSON = resultingJSON + fmt.Sprintf(`,`)
			}
			break
		case xml.CharData:
			break
		}
		prevToken = xml.CopyToken(token)
	}
	//JSON, _ := json.Marshal(nodesList)
	//fmt.Println(string(JSON))
	resultingJSON = resultingJSON + fmt.Sprintf("}")
	return resultingJSON
}

/*
n1 := Node{Tag: "breakfast_menu"}
n21 := Node{Tag: "name", Value: "Belgian Waffles"}
n22 := Node{Tag: "price", Value: "$5.74"}
n2 := Node{Tag: "food"}
n3 := Node{Tag: "food"}

AddChild(&n1, &n2)
AddChild(&n2, &n21)
AddChild(&n2, &n22)
AddChild(&n1, &n3)

nodesList := make([]Node, 0)

nodesList = append(nodesList, n1)
nodesList = append(nodesList, n2)
nodesList = append(nodesList, n21)
nodesList = append(nodesList, n22)
nodesList = append(nodesList, n3)

n23 := Node{Tag: "calories", Value: "4487"}
i := getIndexByTag(nodesList, "food")
AddChild(&nodesList[i], &n23)
nodesList = append(nodesList, n23)
*/

func printAll(nodesList []Node) {
	for _, node := range nodesList {
		fmt.Printf("Tag: %v | Value: %v | Parent: %v\n", node.Tag, node.Value, node.Parent)
	}
}

func getIndexByTag(nodesList []Node, tag string) int {
	for index, node := range nodesList {
		if node.Tag == tag {
			return index
		}
	}
	return -1
}

func cleanXML(data string) string {
	resulting := bytes.Replace([]byte(data), []byte{10, 9, 9}, []byte{}, -1)
	resulting = bytes.Replace([]byte(resulting), []byte{10, 9}, []byte{}, -1)
	resulting = bytes.Replace([]byte(resulting), []byte{10}, []byte{}, -1)
	return string(resulting)
}
