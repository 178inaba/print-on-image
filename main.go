package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

const (
	fontSize = 20
	dpi      = 350
)

func main() {
	var drawValue, outputFilepath string
	flag.StringVar(&drawValue, "d", "", "Draw value.")
	flag.StringVar(&outputFilepath, "o", "", "Output filepath.")
	flag.Parse()

	f, err := os.Create(outputFilepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v.", err)
		os.Exit(1)
	}

	if err := run(drawValue, f); err != nil {
		fmt.Fprintf(os.Stderr, "%v.", err)
		os.Exit(2)
	}
}

func run(drawValue string, w io.Writer) error {
	if drawValue == "" {
		return errors.New("draw value is empty")
	}

	baseImg, err := getBaseImage()
	if err != nil {
		return err
	}

	fo, err := getFont()
	if err != nil {
		return err
	}

	drawImg := drawStringToImage(baseImg, fo, drawValue)
	return jpeg.Encode(w, drawImg, &jpeg.Options{Quality: 100})
}

func getBaseImage() (image.Image, error) {
	f, err := os.Open("static_files/images/src.jpg")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getFont() (*truetype.Font, error) {
	bs, err := ioutil.ReadFile("static_files/fonts/roboto-regular.ttf")
	if err != nil {
		return nil, err
	}

	fo, err := truetype.Parse(bs)
	if err != nil {
		return nil, err
	}

	return fo, nil
}

func drawStringToImage(baseImg image.Image, drawFont *truetype.Font, drawValue string) image.Image {
	r := baseImg.Bounds()
	drawImg := image.NewRGBA(image.Rect(0, 0, r.Dx(), r.Dy()))
	draw.Draw(drawImg, drawImg.Bounds(), baseImg, r.Min, draw.Src)

	d := &font.Drawer{Dst: drawImg, Src: image.Black}
	d.Face = truetype.NewFace(drawFont, &truetype.Options{Size: fontSize, DPI: dpi})
	d.Dot = fixed.Point26_6{
		X: (fixed.I(r.Dx()) - d.MeasureString(drawValue)) / 2,
		Y: fixed.I(r.Dy() / 2),
	}

	d.DrawString(drawValue)
	return d.Dst
}
