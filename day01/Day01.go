package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input01.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	var password = 0
	var pos = 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "L"):
			var numPart, _ = strings.CutPrefix(line, "L")
			var num, _ = strconv.Atoi(numPart)
			pos = (pos - num) % 100
			if pos == 0 {
				password = password + 1
			}
		case strings.HasPrefix(line, "R"):
			var numPart, _ = strings.CutPrefix(line, "R")
			var num, _ = strconv.Atoi(numPart)
			pos = (pos + num) % 100
			if pos == 0 {
				password = password + 1
			}
		}
	}

	fmt.Println("Password = ", password)
}

func part2() {
	file, err := os.Open("input01.txt")
	if err != nil {
		fmt.Printf("Could not open input file: %v\n", err)
	}
	defer file.Close()

	var password = 0
	var pos = 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "L"):
			var numPart, _ = strings.CutPrefix(line, "L")
			var num, _ = strconv.Atoi(numPart)
			password = password + (num / 100)
			num = num % 100
			if pos == 0 {
				pos = 100 - num
			} else {
				if num >= pos {
					password = password + 1
					pos = (pos - num + 100) % 100
				} else {
					pos = pos - num
				}
			}
		case strings.HasPrefix(line, "R"):
			var numPart, _ = strings.CutPrefix(line, "R")
			var num, _ = strconv.Atoi(numPart)
			password = password + (num / 100)
			num = num % 100
			if pos == 0 {
				pos = pos + num
			} else {
				if num >= (100 - pos) {
					password = password + 1
					pos = (pos + num + 100) % 100
				} else {
					pos = pos + num
				}
			}
		}
	}

	fmt.Println("Password = ", password)
}
