package main

import (
	"color-calc/colorCalcUI"

	"fyne.io/fyne/app"
)

func main() {
	application := app.New()
	mainWindow := colorCalcUI.CreateMainWindow(application)

	mainWindow.ShowAndRun()
}
