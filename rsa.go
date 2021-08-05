package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

func savePublicKey(pubPEM string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key.")
	}

	//pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	pubkeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	pubkey, _ := pubkeyInterface.(*rsa.PublicKey)
	checkError(err)
	return pubkey
}

func savePrivateKey(privPEM string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		panic("failed to parse PEM block containing the private key")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	checkError(err)
	return priv
}

func EncryptDataRSA(publicKey *rsa.PublicKey, dataToEncrypt []byte) []byte {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		publicKey,
		dataToEncrypt,
		nil,
	)
	checkError(err)
	return encryptedBytes
}

func DecryptDataRSA(privateKey *rsa.PrivateKey, encryptedBytes []byte) []byte {
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	checkError(err)
	return decryptedBytes
}
