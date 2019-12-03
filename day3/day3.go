package main

import (
	"log"
	"os"
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"math"
)

type Vertex struct {
	X, Y int
}

type Line struct {
	Begin, End Vertex
	NumSteps int
}

func createLine(begin Vertex, move string, beginNumSteps int) Line {
	var end Vertex
	var l Line
	distance, _ := strconv.Atoi(move[1:])
	switch move[0] {
	case 'R':
		end.X = begin.X + distance
		end.Y = begin.Y
	case 'L':
		end.X = begin.X - distance
		end.Y = begin.Y
	case 'U':
                end.Y = begin.Y + distance
                end.X = begin.X
	case 'D':
                end.Y = begin.Y - distance
                end.X = begin.X
	}
	l.Begin = begin
	l.End = Vertex{ end.X, end.Y }
	l.NumSteps = beginNumSteps + int(math.Abs(float64(distance)))
	return l
}

func between(p int, v1 int, v2 int) bool {
	return (v1 <= p && p <= v2) || (v2 <= p && p <= v1)
}

func finalNumSteps(l Line, vert Vertex) int {
	var diff int
	if l.Begin.X != l.End.X {
		diff = int(math.Abs(float64(l.End.X - vert.X)))
	} else {
		diff = int(math.Abs(float64(l.End.Y - vert.Y)))
	}
	return l.NumSteps - diff
}

func compareLines(line1 Line, line2 Line) (Vertex, bool, int) {
	if between(line1.Begin.X, line2.Begin.X, line2.End.X) && between(line2.Begin.Y, line1.Begin.Y, line1.End.Y) {
		vert := Vertex{ line1.Begin.X, line2.Begin.Y }
		return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
	} else if between(line1.End.X, line2.Begin.X, line2.End.X) && between(line2.Begin.Y, line1.Begin.Y, line1.End.Y) {
		vert := Vertex{ line1.End.X, line2.Begin.Y }
		return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
	} else if between(line2.Begin.X, line1.Begin.X, line1.End.X) && between(line1.Begin.Y, line2.Begin.Y, line2.End.Y) {
		vert := Vertex{ line2.Begin.X, line1.Begin.Y }
		return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
	} else if between(line2.End.X, line1.Begin.X, line1.End.X) && between(line1.Begin.Y, line2.Begin.Y, line2.End.Y) {
		vert := Vertex{ line2.End.X, line1.Begin.Y }
                return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
	} else if between(line1.Begin.Y, line2.Begin.Y, line2.End.Y) && between(line2.Begin.X, line1.Begin.X, line1.End.X) {
		vert := Vertex{ line2.Begin.X, line1.Begin.Y }
                return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
        } else if between(line1.End.Y, line2.Begin.Y, line2.End.Y) && between(line2.Begin.X, line1.Begin.X, line1.End.X) {
		vert := Vertex{ line2.End.X, line1.Begin.Y }
                return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
        } else if between(line2.Begin.Y, line1.Begin.Y, line1.End.Y) && between(line1.Begin.X, line2.Begin.X, line2.End.X) {
		vert := Vertex{ line1.Begin.X, line2.Begin.Y }
                return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
        } else if between(line2.End.Y, line1.Begin.Y, line1.End.Y) && between(line1.Begin.X, line2.Begin.X, line2.End.X) {
		vert := Vertex{ line1.End.X, line2.Begin.Y }
                return vert, true, finalNumSteps(line1, vert) + finalNumSteps(line2, vert)
	} else {
		return Vertex {}, false, 0
	}
}

func main() {
        var lines [][]Line
        file, err := os.Open("t")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for n := 0; n < 2; n++ {
                scanner.Scan()
                pathString := scanner.Text()
                path := strings.Split(pathString, ",")
                end := Vertex{ 0, 0 }
		endNumSteps := 0
                var l []Line
                for i, move := range path {
                        l = append(l, createLine(end, move, endNumSteps))
                        end = l[i].End
			endNumSteps = l[i].NumSteps
                }
                lines = append(lines, l)
        }
        for _, l1 := range lines[0] {
		for _, l2 := range lines[1] {
			vert, truth, numSteps := compareLines(l1, l2)
			if truth {
				fmt.Println(fmt.Sprint(vert) + "-" + fmt.Sprint(numSteps))
			}
		}
        }
}
