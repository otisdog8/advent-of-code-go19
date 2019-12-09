package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 int = 382345
	var input2 int = 843167
	var sum int = 0

	for i := input1; i < input2+1; i++ {
		if numberpasses(i) {
			sum++
		}
	}

	fmt.Println(strconv.Itoa(sum))

	sum = 0

	for i := input1; i < input2+1; i++ {
		if numberpassesv2(i) {
			sum++
		}
	}

	fmt.Println(strconv.Itoa(sum))

}

func numberpassesv2(num int) bool {
	ispair := false
	works := true

	for i := 0; i < 5; i++ {
		work := checktwo(getpair(num, i))
		if !work {
			works = false
		}
	}

	ispair = checkpair(num)

	if !ispair {
		works = false
	}

	return works
}

func checktwo(pair int) bool {
	if pair%10 < (pair-pair%10)/10 {
		return false
	}

	return true
}

func checkpair(num int) bool {
	pair := false
	var pairs []int = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		if ispair(getpair(num, i)) {
			pairs[getpair(num, i)%10]++
		} else {
			for i, v := range pairs {
				if v > 1 {
					pairs[i] = 0
				}
			}
		}
	}

	for _, v := range pairs {
		if v == 1 {
			pair = true
		}
	}

	return pair
}

func ispair(pair int) bool {
	if pair%10 == (pair-pair%10)/10 {
		return true
	}
	return false
}

func numberpasses(num int) bool {
	ispair := false
	works := true

	for i := 0; i < 5; i++ {
		work, pair := paircheck(getpair(num, i))

		if pair {
			ispair = true
			//fmt.Println(strconv.Itoa(num))
		}
		if !work {
			works = false
		}
	}

	if !ispair {
		works = false
	}

	return works
}

func paircheck(pair int) (bool, bool) {
	ispair := false
	works := true

	if pair%10 == (pair-pair%10)/10 {
		ispair = true
	}
	if pair%10 < (pair-pair%10)/10 {
		works = false
	}
	return works, ispair
}

func getpair(num, index int) int {
	// Assumes a six digit number
	str := strconv.Itoa(num)

	strslice := str[index : index+2]

	//fmt.Println(strslice)

	result, _ := strconv.Atoi(strslice)

	return result
}
