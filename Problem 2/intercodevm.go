package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			vm(i, j)
		}
	}
}

func vm(noun, verb int) {

	// Setup
	data, err := ioutil.ReadFile("inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	dataparts := strings.Split(string(data), ",")

	var length int = len(dataparts)
	var vmmemory = make([]int, length)
	var instptr int = 0

	for i, v := range dataparts {
		vmmemory[i], _ = strconv.Atoi(v)
	}

	vmmemory[1] = noun
	vmmemory[2] = verb

	for !processopcode(vmmemory, instptr) {
		instptr += 4
	}

	if vmmemory[0] == 19690720 {
		fmt.Println(strconv.Itoa(noun))
		fmt.Println(strconv.Itoa(verb))
	}
}

func processopcode(memory []int, ptr int) bool {
	var retval bool = false // To exit or not

	if memory[ptr] == 1 {
		var inaddr1, inaddr2, outaddr1 int
		inaddr1 = memory[ptr+1]
		inaddr2 = memory[ptr+2]
		outaddr1 = memory[ptr+3]
		retval = opcode1(memory, 0, inaddr1, inaddr2, outaddr1)
	} else if memory[ptr] == 2 {
		var inaddr1, inaddr2, outaddr1 int
		inaddr1 = memory[ptr+1]
		inaddr2 = memory[ptr+2]
		outaddr1 = memory[ptr+3]
		retval = opcode2(memory, 0, inaddr1, inaddr2, outaddr1)
	} else if memory[ptr] == 99 {
		retval = true //Exits
	}

	return retval
}

func opcode1(memory []int, mode, inaddr1, inaddr2, outaddr1 int) bool {
	var output int = fetchvalue(memory, inaddr1, mode) + fetchvalue(memory, inaddr2, mode)
	memory[outaddr1] = output
	return false
}

func opcode2(memory []int, mode, inaddr1, inaddr2, outaddr1 int) bool {
	var output int = fetchvalue(memory, inaddr1, mode) * fetchvalue(memory, inaddr2, mode)
	memory[outaddr1] = output
	return false
}

func fetchvalue(memory []int, address, mode int) int {
	var retval int = 0
	if mode == 0 {
		retval = memory[address]
	} else if mode == 1 {
		retval = address
	}

	return retval
}
