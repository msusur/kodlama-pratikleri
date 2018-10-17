package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y           float64
	DistanceToZero float64
}

func partition(a []Point, lo, hi int) int {

	p := a[hi].DistanceToZero
	for j := lo; j < hi; j++ {
		if a[j].DistanceToZero < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}

func quickSort(a []Point, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func readInput() []string {

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if len(lines) == 1 {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return lines
}

func createPlane(lines []string) []Point {

	plane := make([]Point, len(strings.Fields(lines[0])))

	for index, row := range strings.Fields(lines[0]) {

		arr := strings.Split(row, ",")

		x, _ := strconv.ParseFloat(arr[0], 64)
		y, _ := strconv.ParseFloat(arr[1], 64)

		var point Point
		point.X = x
		point.Y = y
		point.DistanceToZero = math.Sqrt(math.Pow((point.X-0), 2) + math.Pow((point.Y-0), 2))
		plane[index] = point

	}

	return plane
}

func closestPointsByNearToZero(plane []Point) map[int]Point {

	quickSort(plane, 0, len(plane)-1)

	points := make(map[int]Point)

	//min point
	firstPoint := plane[0]

	//check equal min each other
	for index, point := range plane {
		if point.DistanceToZero == firstPoint.DistanceToZero {
			points[index] = point
		} else {
			break
		}
	}

	return points
}

func main() {

	lines := readInput()

	plane := createPlane(lines)

	pointsNearToZero := closestPointsByNearToZero(plane)

	fmt.Println(pointsNearToZero)

	//input		: 2,2 2,3 3,2 3,3 3,4 3,5 2,3 2,5 3,1 1,3 4,3 4,4 4,5 1,4 1,5 4,1 4,2
	//output	: map[0:{2 2 2.8284271247461903}]

	//input		: 2,2 2,3 2,1 3,2 3,3 3,4 3,5 2,3 2,5 3,1 1,2 1,3 4,3 4,4 4,5 1,4 1,5 4,1 4,2
	//output	: map[0:{1 2 2.23606797749979} 1:{2 1 2.23606797749979}]
}
