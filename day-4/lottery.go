package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
			c <- checkCardPoints(s)
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

func checkCardPoints(card string) int {
	numMatches := 0
	_, allNumbers, _ := s.Cut(card, ": ")
	winningStr, cardNumbersStr, _ := s.Cut(allNumbers, " | ")

	winningNumbers := s.Split(winningStr, " ")
	cardNumbers := s.Split(cardNumbersStr, " ")

	for _, win := range winningNumbers {
		for _, num := range cardNumbers {
			if win != "" && win == num {
				numMatches += 1
			}
		}
	}

	if numMatches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(numMatches-1)))
}
