package main

import (
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

/**
必要なものフォントのttfファイル
フォントサイズ
画像の縦と横のサイズ
*/

var (
	fontsize = 125.0
	dpi      = flag.Float64("dpi", float64(fontsize*0.64), "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "/Users/oshikawatakashi/various/golang_project/generate_text_png/src/github.com/TakashiOshikawa/generate_text_png/Hack-Regular.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", float64(fontsize), "font size in points")
	text     = string("Taaaaaaka")
)

func main() {
	var inputText = flag.String("text", "", "Output text")
	flag.Parse()
	text = *inputText

	//fmt.Printf("Loading fontfile %q\n", *fontfile)
	b, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := truetype.Parse(b)
	if err != nil {
		log.Println(err)
		return
	}

	// Freetype context
	fg, bg := image.White, image.NewUniform(color.Alpha16{0}) // background colorless
	img := image.NewRGBA(image.Rect(0, 0, int(*size/1.3*float64(len(text))), int(*size*1.3)))
	draw.Draw(img, img.Bounds(), bg, image.ZP, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(fg)
	c.SetHinting(font.HintingNone)

	// Make some background

	// Truetype stuff
	opts := truetype.Options{}
	opts.Size = float64(*size)
	face := truetype.NewFace(f, &opts)

	// Calculate the widths and print to image
	for i, x := range text {
		awidth, ok := face.GlyphAdvance(rune(x))
		if ok != true {
			log.Println(err)
			return
		}
		iwidthf := int(/*float64(awidth) / */float64(*size/2)/1.1)
		//fmt.Printf("%+v\n", float64(*size/2)/2)

		//pt := freetype.Pt(i*int(*size*0.75)+(int(*size)-iwidthf/2)-int(*size/2), iwidthf*2)
		pt := freetype.Pt(i*int(*size*0.75), iwidthf*2)
		c.DrawString(string(x), pt)
		fmt.Printf("%+v\n", awidth)
	}

	// Save that image to disk.
	outFile, err := os.Create("text.png")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = png.Encode(outFile, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("OK.\n")

}
