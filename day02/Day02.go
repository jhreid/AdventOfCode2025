package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.ReadFile("input02.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := string(file)
	ranges := strings.Split(input, ",")
	total := 0
	for _, r := range ranges {
		if r == "" {
			break
		}

		limits := strings.Split(r, "-")
		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])

		for i := min; i <= max; i++ {
			check := strconv.Itoa(i)
			l := len(check)
			if l%2 > 0 {
				continue
			}

			first := check[0 : l/2]
			second := check[l/2:]

			if first == second {
				total += i
			}
		}
	}
	fmt.Println("Result: ", total)
}

func part2() {
	file, err := os.ReadFile("input02.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	input := string(file)
	ranges := strings.Split(input, ",")
	total := 0
	for _, r := range ranges {
		if r == "" {
			break
		}

		limits := strings.Split(r, "-")
		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])

		for i := min; i <= max; i++ {
			check := strconv.Itoa(i)
			l := len(check)

			for size := 1; size <= l/2; size++ {
				chunks := getChunks(check, size)
				if allSame(chunks) {
					total += i
					break
				}
			}
		}
	}
	fmt.Println("Result: ", total)
}

func getChunks(s string, chunkSize int) []string {
	var chunks []string
	size := len(s) / chunkSize
	for i := 0; i < size; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == size-1 {
			end = len(s)
		}
		chunks = append(chunks, s[start:end])
	}
	return chunks
}

func allSame(chunks []string) bool {
	for i := 1; i < len(chunks); i++ {
		if chunks[i] != chunks[0] {
			return false
		}
	}
	return true
}
