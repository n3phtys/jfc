package main

import (
	"encoding/json"
	"io/ioutil"
)

func main() {
	println("Hello World")
}



func walkJson(raw json.RawMessage, collector map[string][]string) {

}

func printCollector(collector map[string][]string) {

}

func loadJsonFromFile(filepath string) (json.RawMessage, error) {
	bytes, err1 := ioutil.ReadFile(filepath)
	var tmp json.RawMessage
	if (err1 != nil) {
		return  json.RawMessage{}, err1
	} else {
		err2 := json.Unmarshal(bytes, &tmp)
		if (err2 != nil) {
			return json.RawMessage{}, err2
		} else {
			return tmp, nil
		}
	}
}