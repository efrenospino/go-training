package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	jsonData := `
    {
      "glossary": {
        "title": "example glossary",
        "GlossDiv": {
          "title": "S",
          "GlossList": {
            "GlossEntry": {
              "ID": "SGML",
              "SortAs": "SGML",
              "GlossTerm": "Standard Generalized Markup Language",
              "Acronym": "SGML",
              "Abbrev": "ISO 8879:1986",
              "GlossDef": {
                "para": "A meta-markup language, used to create markup languages such as DocBook.",
                "GlossSeeAlso": [
                  "GML",
                  "XML"
                ]
              },
              "GlossSee": "markup"
            }
          }
        }
      }
    }
  `

	var obj map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &obj)
	if err != nil {
		panic(err)
	}
	for _, v := range obj {
		for key := range v {
			fmt.Println(key)
		}

	}

}
