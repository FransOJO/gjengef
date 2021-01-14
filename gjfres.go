package main

import "fmt"

func convertToGcode(coordsArray [100][3]float64) [12]string {
	a := 1
	var result [12]string
	fmt.Println(pointAmountCalc(gstig, hdyb))
	result[0] = fmt.Sprint("G01 X", (coordsArray[0][0]), "\n")
	for a < pointAmountCalc(gstig, hdyb) {
		result[a] = fmt.Sprint("G02 X", coordsArray[a][0], " Y0", " Z", coordsArray[a][2], " I", -(coordsArray[a][0]), "\n")
		a++
	}
	return result
}

// global vars
var gstig float64
var gdia float64
var hdyb float64
var tdia float64
var hdia float64
var pointCount int = int(hdyb / gstig)

func pointAmountCalc(gstig, hdyb float64) int {
	return int(hdyb / gstig)
}

func pointCreator(gstig, gdia, hdyb, tdia, hdia float64) [100][3]float64 {
	var coordsArray [100][3]float64
	var i int
	var u float64
	diacalc := ((gdia - (tdia / 2)) / 2)
	/*	for i < pointAmountCalc(gstig, hdyb) {
		if i%2 == 0 {
			coordsArray[i][0] = -diacalc
		} else {
			coordsArray[i][0] = +diacalc
		}

		coordsArray[i][2] = gstig * u
		i++
		u++
	}*/

	for i < pointAmountCalc(gstig, hdyb) {
		coordsArray[i][0] = -diacalc
		coordsArray[i][2] = gstig * u
		i++
		u++
	}
	return coordsArray
}

func main() {

	fmt.Println("What is your thread diameter?")
	fmt.Scan(&gdia)
	fmt.Println("What is your thread pitch?")
	fmt.Scan(&gstig)
	fmt.Println("What is your hole depth?")
	fmt.Scan(&hdyb)
	fmt.Println("What is your pre bore hole diameter?")
	fmt.Scan(&hdia)
	fmt.Println("What is your tool diameter?")
	fmt.Scan(&tdia)

	fmt.Println(convertToGcode(pointCreator(gstig, gdia, hdyb, tdia, hdia)))
}
