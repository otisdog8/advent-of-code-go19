package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("inputs.txt")
	sum := 0

	if err != nil {
		log.Fatal(err)
	}

	dataparts := strings.Split(string(data), "\r\n")

	for i := 0; i < len(dataparts); i++ {
		num, err := strconv.Atoi(dataparts[i])
		if err != nil {
			log.Fatal(err)
		}
		sum += int(math.Floor(float64(num)/3)) - 2
		fuel := int(math.Floor(float64(num)/3)) - 2

		for int(math.Floor(float64(fuel)/3))-2 > 0 {

			sum += int(math.Floor(float64(fuel)/3)) - 2
			fuel = int(math.Floor(float64(fuel)/3)) - 2

		}
	}

	fmt.Println(strconv.Itoa(sum))

	fuel := sum

	for int(math.Floor(float64(fuel)/3))-2 > 0 {

		sum += int(math.Floor(float64(fuel)/3)) - 2
		fuel = int(math.Floor(float64(fuel)/3)) - 2

	}

	fmt.Println(strconv.Itoa(sum))

}
