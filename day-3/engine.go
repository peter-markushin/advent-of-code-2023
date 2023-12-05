package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var engine []string

	for scanner.Scan() {
		engine = append(engine, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0

	for i, row := range engine {
		for j, symbol := range row {
			if symbol != rune('*') {
				continue
			}

			var gearNums []int

			//find numbers around
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if engine[i+y][j+x] < byte('0') || engine[i+y][j+x] > byte('9') {
						continue
					}

					var s, f int
					// if number found
					for s = x; j+s >= 0 && engine[i+y][j+s] >= byte('0') && engine[i+y][j+s] <= byte('9'); s-- {
					}
					// if number found
					for f = x; j+f < len(engine[i+y]) && engine[i+y][j+f] >= byte('0') && engine[i+y][j+f] <= byte('9'); f++ {
					}

					n, _ := sc.Atoi(engine[i+y][j+s+1 : j+f])
					x = f - 1
					gearNums = append(gearNums, n)
				}
			}

			if len(gearNums) != 2 {
				continue
			}

			sum += gearNums[0] * gearNums[1]
		}

	}

	fmt.Printf("Sum: %d", sum)
}
