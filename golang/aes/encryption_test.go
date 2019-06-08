package aes

import (
	"reflect"
	"testing"
)

func Test_createMD5Hash(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success retrieving MD5 hash sum",
			args: args{
				data: []byte("a-little-data"),
			},
			want: "0a8229f08ae19cf212bfeeb4e7633d21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createMD5Hash(tt.args.data); got != tt.want {
				t.Errorf("createMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_AES_EncryptAndDecrypt(t *testing.T) {
	type args struct {
		data       []byte
		passphrase string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Forward-and-Backward data encryption with AES",
			args: args{
				data:       []byte("plain text data"),
				passphrase: "passphrase",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//
			// Encrypt
			//
			encrypted, err := Encrypt(tt.args.data, tt.args.passphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("encryptWithAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			//
			// Decrypt
			//
			decrypted, err := Decrypt(encrypted, tt.args.passphrase)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(tt.args.data, decrypted) {
				t.Errorf("Decrypt() = %v, want %v", decrypted, tt.args.data)
			}
		})
	}
}
