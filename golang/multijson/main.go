package main

import (
	"encoding/json"
	"github.com/ic2hrmk/snips/golang/multijson/dto"
	"log"
	"reflect"
)

func main() {
	testData := [][]byte{
		[]byte(`
		{
			"type": "a",
			"properties": {
				"cost": 10
			}
		}`),
		[]byte(`
		{
			"type": "b",
			"properties": {
				"is_cancelable": true
			}
		}`),
		[]byte(`		
		{
			"type": "c",
			"properties": {
				"max_duration": 5000
			}
		}`),
	}

	for _, sample := range testData {
		wrapper := &dto.JobWrapper{}

		if err := json.Unmarshal(sample, &wrapper); err != nil {
			log.Fatal(err)
		}

		log.Println(reflect.TypeOf(wrapper.GetJob()))

		switch wrapper.GetJob().(type) {
		case *dto.JobA:
			log.Printf("Type: %s, %+v", dto.JobTypeA, wrapper.GetJob().(*dto.JobA).Properties)

		case *dto.JobB:
			log.Printf("Type: %s, %+v", dto.JobTypeB, wrapper.GetJob().(*dto.JobB).Properties)

		case *dto.JobC:
			log.Printf("Type: %s, %+v", dto.JobTypeC, wrapper.GetJob().(*dto.JobC).Properties)

		}
	}
}
