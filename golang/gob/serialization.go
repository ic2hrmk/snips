package gob

import (
	"bytes"
	"encoding/gob"
)

func Serialize(object interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buffer)

	if err := encoder.Encode(object); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func Deserialize(data []byte, object interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)

	if err := decoder.Decode(object); err != nil {
		return err
	}

	return nil
}
