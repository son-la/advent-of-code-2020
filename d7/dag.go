package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// Vertex is the interface to be implemented for the vertices of the DAG.
type Vertex interface {
	Id() string
}

type DAG struct {
	vertexIds        map[string]Vertex
	inboundEdge      map[Vertex]map[Vertex]int
	outboundEdge     map[Vertex]map[Vertex]int
}

func NewDAG() *DAG {
	return &DAG{
		vertexIds:        make(map[string]Vertex),
		inboundEdge:      make(map[Vertex]map[Vertex]int),
		outboundEdge:     make(map[Vertex]map[Vertex]int),
	}
}

func (d *DAG) AddVertex(v Vertex) error {
	if _, exists := d.vertexIds[v.Id()]; exists {
		return DAGError{v}
	}

	d.vertexIds[v.Id()] = v

	return nil
}

func (d *DAG) GetVertex(id string) (Vertex, bool) {
	if vertex, ok := d.vertexIds[id]; ok {
		return vertex, true
	} else {
		return nil, false
	}
	
}

func (d *DAG) GetAncestors(v Vertex) map[Vertex]bool {
	ancestors := make(map[Vertex]bool)

	if parents, ok := d.inboundEdge[v]; ok {
		for parent := range parents {
			ancestors[parent] = true

			parentAncestors := d.GetAncestors(parent)
			for ancestor := range parentAncestors {
				ancestors[ancestor] = true
			}
		}
	}

	return ancestors
}

func (d *DAG) GetDistanceToLeaves(v Vertex) map[Vertex]int {
	distanceMap := make(map[Vertex]int)
	return distanceMap
}

func (d *DAG) AddEdge(src Vertex, dst Vertex, weight int) error {
	if _, ok := d.vertexIds[src.Id()]; !ok {
		d.AddVertex(src)
	}
	if _, ok := d.vertexIds[dst.Id()]; !ok{
		d.AddVertex(dst)
	}

	if _, exists := d.inboundEdge[dst]; !exists {
		d.inboundEdge[dst] = make(map[Vertex]int)
	}

	if _, exists := d.outboundEdge[src]; !exists {
		d.outboundEdge[src] = make(map[Vertex]int)
	}

	d.outboundEdge[src][dst] = weight	
	d.inboundEdge[dst][src] = weight

	return nil
}

func parseDAGFromFile(fileName string) *DAG {
	d1 := NewDAG()

	file, err := os.Open(fileName)
    if err != nil {
		fmt.Println("Error reading file")
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	pattern1 := regexp.MustCompile(`(.*) bags contain (.*)`)
	pattern2 := regexp.MustCompile(`(\d+) (.*?) bag(s?)`)

	// Read line to slice
    for scanner.Scan() {
		line := scanner.Text()		
		segs := pattern1.FindAllStringSubmatch(line,1)
		
		src, exists := d1.GetVertex(segs[0][1])
		if !exists {
			src = &Bag{segs[0][1]}
		}	

		segs = pattern2.FindAllStringSubmatch(segs[0][2], -1)
	
		for _, value := range segs {
			dst, exists := d1.GetVertex(value[2])
			if !exists {
				dst = &Bag{value[2]}
			}
			
			weight, _ := strconv.Atoi(value[1])
			_ = d1.AddEdge(src, dst, weight)
		}
	}

	return d1
}

type DAGError struct {
	v Vertex
}

// Implements the error interface.
func (e DAGError) Error() string {
	return fmt.Sprintf("Error!")
}
