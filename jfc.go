package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

func main() {
	println("Hello World")
}


type Map map[string]json.RawMessage
type Array []json.RawMessage

func walkJson(raw json.RawMessage, fieldname string, collector map[string][]string) {
	if (len(raw) <= 0) {
		return
	} else {

		if raw[0] == 123 { //  123 is `{` => object
			var cont Map
			json.Unmarshal(raw, &cont)
			for i, v := range cont {
				//println(i, ":")
				walkJson(v, i, collector)
			}
		} else if raw[0] == 91 { // 91 is `[`  => array
			var cont Array
			json.Unmarshal(raw, &cont)
			for i, v := range cont {
				//println(i, ":")
				walkJson(v, fieldname + "array_index_" + strconv.FormatInt(int64(i), 10), collector)
			}

		} else {
			var val interface{}
			json.Unmarshal(raw, &val)
			switch v := val.(type) {
			case float64:
				//println("float:", v)
				value := strconv.FormatFloat(v, 'E', -1, 64)
				appendIfMissing(collector, fieldname, value)
			case string:
				//println("string:", v)
				appendIfMissing(collector, fieldname, v)
			case bool:
				//println("bool:", v)
				value := strconv.FormatBool(v)
				appendIfMissing(collector, fieldname, value)
			case nil:
				//println("nil")
			default:
				//println("unkown type")
			}
		}
	}
}

func appendIfMissing(collector map[string][]string, fieldname string, value string) {
 	//TODO: implement check
	oldlist := collector[fieldname]
	if contains(oldlist, value) {
		//DO NOTHING
		//println("duplicated value for " , fieldname, " and value = ", value)
	} else {
		collector[fieldname] = append(oldlist, value)
	}
	//TODO: for more performance improvement, use sorted list and binary search
}

func contains(slice []string, value string) bool {
	for _, a := range slice {
		if a == value {
			return true
		}
	}
	return false
}

func printCollector(collector map[string][]string) {
	//TODO: print to console with given cmdline parameters
	bytes, err := json.MarshalIndent(collector, "", "    ")
	if (err != nil) {
		println("Panic: could not pretty print collector!")
	} else {
		println(string(bytes))
	}
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