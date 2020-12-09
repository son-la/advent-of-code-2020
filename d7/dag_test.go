
package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	// initialize a new graph
	d1 := parseDAGFromFile("test1.txt")

	b, _ := d1.GetVertex("shiny gold")
	x:= d1.GetAncestors(b)
	if len(x) != 4 {
		t.Errorf("GetAncestors() = %d, want 4", len(x))
	}
	
}

func TestPart2_input1(t *testing.T) {
	// initialize a new graph
	d1 := parseDAGFromFile("test1.txt")

	b, _ := d1.GetVertex("shiny gold")
	distanceMap := d1.GetDistanceFromVertex(b)
	totalDistance := 0
	for _, distance := range distanceMap {
		totalDistance += distance
	}

	if totalDistance != 32 {
		t.Errorf("GetDistanceFromVertex() = %d, want 32", totalDistance)
	}
}

func TestPart2_input2(t *testing.T) {
	// initialize a new graph
	d1 := parseDAGFromFile("test2.txt")

	b, _ := d1.GetVertex("shiny gold")
	distanceMap := d1.GetDistanceFromVertex(b)
	totalDistance := 0
	for _, distance := range distanceMap {
		totalDistance += distance	
	}

	if totalDistance != 126 {
		t.Errorf("GetDistanceFromVertex() = %d, want 126", totalDistance)
	}
}