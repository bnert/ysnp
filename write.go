package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"io/ioutil"

	"golang.org/x/crypto/scrypt"
)

func deriveKey(key, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key, err := scrypt.Key(key, salt, 1048576, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}

	return key, salt, nil
}

func WriteFile(key []byte, filedata FileData, filepath string) {
	dKey, salt, err := deriveKey(key, nil)
	if err != nil {
		panic(err)
	}

	bc, err := aes.NewCipher(dKey)
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(filedata)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(bc)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}

	encryptedText := gcm.Seal(nonce, nonce, jsonBytes, nil)
	encryptedText = append(encryptedText, salt...)
	err = ioutil.WriteFile(filepath, encryptedText, 0660)
	if err != nil {
		panic(err)
	}
}
