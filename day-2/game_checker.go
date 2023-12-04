package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	sc "strconv"
	s "strings"
	"sync"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	c := make(chan int)
	sum := 0

	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			c <- checkGamePower(s)
		}(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for result := range c {
		sum += result
	}

	fmt.Printf("Sum: %d", sum)
}

func checkGamePower(game string) int {
	_, gameLog, _ := s.Cut(game, ":")

	var maxColor = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	cubeSets := s.Split(gameLog, ";")
	for _, cubeSet := range cubeSets {
		cubes := s.Split(cubeSet, ",")

		for _, cube := range cubes {
			num, color, _ := s.Cut(s.Trim(cube, " "), " ")
			n, _ := sc.Atoi(num)

			if n > maxColor[color] {
				maxColor[color] = n
			}
		}
	}

	return maxColor["red"] * maxColor["green"] * maxColor["blue"]
}
