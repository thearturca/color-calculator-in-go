package core

import (
	"color-calc/utils"
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

type Color struct {
	RGBA   color.RGBA
	Weight uint8
}

func SumColors(colors []Color) *Color {
	weightSum := 0
	for _, v := range colors {
		weightSum += int(v.Weight)
	}

	for i := range colors {
		NormalizeColor(&colors[i], int32(len(colors)), uint8(weightSum))
	}

	newRGBA := color.RGBA{0, 0, 0, 255}

	for _, v := range colors {
		newRGBA.R = AddWithOverflowCheck(newRGBA.R, v.RGBA.R)
		newRGBA.G = AddWithOverflowCheck(newRGBA.G, v.RGBA.G)
		newRGBA.B = AddWithOverflowCheck(newRGBA.B, v.RGBA.B)
	}

	return NewColor(newRGBA)
}

func SubColors(colors []Color) *Color {
	weightSum := 0
	for _, v := range colors {
		weightSum += int(v.Weight)
	}

	for _, v := range colors {
		NormalizeColor(&v, int32(len(colors)), uint8(weightSum))
	}

	newRGBA := color.RGBA{0, 0, 0, 255}

	for _, v := range colors {
		newRGBA.R = SubWithOverflowCheck(newRGBA.R, v.RGBA.R)
		newRGBA.G = SubWithOverflowCheck(newRGBA.G, v.RGBA.G)
		newRGBA.B = SubWithOverflowCheck(newRGBA.B, v.RGBA.B)
	}

	return NewColor(newRGBA)
}

func AddWithOverflowCheck(a, b uint8) uint8 {
	res := a + b

	if max := utils.MaxNumber(a, b); res < max {
		res = 255
	}

	return res
}

func SubWithOverflowCheck(a, b uint8) uint8 {
	res := a - b

	if b > a {
		res = 0
	}

	return res
}

func NewColor(color color.RGBA) *Color {
	return &Color{RGBA: color, Weight: 1}
}

func NormalizeColor(c *Color, colorCount int32, weightSum uint8) {
	c.RGBA.R = uint8(float64(c.RGBA.R) * (1 / float64(weightSum)) * float64(c.Weight))
	c.RGBA.G = uint8(float64(c.RGBA.G) * (1 / float64(weightSum)) * float64(c.Weight))
	c.RGBA.B = uint8(float64(c.RGBA.B) * (1 / float64(weightSum)) * float64(c.Weight))
	// c.RGBA.A = uint8(float64(c.RGBA.A) / (float64(c.Weight) / float64(colorCount)))
}

func GetColorFromHexString(hexString string) *Color {
	if !strings.HasPrefix(hexString, "#") || len(hexString) <= 1 {
		return nil
	}

	replacedString := strings.Replace(hexString, "#", "", 1)
	colorsStringSlice := utils.ChunkString(replacedString, 2)
	colorsSlice := make([]uint8, 0, len(colorsStringSlice))
	for _, value := range colorsStringSlice {
		parsedUint, parseError := strconv.ParseUint(value, 16, 8)

		if parseError != nil {
			return nil
		}

		colorsSlice = append(colorsSlice, uint8(parsedUint))
	}

	if len(colorsSlice) < 3 {
		till3 := 3 - len(colorsSlice)
		for i := 0; i < till3; i++ {
			colorsSlice = append(colorsSlice, 0)
		}
	}

	return &Color{RGBA: color.RGBA{R: colorsSlice[0], G: colorsSlice[1], B: colorsSlice[2], A: 255}, Weight: 1}
}

func GetHexStringFromColor(c Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.RGBA.R, c.RGBA.G, c.RGBA.B)
}
