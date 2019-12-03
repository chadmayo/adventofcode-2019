package main

import (
	"fmt"
	"testing"
)

func TestCreateLine(t *testing.T) {
	b := Vertex{ -8, 50 }
	l := createLine(b, "R105", 0)
	expect := fmt.Sprint(Vertex{ 97, 50 })
	got := fmt.Sprint(l.End)
	if expect != got {
		t.Error("Expected " + expect + " but got ", got)
	}
}

func TestBetween(t *testing.T) {
	trueCases := [10][3]int {
		{0,-1,1},
		{-18,-19,-4},
		{555,1,555},
		{4,4,4},
		{0,0,5},
		{1,2,-4},
		{-33,1,-33},
		{500,1000,-1000},
		{60,60,-1000},
		{1,18,0},
	}
	falseCases := [10][3]int {
                {0,1,1},
                {18,-19,-4},
                {555,1,554},
                {4,5,5},
                {0,1,5},
                {1,2,4},
                {-33,-32,33},
                {-1001,1000,-1000},
                {60,59,-1000},
                {1,18,2},
        }
	for _, v := range trueCases {
		if !between(v[0], v[1], v[2]) {
			t.Error("Expected true but got false", fmt.Sprint(v))
		}
	}
	for _, v := range falseCases {
                if between(v[0], v[1], v[2]) {
                        t.Error("Expected false but got true", fmt.Sprint(v))
                }
        }
}

func TestCompareLines(t *testing.T) {
	type TestCase struct {
		L1 Line
		L2 Line
		Intersection Vertex
		FinalNumSteps int
	}
	trueCases := []TestCase {
		{
			Line {
				Vertex{ 1, 5 },
				Vertex{ 4, 5 },
				1,
			},
			Line {
                                Vertex{ 5, 5 },
                                Vertex{ 4, 5 },
				500,
                        },
			Vertex{ 4, 5 },
			501,
		},
		{
                        Line {
                                Vertex{ -5, 5 },
                                Vertex{ -4, 5 },
				50,
                        },
                        Line {
                                Vertex{ -1, 5 },
                                Vertex{ -4, 5 },
				20,
                        },
                        Vertex{ -4, 5 },
			70,
                },
		{
                        Line {
                                Vertex{ -5, 4 },
                                Vertex{ 5, 4 },
				15,
                        },
                        Line {
                                Vertex{ 1, 5 },
                                Vertex{ 1, -5 },
				20,
                        },
			Vertex{ 1, 4 },
			22,
                },
		{
                        Line {
                                Vertex{ 5, 0 },
                                Vertex{ 5, 0 },
				20,
                        },
                        Line {
                                Vertex{ 1, 0 },
                                Vertex{ 10, 0 },
				20,
                        },
                        Vertex{ 5, 0 },
			35,
                },
		{
                        Line {
                                Vertex{ -5, 1 },
                                Vertex{ 5, 1 },
				20,
                        },
                        Line {
                                Vertex{ 1, 5 },
                                Vertex{ 1, -5 },
				15,
                        },
                        Vertex{ 1, 1 },
			25,
                },
		{
                        Line {
                                Vertex{ 6, 5 },
                                Vertex{ 6, -5 },
				12,
                        },
                        Line {
                                Vertex{ 6, 1 },
                                Vertex{ -5, 1 },
				20,
                        },
                        Vertex{ 6, 1 },
			15,
                },
	}
	falseCases := []TestCase {
                {
                        Line {
                                Vertex{ 1, 5 },
                                Vertex{ 4, 5 },
				0,
                        },
                        Line {
                                Vertex{ 5, 5 },
                                Vertex{ 5, 5 },
				0,
                        },
                        Vertex{},
			0,
                },
		{
                        Line {
                                Vertex{ -5, 4 },
                                Vertex{ 5, 4 },
				0,
                        },
                        Line {
                                Vertex{ 5, 8 },
                                Vertex{ -5, 8 },
				0,
                        },
                        Vertex{},
			0,
                },
		{
                        Line {
                                Vertex{ -6, -5 },
                                Vertex{ -6, 5 },
				0,
                        },
                        Line {
                                Vertex{ 8, 6 },
                                Vertex{ -6, 6 },
				0,
                        },
                        Vertex{},
			0,
                },
	}
	for _, v := range trueCases {
		vert, truth, numSteps := compareLines(v.L1, v.L2)
		if !truth || v.Intersection != vert || v.FinalNumSteps != numSteps {
			t.Error(v, vert, truth, numSteps)
		}
	}
	for _, v := range falseCases {
                vert, truth, numSteps := compareLines(v.L1, v.L2)
                if truth || v.Intersection != vert || v.FinalNumSteps != numSteps {
                        t.Error(v, vert, truth, numSteps)
                }
        }
}
