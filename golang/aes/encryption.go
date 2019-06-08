package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

//
// Returns MD5 hash sum for the uploaded data
//
func createMD5Hash(data []byte) string {
	md5Hasher := md5.New()
	md5Hasher.Write(data)
	return hex.EncodeToString(md5Hasher.Sum(nil))
}

//
// Encrypts data by AES algorithm with given passphrase
//
func Encrypt(data []byte, passphrase string) ([]byte, error) {
	cipherKey := []byte(createMD5Hash([]byte(passphrase)))

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, data, nil), nil
}

//
// Decrypts data by AES algorithm with given passphrase
//
func Decrypt(data []byte, passphrase string) ([]byte, error) {
	cipherKey := []byte(createMD5Hash([]byte(passphrase)))

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
