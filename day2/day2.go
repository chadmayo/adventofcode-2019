package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func addOp(position1 int, position2 int, target int, program []string) int {
	valueAtPosition1, _ := strconv.Atoi(program[position1])
	valueAtPosition2, _ := strconv.Atoi(program[position2])
	r := valueAtPosition1 + valueAtPosition2
	program[target] = strconv.Itoa(r)
	return 4
}

func multiplyOp(position1 int, position2 int, target int, program []string) int {
	valueAtPosition1, _ := strconv.Atoi(program[position1])
	valueAtPosition2, _ := strconv.Atoi(program[position2])
	r := valueAtPosition1 * valueAtPosition2
	program[target] = strconv.Itoa(r)
	return 4
}

func main() {
	//var s int
	file, err := os.Open("t")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)
	_, err = file.Read(buffer)

	if err != nil {
		log.Fatal(err)
	}

	program_orig := strings.Split(string(buffer), ",")
	program := make([]string, len(program_orig))
	for x := 0; x<=99; x++ {
		for y := 0; y<=99; y++ {
			copy(program, program_orig)
			program[1] = strconv.Itoa(x)
			program[2] = strconv.Itoa(y)
			i := 0
			for i < len(program) {
				opcode := program[i]
				if opcode == "99" {
					if program[0] == "19690720" {
						fmt.Println(x)
						fmt.Println(y)
						os.Exit(0)
					}
					break
				}
				position1, _ := strconv.Atoi(program[i+1])
				position2, _ := strconv.Atoi(program[i+2])
				target, _ := strconv.Atoi(program[i+3])
				var opLength int
				if opcode == "1" {
					opLength = addOp(position1, position2, target, program)
				} else if opcode == "2" {
					opLength = multiplyOp(position1, position2, target, program)
				} else {
					fmt.Println("Error: invalid opp code: " + program[i])
					os.Exit(1)
				}
				i += opLength
			}
		}
	}
}
