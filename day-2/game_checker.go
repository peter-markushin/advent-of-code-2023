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

var config = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

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
			c <- checkGame(s)
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

func checkGame(game string) int {
	gameId, gameLog, _ := s.Cut(game, ":")
	_, gameId, _ = s.Cut(gameId, " ")

	cubeSets := s.Split(gameLog, ";")
	for _, cubeSet := range cubeSets {
		cubes := s.Split(cubeSet, ",")

		for _, cube := range cubes {
			num, color, _ := s.Cut(s.Trim(cube, " "), " ")

			n, _ := sc.Atoi(num)
			maxNum, exists := config[color]

			if !exists || maxNum < n {
				return 0
			}
		}
	}

	id, _ := sc.Atoi(gameId)
	return id
}
