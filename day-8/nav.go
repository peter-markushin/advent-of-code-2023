package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
)

type Point struct {
	L string
	R string
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := make(map[string]Point)
	instructions := ""

	for scanner.Scan() {
		line := scanner.Text()

		if instructions == "" {
			instructions = line
			continue
		}

		if line == "" {
			continue
		}

		pt, next, _ := s.Cut(line, " = ")
		L, R := next[1:4], next[6:9]
		points[pt] = Point{L: L, R: R}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	cur := "AAA"
	ptr := 0

	for cur != "ZZZ" {
		sum += 1

		if instructions[ptr] == 'L' {
			cur = points[cur].L
		} else {
			cur = points[cur].R
		}

		if ptr < len(instructions)-1 {
			ptr += 1
		} else {
			ptr = 0
		}
	}

	fmt.Printf("Sum %d", sum)
}
