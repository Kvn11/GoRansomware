package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	// privateKey := createPrivateKey()
	// publicKey := &privateKey.PublicKey

	// fmt.Println("Private Key: ", privateKey)
	// fmt.Println("Public Key: ", publicKey)
	// fmt.Println("N: ", publicKey.N)
	// fmt.Println("E: ", publicKey.E)

	// ExportPrivateKey("private.pem", privateKey)
	// ExportPublicKey("public.pem", publicKey)

	//privateKeyImported := ImportPrivateKey("private.pem")
	//fmt.Println("Private Key Imported: ", privateKeyImported)

	// publicKeyImported := ImportPublicKey("public.pem")
	// fmt.Println("Public Key Imported: ", publicKeyImported)
	myKey := make([]byte, 32)
	_, err := rand.Read(myKey)
	checkError(err)

	newFileName := EncryptFileName("message.txt", myKey)
	fmt.Println(newFileName)
}
