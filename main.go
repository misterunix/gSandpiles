package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"time"
)

// var Grid grid
var System system
var timeStart = time.Time{}
var timeImg = time.Time{}

// System struct
type system struct {
	Width     int
	Height    int
	BitsStart int
	BitsEnd   int
	Resume    bool
	Unique    string // Unique part of the filenames
	LastSave  string
}

// Grid struct
type grid struct {
	Width   int
	Height  int
	Bits    int
	RunTime float64
	Cells   []int
}

// Create a new grid with the given width, height, bits and runtime
func NewGrid(w, h, b int, rt float64) grid {
	g := grid{}
	g.Width = w
	g.Height = h
	g.Bits = b
	g.RunTime = rt
	g.Cells = make([]int, w*h)

	return g
}

// Set the value of the cell at x, y
func (g *grid) Set(x, y, v int) {
	g.Cells[y*g.Width+x] = v
}

// Get the value of the cell at x, y
func (g *grid) Get(x, y int) int {
	return g.Cells[y*g.Width+x]
}

// Resize the grid to the new size. Copying the old grid to the center of the new grid
func (g *grid) Resize(w, h int) grid {

	diffx := w - g.Width
	diffy := h - g.Height

	ng := NewGrid(w, h, g.Bits, g.RunTime)

	for y := 0; y < g.Width; y++ {
		xx := y + diffx/2
		for x := 0; x < g.Height; x++ {
			yy := x + diffy/2
			ng.Set(xx, yy, g.Get(x, y))
		}
	}

	//tt := time.Since(timeStart).Seconds()

	ng.RunTime = g.RunTime

	if System.LastSave != "" {
		os.Remove(System.LastSave)
	}

	byteArray, err := json.Marshal(g)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	jsonPath := fmt.Sprintf("savepoint/%s", System.Unique)
	os.MkdirAll(jsonPath, os.ModePerm)
	jsonFN := fmt.Sprintf("%s/%d.json", jsonPath, g.Bits)
	err = os.WriteFile(jsonFN, byteArray, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	System.LastSave = jsonFN
	//fmt.Println("Resized to ", w, h)

	// save System struct
	byteArray, err = json.Marshal(System)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	jsonFN = "system.json"
	err = os.WriteFile(jsonFN, byteArray, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return ng
}

// Topple the grid. Stopping only to resize the grid if the grid is too small
func (g *grid) Topple() {
	changed := true
	resize := false
	timeImg = time.Now()

	for changed {
		changed = false
		for y := 0; y < g.Height; y++ {

			for x := 0; x < g.Width; x++ {

				if g.Get(x, y) >= 4 {
					tv := g.Get(x, y) - 4
					changed = true
					g.Set(x, y, tv)

					if x < 16 || x > g.Width-16 || y < 16 || y > g.Height-16 {
						resize = true
					}

					if x-1 >= 0 {
						g.Set(x-1, y, g.Get(x-1, y)+1)
					}
					if x+1 < g.Width-1 {
						g.Set(x+1, y, g.Get(x+1, y)+1)
					}
					if y-1 >= 0 {
						g.Set(x, y-1, g.Get(x, y-1)+1)
					}
					if y+1 < g.Height-1 {
						g.Set(x, y+1, g.Get(x, y+1)+1)
					}
				}
			}
		}
		if resize {
			*g = g.Resize(g.Width+64, g.Height+64)
			resize = false
		}
	}

	g.RunTime += time.Since(timeImg).Seconds()

}

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

func (g *grid) StartReport() {
	fp := fmt.Sprintf("runreports/%s/", System.Unique)
	err := os.MkdirAll("runreports", os.ModePerm)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fn := fp + "data.md"
	f, _ := os.OpenFile(fn, os.O_CREATE|os.O_WRONLY, 0644)
	s := "Time|Width|Height|Bits|Grains|Image|  \n"
	f.WriteString(s)
	s += "|:-:|:-:|:-:|:-:|:-:|:-:|  \n"
	f.WriteString(s)
	f.Close()
}

// Convert seconds to days, hours, minutes and seconds
func Seconds2Days(s float64) string {

	var result string
	si := int(s)

	weeks := si / (7 * 24 * 3600) // get weeks from seconds
	si = si % (7 * 24 * 3600)

	day := si / (24 * 3600) // get days from seconds
	si = si % (24 * 3600)

	hour := si / 3600 // get hours from remaining seconds

	si %= 3600
	minutes := si / 60 // get minutes from remaining seconds

	si %= 60

	result = fmt.Sprintf("%dw:%dd:%02dh:%02dm:%02ds", weeks, day, hour, minutes, si)

	return result
}

func main() {

	timeStart = time.Now()
	td := time.Since(timeStart).Seconds()

	flag.IntVar(&System.Width, "w", 64, "Width of the grid")
	flag.IntVar(&System.Height, "h", 64, "Height of the grid")
	flag.IntVar(&System.BitsStart, "b", 4, "Bits start.")
	flag.IntVar(&System.BitsEnd, "e", 21, "Bits end.")
	flag.BoolVar(&System.Resume, "r", false, "Resume from savepoint")
	flag.StringVar(&System.Unique, "u", "test", "Unique part of the filenames")

	flag.Parse()

	for b := 4; b < 21; b++ {
		Grid := NewGrid(64, 64, b, td)
		grains := 1 << b
		Grid.Set(32, 32, grains)
		Grid.Topple()
		Grid.SaveImage()
	}

	fmt.Print("Done\n", Seconds2Days(time.Since(timeStart).Seconds()), "\n")

}
