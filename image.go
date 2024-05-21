package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

// Save the image to disk
func (g *grid) SaveImage() {
	// Save the image to disk
	folder := "images"
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.Mkdir(folder, os.ModePerm)
	}
	filename := fmt.Sprintf("%s/%d-%d-%d-%d.png", folder, g.Width, g.Height, g.Bits, len(g.Cells))

	img := image.NewRGBA(image.Rect(0, 0, g.Width, g.Height))
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			v := g.Get(x, y)
			switch {
			case v == 0:
				img.Set(x, y, color.RGBA{0x00, 0x00, 0x00, 255})
			case v == 1:
				img.Set(x, y, color.RGBA{0x00, 0xff, 0x00, 255})
			case v == 2:
				img.Set(x, y, color.RGBA{0xff, 0x00, 0xff, 255})
			case v == 3:
				img.Set(x, y, color.RGBA{0xff, 0xd7, 0x00, 255})
			case v >= 4:
				img.Set(x, y, color.RGBA{0xff, 0xff, 0xff, 255})
			}
		}
	}

	f, _ := os.Create(filename)
	png.Encode(f, img)
	f.Close()

	imgShow := fmt.Sprintf("![%d-%d-%d-%d](%s)", g.Width, g.Height, g.Bits, len(g.Cells), "../../"+filename)
	d := fmt.Sprintf("%s|%d|%d|%d|%d|%s|  \n", Seconds2Days(g.RunTime), g.Width, g.Height, g.Bits, 1<<g.Bits, imgShow)

	fp := fmt.Sprintf("runreports/%s/", System.Unique)
	err := os.MkdirAll(fp, os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fn := fp + "data.md"
	f, _ = os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(d)
	f.Close()

	tt := time.Since(timeStart).Seconds()
	g.RunTime += tt

}
