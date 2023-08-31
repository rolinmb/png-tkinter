package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"math"
	"image"
	"image/color"
    "image/png"
)

const (
	width = 1250
	height = 1250
)

func savePng(fname string, newPng *image.RGBA) {
    out,err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
    err = png.Encode(out, newPng)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Successfully created/rewritten "+fname+"\n")
}

func calcColor1(x, y, geometrySize int, scale, chaos, colorFactor float64) (uint8, uint8, uint8) {
    floatX := float64(x)
    floatY := float64(y)
	patternVal := math.Sin(floatX*scale) + math.Cos(floatY*scale)
	chaosFactor := math.Sin(chaos*(floatX+floatY))
	patternVal += chaosFactor
	geometryValue := math.Abs(math.Sin(patternVal*math.Pi*float64(geometrySize)))
	rShift := math.Sin(chaosFactor*math.Pi) * colorFactor
	gShift := math.Cos(chaosFactor*math.Pi) * colorFactor
	bShift := math.Sin(chaosFactor*2*math.Pi) * colorFactor
	r := uint8((math.Sin(floatX+geometryValue*math.Pi) + rShift) * 0.5 * 255)
	g := uint8((math.Cos(floatY+geometryValue*math.Pi) + gShift) * 0.5 * 255)
	b := uint8((math.Sin(floatX+floatY+geometryValue*math.Pi*2) + bShift) * 0.5 * 255)
	return r, g, b
}

func trippyPng1(fname string, scale, chaos, colorFactor float64, geometrySize, width, height int) {
    newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b := calcColor1(x, y, geometrySize, scale, chaos, colorFactor)
			newPng.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
    savePng(fname, newPng)
}

func calcColor2(x, y, geometrySize int, scale, chaos, colorFactor float64) (uint8, uint8, uint8) {
    floatX := float64(x)
    floatY := float64(y)
	patternVal := math.Sin(floatX*scale) + math.Cos(floatY*scale)
	chaosFactor := math.Sin(chaos*(floatX+floatY))
	patternVal += chaosFactor
	geometryValue := math.Abs(math.Sin(patternVal*math.Pi*float64(geometrySize)))
	rShift := math.Sin(chaosFactor*math.Pi) * colorFactor
	gShift := math.Cos(chaosFactor*math.Pi) * colorFactor
	bShift := math.Sin(chaosFactor*2*math.Pi) * colorFactor
	r := uint8((math.Sin(floatX+geometryValue*math.Pi) + rShift) * 0.5 * 255)
	g := uint8((math.Cos(floatY+geometryValue*math.Pi) + gShift) * 0.5 * 255)
	b := uint8((math.Sin(floatX+floatY+geometryValue*math.Pi*2) + bShift) * 0.5 * 255)
	return r, g, b
}

func trippyPng2(fname string, scale, chaos, colorFactor float64, geometrySize, width, height int) {
	newPng := image.NewRGBA(image.Rect(0,0,width,height))
    for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b := calcColor2(x, y, geometrySize, scale, chaos, colorFactor)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, 255})
		}
	}
    savePng(fname, newPng)
}

func calcColor3(x, y, geometrySize int, scale, chaos, colorFactor float64) (uint8, uint8, uint8) {
	floatX := float64(x)
    floatY := float64(y)
	patternVal := math.Sin(floatX*scale) + math.Cos(floatY*scale)
	chaosFactor := math.Sin(chaos*(floatX+floatY))
	patternVal += chaosFactor
	geometryValue := math.Abs(math.Sin(patternVal*math.Pi*float64(geometrySize)))
	rShift := math.Sin(chaosFactor*math.Pi) * colorFactor
	gShift := math.Cos(chaosFactor*math.Pi) * colorFactor
	bShift := math.Sin(chaosFactor*2*math.Pi) * colorFactor
	r := uint8((math.Sin(floatX+geometryValue*math.Pi) + rShift) * 0.5 * 255)
	g := uint8((math.Cos(floatY+geometryValue*math.Pi) + gShift) * 0.5 * 255)
	b := uint8((math.Sin(floatX+floatY+geometryValue*math.Pi*2) + bShift) * 0.5 * 255)
	return r, g, b
}

func trippyPng3(fname string, scale, chaos, colorFactor float64, geometrySize, width, height int) {
	newPng := image.NewRGBA(image.Rect(0,0,width,height))
    for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b := calcColor3(x, y, geometrySize, scale, chaos, colorFactor)
			newPng.SetRGBA(x, y, color.RGBA{r, g, b, 255})
		}
	}
    savePng(fname, newPng)
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

func applyDistortion(x, y int) (int, int) {
	amplitude := 50.0
	frequency := 0.02
	phase := 0.0
	dx := x + int(amplitude*math.Sin(frequency*float64(x)+phase))
	dy := y + int(amplitude*math.Sin(frequency*float64(y)+phase))
	dx = clamp(dx, 0, width-1)
	dy = clamp(dy, 0, height-1)
	return dx, dy
}

func calcColor4(x, y, geometrySize int, scale, chaos, colorFactor float64) (uint8, uint8, uint8) {
	floatX := float64(x)
    floatY := float64(y)
	patternVal := math.Sin(floatX*scale) + math.Cos(floatY*scale)
	chaosFactor := math.Sin(chaos*(floatX+floatY))
	patternVal += chaosFactor
	geometryValue := math.Abs(math.Sin(patternVal*math.Pi*float64(geometrySize)))
	rShift := math.Sin(chaosFactor*math.Pi) * colorFactor
	gShift := math.Cos(chaosFactor*math.Pi) * colorFactor
	bShift := math.Sin(chaosFactor*2*math.Pi) * colorFactor
	r := uint8((math.Sin(floatX+geometryValue*math.Pi) + rShift) * 0.5 * 255)
	g := uint8((math.Cos(floatY+geometryValue*math.Pi) + gShift) * 0.5 * 255)
	b := uint8((math.Sin(floatX+floatY+geometryValue*math.Pi*2) + bShift) * 0.5 * 255)
	return r, g, b
}

func trippyPng4(fname string, scale, chaos, colorFactor float64, geometrySize, width, height int) {
	newPng := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
            dx, dy := applyDistortion(x, y)
			r, g, b := calcColor4(dx, dy, geometrySize, scale, chaos, colorFactor)
			newPng.Set(x, y, color.RGBA{r, g, b, 255})
		}
	}
    savePng(fname, newPng)
}

func main() {
	pngOut := "C:\\code\\png-tkinter\\go_png_out\\"+os.Args[1]+".png"
	clrFactor, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Error converting Arg[2] (clrFactor) to float64:", err)
	}
	fmt.Println("Color Factor:", clrFactor)
	geoSize, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Error converting Args[3] (geoSize) to integer:", err)
	}
	fmt.Println("Geometry Size:", geoSize)
	scale, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Println("Error converting Args[4] (scale) to float64:", err)
	}
	fmt.Println("Scale:", scale)
	chaos, err := strconv.ParseFloat(os.Args[5], 64)
	if err != nil {
		fmt.Println("Error converting Args[5] (chaos) to float64:", err)
	}
	fmt.Println("Chaos:", chaos)
	whichGenerator := os.Args[6]
	fmt.Println("Selected Generator: "+whichGenerator)
	if whichGenerator == "Generator 1" {
		trippyPng1(pngOut, scale, chaos, clrFactor, geoSize, width, height)
	} else if whichGenerator == "Generator 2" {
		trippyPng2(pngOut, scale, chaos, clrFactor, geoSize, width, height)
	} else if whichGenerator == "Generator 3" {
		trippyPng3(pngOut, scale, chaos, clrFactor, geoSize, width, height)
	} else if whichGenerator == "Generator 4" {
		trippyPng4(pngOut, scale, chaos, clrFactor, geoSize, width, height)
	}
}