package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	min, _ := strconv.Atoi(os.Args[1])
	max, _ := strconv.Atoi(os.Args[2])
	count := 0
	fmt.Println(min)
	fmt.Println(max)
	for min <= max {
		bs := []byte(strconv.Itoa(min))
		adjacentDigits := make(map[int]int)
		neverDecreases := true
		previousDigit := -1
		for _, v := range bs {
			digit, _ := strconv.Atoi(string(v))
			if digit == previousDigit {
				adjacentDigits[digit]++
			}
			if digit < previousDigit {
				neverDecreases = false
				break
			}
			//fmt.Println(fmt.Sprint(previousDigit) + "|" + fmt.Sprint(digit) + "|" + fmt.Sprint(neverDecreases) + "|" + fmt.Sprint(hasAdjacentDigits))
			previousDigit = digit
		}
		if neverDecreases {
			//fmt.Println(string(bs))
			for _, v := range adjacentDigits {
				if v == 1 {
					count++
					break
				}
			}
		}
		min++
	}
	fmt.Println(count)
}
