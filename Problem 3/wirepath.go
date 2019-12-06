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
	x1, y1, x2, y2, steps int
}

type movement struct {
	dir, mag int
}

type point struct {
	x, y int
}

func main() {
	data, err := ioutil.ReadFile("inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	var wires []string = strings.Split(string(data), "\n")

	var wirestr1 []string = strings.Split(wires[0], ",")
	var wirestr2 []string = strings.Split(wires[1], ",")
	var wire1 []movement = make([]movement, len(wirestr1))
	var wire2 []movement = make([]movement, len(wirestr2))
	var lines1 []line = make([]line, len(wire1))
	var lines2 []line = make([]line, len(wire2))

	for i, v := range wirestr1 {
		wire1[i] = movement{calcdir(string(v[0])), sliceindexconvert(v[1:])}
	}

	for i, v := range wirestr2 {
		wire2[i] = movement{calcdir(string(v[0])), sliceindexconvert(v[1:])}
	}

	lines1 = calcline(wire1)
	lines2 = calcline(wire2)

	//Find shortest distance

	var shortest int = 0

	for _, v1 := range lines1 {
		for _, v2 := range lines2 {
			if lineintersects(v1, v2) {
				x, y := findintersect(v1, v2)

				if x != 0 || y != 0 {
					if shortest == 0 {
						shortest = abs(x) + abs(y)
					} else {
						shortest = int(min(shortest, abs(x)+abs(y)))
					}
				}
			}
		}
	}

	fmt.Println(strconv.Itoa(shortest))

	//Find shortest steps

	shortest = 0

	for _, v1 := range lines1 {
		for _, v2 := range lines2 {
			if lineintersects(v1, v2) {
				steps := calculatesteps(v1, v2)

				if steps != 0 {
					if shortest == 0 {
						shortest = steps
					} else {
						shortest = int(min(shortest, steps))
					}
				}
			}
		}
	}

	fmt.Println(strconv.Itoa(shortest))

}

func calculatesteps(line1, line2 line) int {
	var steps int = 0
	steps += line1.steps + line2.steps

	if line1.x1 == line1.x2 {
		if line2.x1 == line2.x2 {
			steps += abs(line2.x1 - line1.x1)
		}
		if line2.y1 == line2.y2 {
			steps += abs(line1.y1-line2.y1) + abs(line2.x1-line1.x1)
		}
	}
	if line1.y1 == line1.y2 {
		if line2.x1 == line2.x2 {
			steps += abs(line2.y1-line1.y1) + abs(line1.x1-line2.x1)
		}
		if line2.y1 == line2.y2 {
			steps += abs(line2.y1 - line1.y1)
		}
	}

	return steps
}

func findintersect(line1, line2 line) (int, int) {
	if line1.x1 == line1.x2 {
		if line2.x1 == line2.x2 {
			if max(line1.y1, line2.y1) < 0 && 0 < min(line1.y2, line2.y2) {
				return line1.x1, 0
			}
			if math.Copysign(1, float64(line1.y1)) == -1 {
				return line1.x1, int(max(line1.y1, line2.y1))
			} else if math.Copysign(1, float64(line1.x1)) == 1 {
				return line1.x1, int(min(line1.y1, line2.y1))
			}
		}
		if line2.y1 == line2.y2 {
			return line1.x1, line2.y1
		}
	}
	if line1.y1 == line1.y2 {
		if line2.x1 == line2.x2 {
			return line2.x1, line1.y1
		}
		if line2.y1 == line2.y2 {
			if max(line1.x1, line2.x1) < 0 && 0 < min(line1.x2, line2.x2) {
				return 0, line1.y1
			}
			if math.Copysign(1, float64(line1.x1)) == -1 {
				return int(max(line1.x1, line2.x1)), line1.y1
			} else if math.Copysign(1, float64(line1.x1)) == 1 {
				return int(min(line1.x1, line2.x1)), line1.y1
			}

		}
	}
	return 0, 0
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
	var steps int = 0
	var lines []line
	lines = make([]line, len(wire))

	for i, v := range wire {
		x2 = x1 + (v.dir%3)*v.mag
		y2 = y1 + (v.dir%2)*v.mag
		lines[i] = line{x1, y1, x2, y2, steps}
		x1 = x2
		y1 = y2
		steps += abs(v.mag)
	}

	return lines
}

func lineintersects(line1, line2 line) bool {
	// If thing is congruent
	var xcongruent bool = false
	var ycongruent bool = false

	if min(line1.x1, line1.x2) <= float64(line2.x1) && float64(line2.x1) <= max(line1.x1, line1.x2) {
		xcongruent = true
	}
	if min(line2.x1, line2.x2) <= float64(line1.x1) && float64(line1.x1) <= max(line2.x1, line2.x2) {
		xcongruent = true
	}

	if min(line1.y1, line1.y2) <= float64(line2.y1) && float64(line2.y1) <= max(line1.y1, line1.y2) {
		ycongruent = true
	}
	if min(line2.y1, line2.y2) <= float64(line1.y1) && float64(line1.y1) <= max(line2.y1, line2.y2) {
		ycongruent = true
	}
	if xcongruent && ycongruent {
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

func abs(num1 int) int {
	return int(math.Abs(float64(num1)))
}

func sign(num1 int) int {
	return int(math.Copysign(1, float64(num1)))
}
