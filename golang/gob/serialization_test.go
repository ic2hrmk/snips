package gob

import (
	"fmt"
	"reflect"
	"testing"
)

type gobTestObject struct {
	ID     string
	Number int
}

func Test_GOB_SerializeAndDeserialize(t *testing.T) {
	type args struct {
		data gobTestObject
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Forward-and-Backward data encryption with AES",
			args: args{
				data: gobTestObject{
					ID:     "uuid",
					Number: 42,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//
			// Encrypt
			//
			serialized, err := Serialize(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(len(serialized))

			//
			// Decrypt
			//
			deserializedObject := new(gobTestObject)

			err = Deserialize(serialized, deserializedObject)
			if (err != nil) != tt.wantErr {
				t.Errorf("Deserialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.args.data, *deserializedObject) {
				t.Errorf("decryptWithAES() = %v, want %v", deserializedObject, tt.args.data)
			}
		})
	}
}
