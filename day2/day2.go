package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
)

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

	s_orig := strings.Split(string(buffer), ",")
	s := make([]string, len(s_orig))
	for x := 0; x<=99; x++ {
		for y := 0; y<=99; y++ {
			copy(s, s_orig)
			s[1] = strconv.Itoa(x)
			s[2] = strconv.Itoa(y)
			for i := 0; i < len(s)/4; i++ {
				opcode := s[i*4]
				if opcode == "99" {
					if s[0] == "19690720" {
						fmt.Println(x)
						fmt.Println(y)
						os.Exit(0)
					}
					//fmt.Println(s[0])
					//os.Exit(0)
					break
				}
				position1, _ := strconv.Atoi(s[i*4+1])
				valueAtPosition1, _ := strconv.Atoi(s[position1])
				position2, _ := strconv.Atoi(s[i*4+2])
				valueAtPosition2, _ := strconv.Atoi(s[position2])
				target, _ := strconv.Atoi(s[i*4+3])
				var result int
				if opcode == "1" {
					result = valueAtPosition1 + valueAtPosition2
				} else if opcode == "2" {
					result = valueAtPosition1 * valueAtPosition2
				} else {
					fmt.Println("Error: invalid opp code: " + s[i*4])
					os.Exit(1)
				}
				//fmt.Println(opcode + "-" + s[i*4+1] + "-" + string(valueAtPosition2) + "-" + string(target) + ": " + string(result))
				s[target] = strconv.Itoa(result)
			}
		}
	}
}
