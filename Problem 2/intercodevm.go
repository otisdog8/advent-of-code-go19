package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	setupvm()
}

func setupvm() {

	// Setup
	data, err := ioutil.ReadFile("inputs.txt")

	if err != nil {
		log.Fatal(err)
	}

	dataparts := strings.Split(string(data), ",")

	var length int = len(dataparts)
	var vmmemory = make([]int, length)

	for i, v := range dataparts {
		vmmemory[i], _ = strconv.Atoi(v)
		fmt.Println(v)
	}

	vm(vmmemory)

	fmt.Println(strconv.Itoa(vmmemory[0]))
}

func vm(memory []int) []int {
	var instptr int = 0

	for !processopcode(memory, instptr) {
		instptr += 4
	}

	return memory
}

func processopcode(memory []int, ptr int) bool {
	var retval bool = false // To exit or not
	// Make mode and inputs an int array?
	// Change instruction ptr in here too?

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

	printint(memory[6])

	return retval
}

func opcode1(memory []int, mode, inaddr1, inaddr2, outaddr1 int) bool {
	//fmt.Println("1")
	//printint(fetchvalue(memory, inaddr1, mode))
	//printint(fetchvalue(memory, inaddr2, mode))
	var output int = fetchvalue(memory, inaddr1, mode) + fetchvalue(memory, inaddr2, mode)
	memory[outaddr1] = output
	//printint(output)
	return false
}

func opcode2(memory []int, mode, inaddr1, inaddr2, outaddr1 int) bool {
	//fmt.Println("2")
	//printint(fetchvalue(memory, inaddr1, mode))
	//printint(fetchvalue(memory, inaddr2, mode))
	var output int = fetchvalue(memory, inaddr1, mode) * fetchvalue(memory, inaddr2, mode)
	memory[outaddr1] = output
	//printint(output)
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

func printint(num int) {
	fmt.Println(strconv.Itoa(num))
}
