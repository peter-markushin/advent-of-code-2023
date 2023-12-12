package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Galaxy struct {
	Row, Col int
}

const EXP_FACTOR = 999999

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	spaceHeight := 0
	spaceWidth := 0
	galaxies := make([]Galaxy, 0)

	for scanner.Scan() {
		line := scanner.Text()
		spaceWidth = len(line)
		hasGalaxies := false

		for i, spaceBit := range line {
			if spaceBit == '#' {
				hasGalaxies = true

				galaxies = append(galaxies, Galaxy{Row: spaceHeight, Col: i})
			}
		}

		if !hasGalaxies {
			spaceHeight += EXP_FACTOR
		}

		spaceHeight += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	hasGalaxiesAtCol := make([]bool, spaceWidth)
	expandedBy := 0

	for _, g := range galaxies {
		hasGalaxiesAtCol[g.Col] = true
	}

	for i, hasGalaxies := range hasGalaxiesAtCol {
		if hasGalaxies {
			continue
		}

		for x := range galaxies {
			if galaxies[x].Col-expandedBy > i {
				galaxies[x].Col += EXP_FACTOR
			}
		}

		expandedBy += EXP_FACTOR
	}

	sum := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += abs(galaxies[j].Col-galaxies[i].Col) + abs(galaxies[j].Row-galaxies[i].Row)
		}
	}

	fmt.Printf("\nSum %d", sum)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
