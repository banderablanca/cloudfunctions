package utils

import (
	"bytes"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

const offset = 15

// SetLogoToImage mark image with logo bandera blanca
func SetLogoToImage(resourceOriginal io.ReadCloser, resourceType string) bytes.Buffer {
	var imgOriginal image.Image

	if resourceType == "image/png" {
		imgOriginal, _ = png.Decode(resourceOriginal)
	} else {
		imgOriginal, _ = jpeg.Decode(resourceOriginal)
	}

	imgOriginalBounds := imgOriginal.Bounds()
	imgOriginalWidth := imgOriginalBounds.Max.X

	resourceLogo, _ := os.Open("assets/logo_banderablanca.png")
	imgLogo, _ := png.Decode(resourceLogo)
	imgLogoBounds := imgLogo.Bounds()
	imgLogoWidth := imgLogoBounds.Max.X

	defer resourceLogo.Close()

	resourceOutput := image.NewRGBA(imgOriginalBounds)
	point := image.Point{X: imgOriginalWidth - (imgLogoWidth + offset), Y: offset}

	draw.Draw(resourceOutput, imgOriginalBounds, imgOriginal, image.ZP, draw.Src)
	draw.Draw(resourceOutput, imgLogoBounds.Add(point), imgLogo, image.ZP, draw.Over)

	var buffer bytes.Buffer
	jpeg.Encode(&buffer, resourceOutput, &jpeg.Options{})

	return buffer
}
