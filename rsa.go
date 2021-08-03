package main

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func savePublicKey(pubPEM string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pubPEM))
	if block == nil {
		panic("failed to parse PEM block containing the public key.")
	}

	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	checkError(err)
	return pub
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

func ImportPublicKey(fileName string) *rsa.PublicKey {
	publicKeyFile, err := os.Open(fileName)
	checkError(err)
	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))
	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKCS1PublicKey(data.Bytes)
	checkError(err)
	return publicKeyImported
}

func createPrivateKey() *rsa.PrivateKey {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	checkError(err)
	return privatekey
}

func ExportPrivateKey(fileName string, privateKey *rsa.PrivateKey) int {
	pemPrivateFile, err := os.Create(fileName)
	checkError(err)

	var pemPrivateBlock = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	checkError(err)
	pemPrivateFile.Close()
	return 0
}

func ExportPublicKey(filename string, publicKey *rsa.PublicKey) int {
	pemPublicFile, err := os.Create(filename)
	checkError(err)

	var pemPublicBlock = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	}

	err = pem.Encode(pemPublicFile, pemPublicBlock)
	checkError(err)
	pemPublicFile.Close()
	return 0
}

func ImportPrivateKey(fileName string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(fileName)
	checkError(err)
	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	checkError(err)
	return privateKeyImported
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
