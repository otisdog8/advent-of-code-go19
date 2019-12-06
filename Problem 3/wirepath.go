package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type line struct {
	x1, y1, x2, y2 int
}

type movement struct {
	dir, mag int
}

func main() {
	data, err := ioutil.ReadFile("inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	var wires []string = strings.Split(string(data), "\r\n")

	var wirestr1 []string = strings.Split(wires[0], ",")
	var wirestr2 []string = strings.Split(wires[0], ",")
	var wire1 []movement = make([]movement, len(wirestr1))
	var wire2 []movement = make([]movement, len(wirestr2))
	var lines1 []line = make([]line, len(wire1))
	var lines2 []line = make([]line, len(wire2))

	for i, v := range wirestr1 {
		wire1[i] = movement{calcdir(string(v[0])), sliceindexconvert(v[0:])}
	}

	for i, v := range wirestr1 {
		wire2[i] = movement{calcdir(string(v[0])), sliceindexconvert(v[0:])}
	}

	lines1 = calcline(wire1)
	lines2 = calcline(wire2)

	for _, v1 := range lines1 {
		for _, v2 := range lines2 {
			if lineintersects(v1, v2) {
				fmt.Println("intersect")
			}
		}
	}

}

func sliceindexconvert(stuff string) int {
	var num int = 0
	num, _ = strconv.Atoi(stuff)

	return num
}

func calcdir(s string) int {
	if s == "U" {
		return 3
	}
	if s == "D" {
		return -3
	}
	if s == "R" {
		return 4
	}
	if s == "L" {
		return -4
	}
	return 0
}

func calcline(wire []movement) []line {
	var x1 int = 0
	var y1 int = 0
	var x2 int = 0
	var y2 int = 0
	var lines []line
	lines = make([]line, len(wire))

	for i, v := range wire {
		x2 = x1 + (v.dir%3)*v.mag
		y2 = y1 + (v.dir%2)*v.mag
		lines[i] = line{x1, y1, x2, y2}
	}

	return lines
}

func lineintersects(line1, line2 line) bool {
	if min(line1.x1, line1.x2) <= float64(line2.x1) && float64(line2.x1) <= max(line1.x1, line1.x2) {
		return true
	}
	if min(line2.x1, line2.x2) <= float64(line1.x1) && float64(line1.x1) <= max(line2.x1, line2.x2) {
		return true
	}
	if min(line1.y1, line1.y2) <= float64(line2.y1) && float64(line2.y1) <= max(line1.y1, line1.y2) {
		return true
	}
	if min(line2.y1, line2.y2) <= float64(line1.y1) && float64(line1.y1) <= max(line2.y1, line2.y2) {
		return true
	}
	return false
}

func min(num1, num2 int) float64 {
	return math.Min(float64(num1), float64(num2))
}

func max(num1, num2 int) float64 {
	return math.Max(float64(num1), float64(num2))
}
