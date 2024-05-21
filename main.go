package main

import (
	"flag"
	"fmt"

	"time"
)

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
