package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	if 2+2 != 4 {
		t.Error("Addition not working")
	}
}

func TestCollectFromFile(t *testing.T) {
	//noinspection ALL
	var filepath string = "jfc_testdata/simple.json"
	jsn, err := loadJsonFromFile(filepath)
	if err != nil {
		t.Error(err.Error())
	}
	var collector = make(map[string][]string)
	walkJson(jsn, "default_field", collector)
	printCollector(collector, 5, "")
	if collector["a"][0] != "something" {
		t.Error("Parsing didn't work")
	}
}

//noinspection SpellCheckingInspection
func CompareStringsByLine(a, b string) error {
	var listsplit = strings.Split(b, "\n")
	var resultsplit = strings.Split(a, "\n")
	if len(listsplit) != len(resultsplit) {
		return errors.New(fmt.Sprintf("Unequal line numbers: %d vs %d", len(listsplit), len(resultsplit)))
	}
	var re1 = regexp.MustCompile(`\r`)
	var re2 = regexp.MustCompile(`\v`)
	for i := 0; i < len(listsplit); i++ {
		resultsplit[i] = re1.ReplaceAllString(resultsplit[i], "")
		resultsplit[i] = re2.ReplaceAllString(resultsplit[i], "")
		listsplit[i] = re1.ReplaceAllString(listsplit[i], "")
		listsplit[i] = re2.ReplaceAllString(listsplit[i], "")
		if resultsplit[i] != listsplit[i] {
			fmt.Println("Comparing string line should vs is for line number (differs here) : ", i)
			fmt.Println(resultsplit[i])
			fmt.Println(listsplit[i])
			fmt.Println([]byte(resultsplit[i]))
			fmt.Println([]byte(listsplit[i]))
			return errors.New(fmt.Sprint("String result differing from should be on line ", i))
		}
	}
	return nil
}

//noinspection SpellCheckingInspection
func loadStringFromFile(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
