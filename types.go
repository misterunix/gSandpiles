package main

import "time"

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
