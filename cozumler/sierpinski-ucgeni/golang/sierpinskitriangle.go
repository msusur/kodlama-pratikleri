package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var width = 4096
var img = image.NewRGBA(image.Rect(0, 0, width, width))
var col color.Color

func drawTriangleRotate180(len int, _x int, _y int) {
	for x := 0; x < len; x++ {
		for y := _y; y < (x+_y/2)*2; y++ {
			img.Set(x+_x, y, col)
		}
	}

	for x := 0; x < len; x++ {
		for y := _y; y < (_y+len*2)-x*2; y++ {
			img.Set(x+_x+len, y, col)
		}
	}
}

func drawSierpinskisTriangle(len, x, y int) {
	for {
		len = len / 2
		if len > 0 {
			drawTriangleRotate180(len, x+len, y-len*2) // top
			drawSierpinskisTriangle(len, x+len, y-len*2)

			drawTriangleRotate180(len, x-len, y+len*2) // left
			drawSierpinskisTriangle(len, x-len, y+len*2)

			drawTriangleRotate180(len, x+len*3, y+len*2) // right
			drawSierpinskisTriangle(len, x+len*3, y+len*2)
		} else {
			break
		}
	}

	// calc formula
	// top		=> len/2, x+len, y-len*2 => top => 100, 600+100, 400/2
	// left		=> len/2, x-len, y+len*2 => 100, 600-100, 400+100*2
	// right 	=> len/2, x+len*3, y+len*2 => 100, 600+100*3, 400+100*2
}

func main() {

	col = color.White
	for index := 0; index <= width; index++ {
		for index1 := 0; index1 < width; index1++ {
			img.Set(index, index1, col)
		}
	}

	//step 1=> main triangle
	col = color.Black
	for x := width / 2; x > 0; x-- {
		for y := 0; y < width; y++ {
			img.Set(x, y-(x-width/2)*2, col)
		}
	}
	for x := width / 2; x < width; x++ {
		for y := 0; y < width; y++ {
			img.Set(x, y-(width/2-x)*2, col)
		}
	}

	col = color.White

	len := width / 4
	x := width / 4
	y := width / 2

	//step 2=> draw first reverse triangle
	drawTriangleRotate180(len, x, y)

	//step 3=> draw drawSierpinskisTriangle by step 2
	drawSierpinskisTriangle(len, x, y)

	f, err := os.Create("SierpinskisTriangle.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
