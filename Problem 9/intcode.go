package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
)

var memory []*big.Int
var inputs []*big.Int
var inputptr int
var outputs []*big.Int
var outputptr int
var relptr int

var globaldump string = ""

func main() {
	setupvm()
}

//Global Scope Management

func setupvm() {
	var memory []*big.Int = readmem("inputs.txt")
	loadmem(memory)

	//VM runtime
	inputs = make([]*big.Int, 1, 1)
	inputptr = 0
	outputs = make([]*big.Int, 0, 10)
	outputptr = 0
	relptr = 0

	inputs[0] = tobig(2)

	vm()

	//f, _ := os.Create("file2.txt")
	//f.WriteString(globaldump)
	//f.Sync()

	for _, v := range outputs {
		printbig(v)
	}

}

func readmem(filename string) []*big.Int {
	// Setup
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	dataparts := strings.Split(string(data), ",")

	var length int = len(dataparts)
	var vmmemory []*big.Int = make([]*big.Int, length, length)

	for i, v := range dataparts {
		temp, err := strconv.Atoi(v)
		if err != nil {
			log.Println(err.Error())
		}
		vmmemory[i] = tobig(temp)
	}

	return vmmemory
}

func loadmem(mem []*big.Int) {
	memory = mem
}

func loadinput(input []*big.Int) {
	inputs = input
}

func getoutput() []*big.Int {
	return outputs
}

func vm() {
	var instptr int = 0
	var exit bool = false

	for !exit {
		instptr, exit = processopcode(instptr)
	}
}

func processopcode(ptr int) (int, bool) {
	var retval bool = false // To exit or not
	var tempptr int = 0
	// Make mode and inputs an int array?
	// Change instruction ptr in here too?
	if ptrinst(ptr, 1) {
		retval, tempptr = opcode1(ptr)
	} else if ptrinst(ptr, 2) {
		retval, tempptr = opcode2(ptr)
	} else if ptrinst(ptr, 3) {
		retval, tempptr = opcode3(ptr)
	} else if ptrinst(ptr, 4) {
		retval, tempptr = opcode4(ptr)
	} else if ptrinst(ptr, 5) {
		retval, tempptr = opcode5(ptr)
	} else if ptrinst(ptr, 6) {
		retval, tempptr = opcode6(ptr)
	} else if ptrinst(ptr, 7) {
		retval, tempptr = opcode7(ptr)
	} else if ptrinst(ptr, 8) {
		retval, tempptr = opcode8(ptr)
	} else if ptrinst(ptr, 9) {
		retval, tempptr = opcode9(ptr)
	} else if ptrinst(ptr, 99) {
		retval = true
		tempptr = 1
	} else {
		printint(ptr)
		printbig(getval(ptr))
	}

	return ptr + tempptr, retval
}

func ptrinst(index, value int) bool {
	return big.NewInt(0).Mod(getval(index), big.NewInt(100)).Cmp(big.NewInt(int64(value))) == 0
}

func intindex(val *big.Int, index int) int {
	var num int = int(val.Int64())

	return (num%pow(10, index+1) - num%pow(10, index)) / pow(10, index)
}

func opcode000(ptr int) (bool, int) {
	//Read and parse instruction
	//Get values needed for instruction
	//Do instruction
	//Return result
	return false, 1
}
func opcode1(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)
	var param2 int = intindex(opcode, 3)
	var param3 int = intindex(opcode, 4)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)
	var paramval2 *big.Int = fetchvalue(getval(ptr+2), param2)
	var paramval3 int = toint(fetchwriteval(getval(ptr+3), param3))

	setval(paramval3, tobig(0).Add(paramval1, paramval2))

	return false, 4
}

func opcode2(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)
	var param2 int = intindex(opcode, 3)
	var param3 int = intindex(opcode, 4)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)
	var paramval2 *big.Int = fetchvalue(getval(ptr+2), param2)
	var paramval3 int = toint(fetchwriteval(getval(ptr+3), param3))

	setval(paramval3, tobig(0).Mul(paramval1, paramval2))

	return false, 4
}

func opcode3(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)

	var paramval1 int = toint(fetchwriteval(getval(ptr+1), param1))

	setval(paramval1, inputs[inputptr])
	inputptr++
	return false, 2
}

func opcode4(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)

	outputs = append(outputs, paramval1)
	return false, 2
}

func opcode5(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)
	var param2 int = intindex(opcode, 3)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)
	var paramval2 *big.Int = fetchvalue(getval(ptr+2), param2)

	if paramval1.Cmp(tobig(0)) == 0 {
		ptr = 3
	} else {
		ptr = -1 * (ptr - toint(paramval2))
	}

	return false, ptr
}

func opcode6(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)
	var param2 int = intindex(opcode, 3)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)
	var paramval2 *big.Int = fetchvalue(getval(ptr+2), param2)

	if paramval1.Cmp(tobig(0)) == 0 {
		ptr = -1 * (ptr - toint(paramval2))
	} else {
		ptr = 3
	}

	return false, ptr
}

func opcode7(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)
	var param2 int = intindex(opcode, 3)
	var param3 int = intindex(opcode, 4)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)
	var paramval2 *big.Int = fetchvalue(getval(ptr+2), param2)
	var paramval3 int = toint(fetchwriteval(getval(ptr+3), param3))

	if paramval1.Cmp(paramval2) == -1 {
		setval(paramval3, tobig(1))
	} else {
		setval(paramval3, tobig(0))
	}

	return false, 4
}

func opcode8(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)
	var param2 int = intindex(opcode, 3)
	var param3 int = intindex(opcode, 4)

	var paramval1 *big.Int = fetchvalue(getval(ptr+1), param1)
	var paramval2 *big.Int = fetchvalue(getval(ptr+2), param2)
	var paramval3 int = toint(fetchwriteval(getval(ptr+3), param3))

	if paramval1.Cmp(paramval2) == 0 {
		setval(paramval3, tobig(1))
	} else {
		setval(paramval3, tobig(0))
	}

	return false, 4
}

func opcode9(ptr int) (bool, int) {
	var opcode *big.Int = getval(ptr)

	var param1 int = intindex(opcode, 2)

	var paramval1 int = toint(fetchvalue(getval(ptr+1), param1))

	relptr += paramval1
	return false, 2
}

func fetchvalue(address *big.Int, mode int) *big.Int {

	var retval *big.Int = big.NewInt(int64(0))
	if mode == 0 {
		retval = getval(int(address.Int64()))
	} else if mode == 1 {
		retval = address
	} else if mode == 2 {
		retval = getval(int(address.Int64()) + relptr)
	}

	return retval
}

func fetchwriteval(address *big.Int, mode int) *big.Int {
	var retval *big.Int = big.NewInt(int64(0))
	if mode == 0 {
		retval = address
	} else if mode == 1 {
		retval = address
	} else if mode == 2 {
		retval = retval.Add(address, tobig(relptr))
	}

	return retval
}

func getval(index int) *big.Int {
	if index+1 > cap(memory) {
		newmem := make([]*big.Int, index+1, index+1)
		copy(newmem, memory)
		memory = newmem
	}

	if memory[index] == nil {
		memory[index] = tobig(0)
	}
	return memory[index]
}

func setval(index int, value *big.Int) {
	if index+1 > cap(memory) {
		newmem := make([]*big.Int, index+1, index+1)
		copy(newmem, memory)
		memory = newmem
	}

	memory[index] = value
}

func dumpmem() {
	for i := 0; i < len(memory); i++ {
		globaldump += strconv.Itoa(toint(memory[i])) + "\n"
	}
}

func pow(base, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

func log10(num int) int {
	return int(math.Floor(math.Log10(float64(num))))
}

func toint(num *big.Int) int {
	return int(num.Int64())
}

func tobig(num int) *big.Int {
	return big.NewInt(int64(num))
}

func printbig(num *big.Int) {
	printint(toint(num))
}

func printint(num int) {
	fmt.Println(strconv.Itoa(num))
}
