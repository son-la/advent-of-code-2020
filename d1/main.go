package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
	"time"
)

func solution1(fileName string, value int) int {
	file, err := os.Open(fileName)
    if err != nil {
		fmt.Println("Error reading file")
        return 0
	}
	
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	
	// Read line to slice
    for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Error reading file")
        	return 0
		}
		
		lines = append(lines, line)
	}

	for _, x := range lines {
		for _,y := range lines {
			if (x+y == value) {
				return x*y
			}
		}
	}

	return 0
}

func solution2(fileName string, value int) int {
	file, err := os.Open(fileName)
    if err != nil {
		fmt.Println("Error reading file")
        return 0
	}
	
	defer file.Close()

	lines := make(map[int]bool)
	scanner := bufio.NewScanner(file)
	
	// Read line to slice
    for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Error reading file")
        	return 0
		}

		_, ok := lines[value - x]
		if (ok) {
			return x*(value-x)
		} else 
		{
			lines[x] = true
		}
	}

	return 0
}

func main() {
	record := time.Now()
	fmt.Println("Solution 1 find: ", solution1("input.txt", 2020), "in ", time.Since(record))

	record = time.Now()
	fmt.Println("Solution 2 find: ", solution2("input.txt", 2020), "in ", time.Since(record))
}