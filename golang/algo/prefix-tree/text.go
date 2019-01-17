package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadFile(filepath string) (text string, err error) {
	if data, err := ioutil.ReadFile(filepath); err != nil {
		return text, err
	} else {
		text = string(data)
	}

	return
}

func SanitizeText(dirty string) string {
	reg, _ := regexp.Compile("[^A-ZА-ЯЇҐа-яїіґa-z0-9\\(\\)\\[\\]_ ]+")
	return reg.ReplaceAllString(dirty, " ")
}

func LoadDictionary(filepath string) (dictionary []string, err error) {
	text, err := ReadFile(filepath)
	if err != nil {
		return
	}

	text = SanitizeText(text)
	text = strings.ToLower(text)

	dictionary = strings.Split(text, " ")

	return
}