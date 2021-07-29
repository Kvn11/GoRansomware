package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func EncryptDataAES(key []byte, data []byte) []byte {
	block, err := aes.NewCipher(key)
	checkError(err)

	aesGCM, err := cipher.NewGCM(block)
	checkError(err)
	nonce := make([]byte, aesGCM.NonceSize())
	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)

	return ciphertext
}

func DecryptDataAES(key []byte, data []byte) []byte {
	block, err := aes.NewCipher(key)
	checkError(err)

	aesGCM, err := cipher.NewGCM(block)
	checkError(err)

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	checkError(err)

	return plaintext
}
