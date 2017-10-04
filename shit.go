package shitter

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/disintegration/imaging"
)

// Shit contains an image that needs to be shit on
type Shit struct {
	In  io.Reader
	Out io.Writer
}

// Shit will shit on the incoming image
func (s Shit) Shit() error {
	rand.Seed(time.Now().Unix())

	btz, err := ioutil.ReadAll(s.In)
	if err != nil {
		return err
	}

	contentType := http.DetectContentType(btz)
	buff := bytes.NewBuffer(btz)
	var img image.Image

	switch contentType {
	case "image/png":
		img, err = png.Decode(buff)
		if err != nil {
			return err
		}
	case "image/jpeg":
		img, err := jpeg.Decode(buff)
		if err != nil {
			fail("Failed to decode jpeg", err)
		}
	default:
		return nil
	}

	shitData, err := Asset("shit.png")
	if err != nil {
		return err
	}
	shitBuff := bytes.NewBuffer(shitData)
	shitImage, err := png.Decode(shitBuff)
	if err != nil {
		return err
	}

	output := image.NewRGBA(img.Bounds())
	outputBounds := img.Bounds()
	draw.Draw(output, outputBounds, img, image.ZP, draw.Src)

	howManyPoops := random(0, 100)
	for j := 0; j <= howManyPoops; j++ {
		min := image.Point{
			random(0, outputBounds.Max.X),
			random(0, outputBounds.Max.Y),
		}
		rect := image.Rectangle{
			Min: min,
			Max: outputBounds.Max,
		}

		rotatedImage := imaging.Rotate(shitImage, float64(random(0, 360)), color.Transparent)
		draw.DrawMask(output, rect, rotatedImage, image.ZP, rotatedImage, image.ZP, draw.Over)
	}

	return png.Encode(s.Out, output)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
