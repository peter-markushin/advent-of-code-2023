package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func parseNumber(s string) int {
	var n1, n2 byte
	i := 0

	for i = 0; s[i] < byte('0') || s[i] > byte('9'); i++ {
	}
	n1 = s[i] - byte('0')

	for i = len(s) - 1; s[i] < byte('0') || s[i] > byte('9'); i-- {
	}
	n2 = s[i] - byte('0')

	return int(n1*10 + n2)
}
