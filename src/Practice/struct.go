package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID                int
	Name, Address     string
	DoB               time.Time
	Position          string
	Salary, ManagerID int
}

type address struct {
	hostname string
	port     int
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
