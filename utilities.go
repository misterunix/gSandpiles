package main

import "fmt"

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
