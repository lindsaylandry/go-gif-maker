package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	_ "image/jpeg" // Register JPEG decoder
	_ "image/png"  // Register PNG decoder
	"os"

	"github.com/soniakeys/quant/median"

	"github.com/lindsaylandry/go-gif-maker/src/config"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	outGif := &gif.GIF{
		LoopCount: 0, // 0 means loop infinitely
	}

	for _, file := range c.InputFiles {
		path := c.Path + "/" + file
		// Open the image file
		f, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		
		// Decode into an image.Image object
		img, _, err := image.Decode(f)
		if err != nil {
			fmt.Printf("Error decoding %s: %v\n", path, err)
			return
		}

		q := median.Quantizer(c.Colors)
		cp := q.Quantize(make(color.Palette, 0, c.Colors), img)

		// 3. Convert image.Image to *image.Paletted
		bounds := img.Bounds()
		palettedImg := image.NewPaletted(bounds, cp)

		draw.FloydSteinberg.Draw(palettedImg, bounds, img, image.Point{})

		// 4. Append the frame and configuration to the GIF object
		outGif.Image = append(outGif.Image, palettedImg)
		outGif.Delay = append(outGif.Delay, 100*c.DelaySeconds) // Delay unit is 10ms
	}

	outputPath := c.Path + "/" + c.OutputFile

	fOut, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer fOut.Close()

	// 6. Encode and save the GIF
	err = gif.EncodeAll(fOut, outGif)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully created %s!\n", c.OutputFile)
}

