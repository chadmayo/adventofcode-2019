package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
)

func main() {
	var s int
	file, err := os.Open("t")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		added := (i / 3) - 2
		s += added
		for {
			added = (added / 3) - 2
			if added < 0 {
				break
			}
			s += added
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
