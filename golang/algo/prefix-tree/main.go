package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	rootEdge := &Edge{}

	dictionary, err := LoadDictionary("test.txt")
	if err != nil {
		panic(err.Error())
	}

	for i := range dictionary {
		rootEdge.AddWord(dictionary[i])
	}

	fmt.Println(len(rootEdge.GetWordList()))
	data, _ := json.MarshalIndent(rootEdge.GetWordList(), "   ", "   ")
	fmt.Println(string(data))
}
