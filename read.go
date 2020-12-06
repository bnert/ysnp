package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"io/ioutil"

	"github.com/atotto/clipboard"
)

func ReadFile(key []byte, filename string) {
	encryptedText, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	etl := len(encryptedText)
	data, salt := encryptedText[:etl - 32], encryptedText[etl - 32:]

	dKey, _, err := deriveKey(key, salt)
	if err != nil {
		panic(err)
	}

	bc, err := aes.NewCipher(dKey)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(bc)
	ct, nonce := data[gcm.NonceSize():], data[:gcm.NonceSize()]

	jsonBytes, err := gcm.Open(nil, nonce, ct, nil)
	if err != nil {
		panic(err)
	}

	var f FileData
	json.Unmarshal(jsonBytes, &f)
	clipboard.WriteAll(f.Passwd)
	fmt.Println("username>", f.User)
	fmt.Println("password> copied to clipboard!")
}
