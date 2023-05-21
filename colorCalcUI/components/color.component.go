package components

import (
	"color-calc/core"
	"strconv"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type Component struct {
	OnChanged     func()
	OnDelete      func()
	FyneComponent *fyne.Container
}

type ResultComponent struct {
	OnChanged             func()
	ColorBoxFyneComponent *fyne.Container
	SetResultHex          func(string)
	TextFyneComponent     *fyne.Container
}

func NewColorComponent(color *core.Color) *Component {
	vBox := container.NewVBox()
	component := &Component{FyneComponent: vBox}
	box := canvas.NewRectangle(&color.RGBA)
	box.SetMinSize(fyne.NewSize(100, 100))

	colorInput := widget.NewEntry()
	colorInput.SetPlaceHolder("Enter HEX Color")
	colorInput.OnChanged = func(text string) {
		newColor := core.GetColorFromHexString(text)
		if newColor != nil {
			*color = *newColor
			box.Refresh()
			if component.OnChanged != nil {
				component.OnChanged()
			}
		}
	}
	weightInput := widget.NewEntry()
	weightInput.SetPlaceHolder("Weight")
	weightInput.SetText(strconv.FormatUint(uint64(color.Weight), 10))
	weightInput.OnChanged = func(text string) {
		parsedWeight, parsedWeightError := strconv.ParseUint(text, 10, 64)

		if parsedWeightError != nil {
			parsedWeight = 1
		}

		color.Weight = uint8(parsedWeight)

		if component.OnChanged != nil {
			component.OnChanged()
		}
	}
	deleteButton := widget.NewButton("Delete", func() {
		// vBox.Hide()
		// weightInput.SetText("0")

		if component.OnDelete != nil {
			component.OnDelete()
		}

	})
	vBox.Add(box)
	vBox.Add(colorInput)
	vBox.Add(weightInput)
	vBox.Add(deleteButton)

	return component
}

func NewResultColorComponent(color *core.Color, w fyne.Window) *ResultComponent {
	box := canvas.NewRectangle(&color.RGBA)
	c := container.NewMax(box)
	hexString := core.GetHexStringFromColor(*color)
	resHexColor := widget.NewEntry()
	resHexColor.Disable()
	resHexColor.SetText(hexString)
	resHexColorButton := widget.NewButton("Copy", func() {})
	resHexColorButton.OnTapped = func() {
		w.Clipboard().SetContent(resHexColor.Text)
		prevText := resHexColorButton.Text
		resHexColorButton.SetText("Copied!")
		go func() {
			time.Sleep(3 * time.Second)
			resHexColorButton.SetText(prevText)
		}()
	}
	resHexBox := container.NewHBox(resHexColor, resHexColorButton)
	resHexBox.Move(fyne.NewPos(20, 20))

	component := &ResultComponent{ColorBoxFyneComponent: c, TextFyneComponent: fyne.NewContainerWithoutLayout(resHexBox)}
	component.SetResultHex = resHexColor.SetText
	return component
}
