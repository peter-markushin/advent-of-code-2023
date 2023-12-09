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
	cur := make([]string, 0)

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

		if pt[2] == 'A' {
			cur = append(cur, pt)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	pathsLen := make([]int, len(cur))

	for i, _ := range cur {
		ptr := 0

		for cur[i][2] != 'Z' {
			if instructions[ptr] == 'L' {
				cur[i] = points[cur[i]].L
			} else {
				cur[i] = points[cur[i]].R
			}

			if ptr == len(instructions)-1 {
				ptr = 0
			} else {
				ptr += 1
			}

			pathsLen[i] += 1
		}

	}

	sum := pathsLen[0]

	if len(pathsLen) > 1 {
		for i := 1; i < len(pathsLen); i++ {
			sum = lcm(sum, pathsLen[i])
		}
	}

	fmt.Printf("\nSum %d", sum)
}

func lcm(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	return abs(a*b) / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
