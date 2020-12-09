package main

import (
	"fmt"
)

// data structure that will be used as vertex in the graph
type Bag struct {
	Label string
}

// implement the Vertex interface
func (bag Bag) String() string {
	return bag.Label
}

// implement the Vertex interface
func (bag Bag) Id() string {
	return bag.Label
}

func main() {
	// initialize a new graph
	d1 := parseDAGFromFile("input.txt")

	b, _ := d1.GetVertex("shiny gold")
	x:= d1.GetAncestors(b)
	fmt.Println("Number of ancestors of shiny gold:", len(x))

	distanceMap := d1.GetDistanceFromVertex(b)
	totalDistance := 0
	for _, distance := range distanceMap {
		totalDistance += distance
	}

	fmt.Println("Total distance from shiny gold:", totalDistance)
}