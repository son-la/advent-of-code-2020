package main

import (
    "bufio"
    "fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func solution1(fileName string) int {
	file, err := os.Open(fileName)
    if err != nil {
		fmt.Println("Error reading file")
        return 0
	}
	
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validPassword := 0

	// Pattern to match line
	pattern := regexp.MustCompile(`(.*)-(.*) (.*):(.*)`)

    for scanner.Scan() {
		line := scanner.Text()		
		segs := pattern.FindAllStringSubmatch(line,-1)
		
		min, _ := strconv.Atoi(segs[0][1])
		max, _ := strconv.Atoi(segs[0][2])
		occurence := len(regexp.MustCompile(segs[0][3]).FindAllString(segs[0][4], -1))
		if (occurence <= max && occurence >= min) {
			validPassword += 1
		}
	}

	return validPassword
}

func main() {
	record := time.Now()
	fmt.Println("Solution 1 find: ", solution1("input.txt"), "in ", time.Since(record))
}