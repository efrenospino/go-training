package jsontoxml

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func iterateSlice(dataSlice []interface{}, xmldata, key string) string {
	for _, value := range dataSlice {
		if t, ok := value.(map[string]interface{}); ok {
			xmldata = xmldata + fmt.Sprintf("<%v>", key)
			xmldata = iterateMap(t, xmldata)
			xmldata = xmldata + fmt.Sprintf("</%v>", key)
		}
		if t, ok := value.(string); ok {
			xmldata = xmldata + fmt.Sprintf("<%v>%v</%v>", key, t, key)
		}
		if t, ok := value.([]interface{}); ok {
			xmldata = iterateSlice(t, xmldata, key)

		}
	}
	return xmldata
}

func iterateMap(dataMap map[string]interface{}, xmldata string) string {
	for key, value := range dataMap {

		if t, ok := value.(string); ok {
			xmldata = xmldata + fmt.Sprintf("<%v>%v</%v>", key, t, key)
		}
		if t, ok := value.(map[string]interface{}); ok {
			xmldata = xmldata + fmt.Sprintf("<%v>", key)
			xmldata = iterateMap(t, xmldata)
			xmldata = xmldata + fmt.Sprintf("</%v>", key)
		}
		if t, ok := value.([]interface{}); ok {
			xmldata = iterateSlice(t, xmldata, key)
		}
	}
	return xmldata
}

//JSONToXML converts JSON to XML, max 3 levels
func JSONToXML(jsondata string) string {

	c := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsondata), &c)
	if err != nil {
		panic(err)
	}
	xmldata := `<?xml version="1.0" encoding="UTF-8"?>`
	for index := range c {
		xmldata = xmldata + fmt.Sprintf("<%v>", index)
		if t, ok := c[index].(map[string]interface{}); ok {
			xmldata = iterateMap(t, xmldata)
		}
		xmldata = xmldata + fmt.Sprintf("</%v>", index)
	}

	return cleanXML(xmldata)
}

func cleanXML(data string) string {
	resulting := bytes.Replace([]byte(data), []byte{10, 9, 9}, []byte{}, -1)
	resulting = bytes.Replace([]byte(resulting), []byte{10, 9}, []byte{}, -1)
	resulting = bytes.Replace([]byte(resulting), []byte{10}, []byte{}, -1)
	return string(resulting)
}
