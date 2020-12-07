package main

import (
	"bufio"
	"os"
	"fmt"
	"regexp"
	"github.com/heimdalr/dag"
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
	d := dag.NewDAG()

	file, err := os.Open("input.txt")
    if err != nil {
		fmt.Println("Error reading file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	pattern1 := regexp.MustCompile(`(.*) bags contain (.*)`)
	pattern2 := regexp.MustCompile(`(\d+) (.*?) bag(s?)`)

	// Read line to slice
    for scanner.Scan() {
		line := scanner.Text()		
		segs := pattern1.FindAllStringSubmatch(line,1)
		
		// fmt.Println(segs[0][2])
		b, err := d.GetVertex(segs[0][1])
		if err != nil {
			b = &Bag{segs[0][1]}
		}
		
		segs = pattern2.FindAllStringSubmatch(segs[0][2], -1)

		// fmt.Println(line)		
		for _, value := range segs {
			b1, err := d.GetVertex(value[2])
			if err != nil {
				b1 = &Bag{value[2]}
			}

			_ = d.AddEdge(b, b1)
		}
	}

	b, err := d.GetVertex("shiny gold")
	x, _ := d.GetAncestors(b)
	
	cnt := 0
	for _, v := range x {
		if v == true {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}