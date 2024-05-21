package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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
