package dto

import (
	"encoding/json"
	"fmt"
)

type Job interface{}

type JobWrapper struct {
	wrapper   jobWrapper `json:"-"`
	castedJob Job        `json:"-"`
}

type jobWrapper struct {
	Type       string          `json:"type"`
	Properties json.RawMessage `json:"properties"`
}

func (rcv *JobWrapper) GetJob() Job {
	return rcv.castedJob
}

func (rcv *JobWrapper) UnmarshalJSON(data []byte) error {
	var err error

	if err = json.Unmarshal(data, &rcv.wrapper); err != nil {
		return err
	}

	switch rcv.wrapper.Type {
	case JobTypeA:
		job := &JobA{}
		rcv.castedJob = job
		err = json.Unmarshal(rcv.wrapper.Properties, &job.Properties)

	case JobTypeB:
		job := &JobB{}
		rcv.castedJob = job
		err = json.Unmarshal(rcv.wrapper.Properties, &job.Properties)

	case JobTypeC:
		job := &JobC{}
		rcv.castedJob = job
		err = json.Unmarshal(rcv.wrapper.Properties, &job.Properties)

	default:
		return fmt.Errorf("job type unrecognized [%s]", rcv.wrapper.Type)
	}

	if err != nil {
		return err
	}

	return nil
}

//
// Custom Jobs
//

const (
	JobTypeA = "a"
	JobTypeB = "b"
	JobTypeC = "c"
)

type JobA struct {
	Properties *JobAProperties
}

type JobB struct {
	Properties *JobBProperties
}

type JobC struct {
	Properties *JobCProperties
}

type JobAProperties struct {
	Cost int `json:"cost"`
}

type JobBProperties struct {
	IsCancelable bool `json:"is_cancelable"`
}

type JobCProperties struct {
	MaxDuration int64 `json:"max_duration"`
}
