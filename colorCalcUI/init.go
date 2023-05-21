package colorCalcUI

import (
	"color-calc/colorCalcUI/components"
	"color-calc/core"
	"color-calc/utils"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func CreateMainWindow(application fyne.App) fyne.Window {
	mainWindow := application.NewWindow("Color calculator")
	mainBox := container.NewMax()

	firstColor := core.NewColor(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	colors := []*core.Color{firstColor}

	resColor := core.SumColors(utils.GetValuesSliceFromPointerSlice(colors))
	resutlColorBox := components.NewResultColorComponent(resColor, mainWindow)

	calcColorsFunc := func(option string) {
		switch option {
		case "Add":
			*resColor = *core.SumColors(utils.GetValuesSliceFromPointerSlice(colors))
		case "Sub":
			*resColor = *core.SubColors(utils.GetValuesSliceFromPointerSlice(colors))
		}
		newResHexString := core.GetHexStringFromColor(*resColor)
		resutlColorBox.SetResultHex(newResHexString)
		for _, v := range resutlColorBox.ColorBoxFyneComponent.Objects {
			v.Refresh()
		}
	}

	calcSwitch := widget.NewRadioGroup([]string{"Add", "Sub"}, calcColorsFunc)
	calcSwitch.SetSelected("Add")

	hBoxColors := container.NewHBox()

	addColorButton := widget.NewButton("Add color", func() {
		newColor := core.NewColor(color.RGBA{R: 100, G: 100, B: 100, A: 255})
		colors = append(colors, newColor)
		newColorBox := components.NewColorComponent(newColor)
		newColorBox.OnChanged = func() { calcColorsFunc(calcSwitch.Selected) }
		newColorBox.OnDelete = func() {
			hBoxColors.Remove(newColorBox.FyneComponent)
		}
		newColorBox.OnChanged()
		hBoxColors.Add(newColorBox.FyneComponent)
	})
	hBoxColors.Add(addColorButton)

	firstColorVBox := components.NewColorComponent(firstColor)
	hBoxColors.Add(firstColorVBox.FyneComponent)
	firstColorVBox.OnDelete = func() {
		hBoxColors.Remove(firstColorVBox.FyneComponent)
	}

	firstColorVBox.OnChanged = func() {
		calcColorsFunc(calcSwitch.Selected)
	}

	// resutlColorBox.TextFyneComponent.Add(calcSwitch)
	colorsHScroll := container.NewHScroll(container.NewPadded(hBoxColors))
	colorsCenter := container.NewCenter(colorsHScroll)
	colorsHScroll.SetMinSize(fyne.NewSize(800, 100))
	mainWindow.SetContent(mainBox)
	mainBox.Add(resutlColorBox.ColorBoxFyneComponent)
	overlayStack := mainWindow.Canvas().Overlays()
	overlayStack.Add(container.NewBorder(resutlColorBox.TextFyneComponent, colorsCenter, nil, nil))

	mainWindow.Resize(fyne.NewSize(800, 400))

	return mainWindow
}
