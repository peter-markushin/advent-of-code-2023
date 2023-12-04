package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	s "strings"
	"sync"
)

var digits = map[string]int{
	"0":     0,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
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
			c <- parseNumber(s)
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

func parseNumber(str string) int {
	n1, n2, pos1, pos2 := 0, 0, -1, -1

	for num, val := range digits {
		if !s.Contains(str, num) {
			continue
		}

		idx := s.Index(str, num)
		lastIdx := s.LastIndex(str, num)

		if pos1 == -1 || idx < pos1 {
			n1 = val
			pos1 = idx
		}

		if pos2 == -1 || lastIdx > pos2 {
			n2 = val
			pos2 = lastIdx
		}
	}

	return n1*10 + n2
}
